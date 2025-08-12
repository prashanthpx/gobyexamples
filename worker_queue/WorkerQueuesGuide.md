# Worker Queues: Advanced Developer Guide

## Table of Contents
1. When to Use a Worker Queue vs Plain Goroutines
2. Core Designs (Shared Jobs Chan vs Ready-Pool Dispatcher)
3. Queue API Design (Submit/Stop/Results)
4. Shutdown Semantics and Draining
5. Timeouts, Cancellation, and Retries
6. Load Shedding and Backpressure
7. Common Mistakes and Gotchas (with fixes)
8. Best Practices
9. Performance Considerations
10. Advanced Challenge Questions

---

## 1) When to Use a Worker Queue vs Plain Goroutines

Use a worker queue when you need:
- Bounded concurrency with backpressure
- Long-running worker lifecycles (warm caches, connections)
- Centralized scheduling or prioritization
- Structured shutdown (finish in-flight, stop accepting)

Plain goroutines + a counting semaphore (buffered chan) are fine for simple burst control. Worker queues shine for control and observability.

---

## 2) Core Designs (Shared Jobs Chan vs Ready-Pool Dispatcher)

### A) Shared Jobs Channel (simplest)
- N workers read from one `jobs <-chan Job`
- Producer closes jobs; a closer goroutine closes results after workers finish

```go
package main
import (
  "fmt"
  "sync"
)

type Job struct{ ID int }
func (j Job) Run() string { return fmt.Sprintf("done-%d", j.ID) }

type Result struct{ ID int; Out string }

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
  defer wg.Done()
  for j := range jobs { results <- Result{ID: j.ID, Out: j.Run()} }
}

func main() {
  jobs := make(chan Job, 8); results := make(chan Result, 8)
  var wg sync.WaitGroup
  for w := 0; w < 3; w++ { wg.Add(1); go worker(w, jobs, results, &wg) }
  go func(){ for i:=1;i<=5;i++{ jobs<-Job{ID:i} }; close(jobs) }()
  go func(){ wg.Wait(); close(results) }()
  for r := range results { fmt.Print(r.Out, " ") }
}
```

### B) Ready-Pool Dispatcher (more control)
- Each worker has its own input channel `chan Job`
- Idle workers advertise readiness by sending their input chan on `readyPool chan chan Job`
- Dispatcher selects: when a job arrives and a worker is ready, it forwards the job to that worker

```go
// Core channels
// inputQueue: submissions; readyPool: idle worker mailboxes
// dispatcher matches inputQueue with readyPool
```

Pros: precise matching, per-worker state, easy to plug priorities. Cons: more code.

---

## 3) Queue API Design (Submit/Stop/Results)

A minimal queue with explicit lifecycle:

```go
package wq
import "sync"

type Job interface{ Run() error }

type Queue struct {
  inputQueue chan Job
  readyPool  chan chan Job
  stop       chan struct{}
  wg         sync.WaitGroup // tracks workers
}

func New(size, workers int) *Queue {
  q := &Queue{
    inputQueue: make(chan Job, size),
    readyPool:  make(chan chan Job, workers),
    stop:       make(chan struct{}),
  }
  // start dispatcher
  go q.dispatch()
  // start workers
  for i := 0; i < workers; i++ { go q.worker() }
  return q
}

func (q *Queue) Submit(j Job) error {
  select {
  case q.inputQueue <- j:
    return nil
  case <-q.stop:
    return ErrStopped
  }
}

func (q *Queue) Stop() {
  // signal dispatcher and wait workers to drain
  close(q.stop)
  // Optional: wait externally with a separate Wait method
}

func (q *Queue) dispatch() {
  for {
    select {
    case j := <-q.inputQueue:
      // wait for a ready worker
      w := <-q.readyPool
      w <- j
    case <-q.stop:
      // best-effort drain inputQueue then exit
      for {
        select { case j := <-q.inputQueue: _ = j // drop or route; case <-q.stop: default: return }
      }
    }
  }
}

func (q *Queue) worker() {
  inbox := make(chan Job)
  q.wg.Add(1)
  defer q.wg.Done()
  for {
    // advertise readiness
    select {
    case q.readyPool <- inbox:
    case <-q.stop:
      return
    }
    // receive work or stop
    select {
    case j := <-inbox: _ = j.Run()
    case <-q.stop: return
    }
  }
}
```

Notes:
- In production, add a `Wait()` method that waits on `q.wg` and ensure dispatcher exit before returning.
- Decide whether `Stop()` blocks (graceful) or is fire-and-forget.

---

## 4) Shutdown Semantics and Draining

Goals:
- Stop accepting new jobs
- Finish in-flight jobs (graceful) or cancel them (fast)
- Avoid leaking goroutines/channels

Patterns:
- `Stop()` closes a quit chan; dispatcher drains input; workers exit after inbox is empty
- `StopNow()` cancels via context passed to jobs (cooperative cancellation); workers check ctx

Ensure the closer goroutine closes results channels after all workers are done, if you expose results.

---

## 5) Timeouts, Cancellation, and Retries

Pass `context.Context` into `Job.Run(ctx)` to support deadlines.

```go
type CtxJob interface{ Run(ctx context.Context) error }
// Workers: select on ctx.Done() in long operations; return early on cancel.
```

Retries belong close to the side-effect where idempotency is understood. Use exponential backoff + jitter; cap attempts.

---

## 6) Load Shedding and Backpressure

- Use a bounded `inputQueue` to apply backpressure
- For overload, prefer dropping policy or a separate overflow queue with metrics
- Expose `len(inputQueue)` and worker utilization for autoscaling decisions

---

## 7) Common Mistakes and Gotchas (with fixes)

1) Closing channels from the wrong side
```go
// Only the submitter side closes inputQueue when no more jobs will arrive.
```

2) Single receive then exit (select) — only one job processed
```go
// Loop on receive; range over channel or for-select. Avoid one-shot selects.
```

3) Leaking goroutines on Stop
```go
// Ensure workers select on stop signal and dispatcher drains inputQueue.
```

4) Sending on closed channels
```go
// Guard Submit with stop chan; after Stop, return ErrStopped.
```

5) Unbounded growth
```go
// Always set a capacity; shed load if full to protect upstream systems.
```

6) Head-of-line blocking
```go
// Long jobs monopolize workers; consider job sharding or size-based queues.
```

---

## 8) Best Practices

- Start with the shared-jobs-channel design; move to ready-pool only if needed
- Keep Job small (data + method); avoid capturing huge closures
- Make jobs idempotent; add retries with backoff at the worker
- Provide observability: queue length, in-flight, success/error counts, latency
- Encapsulate Stop/Wait to make shutdown deterministic

---

## 9) Performance Considerations

- Prefer few hot workers over many contending threads; benchmark
- Size buffers to reduce contention but avoid large resident sets
- Avoid boxing to `interface{}` in hot paths; define explicit Job interfaces
- Consider batching jobs in workers to amortize sync cost
- Use `-race` during development to catch misuse; load test under production-like conditions

---

## 10) Advanced Challenge Questions

1) Compare the shared-jobs-channel pool with the ready-pool dispatcher.
- Shared jobs is simpler and adequate for most; ready-pool enables matching, per-worker state, and custom scheduling at the cost of complexity.

2) How do you guarantee no jobs are lost during shutdown?
- Stop accepting new jobs, drain the input queue, track in-flight with a WaitGroup, and close results only after workers complete.

3) Where do retries belong and how do you ensure safety?
- Close to the side-effect (inside Job) with idempotency guarantees; use capped exponential backoff with jitter.

4) How do you handle overload without crashing upstreams?
- Bounded queues (backpressure) plus dropping policy or 429-style rejection with metrics and autoscaling hooks.

5) How would you add priorities?
- Multiple input queues (high/low) with a dispatcher that selects preferentially; or a heap-backed priority queue feeding workers.

