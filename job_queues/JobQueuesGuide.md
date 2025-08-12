# Job Queues and Worker Pools: Advanced Developer Guide

## Table of Contents
1. Purpose and Basics (what/why)
2. Queue Shapes: Unbounded, Bounded, Dropping, Backpressure
3. Worker Pool Design (jobs/results channels)
4. Cancellation, Timeouts, and Draining
5. Retries, Idempotency, and Ordering
6. Rate Limiting and Load Shedding
7. Common Mistakes and Gotchas
8. Best Practices
9. Performance Considerations
10. Advanced Challenge Questions

---

## 1) Purpose and Basics (what/why)

Job queues decouple producers from consumers. They:
- Smooth spikes in load (producer faster than consumer)
- Control concurrency with worker pools
- Provide a place to add retries, backoff, and metrics

---

## 2) Queue Shapes: Unbounded, Bounded, Dropping, Backpressure

- Unbounded: risky (can grow without limit) — usually avoid
- Bounded (buffered channel): applies backpressure when full
- Dropping: shed load by discarding new jobs or oldest
- Blocking: producers block until space (natural backpressure)

```go
// Bounded queue with backpressure
jobs := make(chan Job, 64) // cap controls buffering

// Dropping queue (shed load when full)
select {
case jobs <- j: // enqueued
default:
    // drop or count metric
}
```

---

## 3) Worker Pool Design (jobs/results channels)

Canonical pattern: N workers read jobs, write results; a closer goroutine closes results when workers finish.

```go
package main
import (
  "fmt"
  "sync"
)

type Job struct{ ID, X int }
type Result struct{ ID, Y int }

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
  defer wg.Done()
  for j := range jobs {
    results <- Result{ID: j.ID, Y: j.X * j.X}
  }
}

func main() {
  jobs := make(chan Job, 8)
  results := make(chan Result, 8)

  var wg sync.WaitGroup
  for w := 0; w < 3; w++ { wg.Add(1); go worker(w, jobs, results, &wg) }

  go func(){ // producer
    for i := 1; i <= 5; i++ { jobs <- Job{ID:i, X:i} }
    close(jobs)
  }()

  go func(){ wg.Wait(); close(results) }()

  for r := range results { fmt.Printf("%d:%d ", r.ID, r.Y) }
}
```

Notes:
- Sender closes jobs when done; closer closes results after workers finish
- Buffer sizes reduce contention; benchmark to choose

---

## 4) Cancellation, Timeouts, and Draining

Use context to cancel work; ensure goroutines select on ctx.Done() and that channels are closed/drained.

```go
package main
import (
  "context"
  "fmt"
  "sync"
  "time"
)

type Task func(context.Context) error

func worker(ctx context.Context, in <-chan Task, wg *sync.WaitGroup) {
  defer wg.Done()
  for {
    select {
    case <-ctx.Done():
      return
    case t, ok := <-in:
      if !ok { return }
      _ = t(ctx) // honor ctx inside task
    }
  }
}

func main() {
  ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
  defer cancel()

  in := make(chan Task, 4)
  var wg sync.WaitGroup
  for i := 0; i < 2; i++ { wg.Add(1); go worker(ctx, in, &wg) }

  in <- func(ctx context.Context) error { time.Sleep(100*time.Millisecond); fmt.Print("A "); return nil }
  in <- func(ctx context.Context) error { time.Sleep(300*time.Millisecond); fmt.Print("B "); return nil }
  close(in)
  wg.Wait()
}
```

Draining: if cancelling early, drain or close channels to avoid leaks; producers should stop producing on cancel.

---

## 5) Retries, Idempotency, and Ordering

Retries:
```go
func withRetry(ctx context.Context, max int, f func() error) error {
  var err error
  backoff := 10 * time.Millisecond
  for attempt := 1; attempt <= max; attempt++ {
    if err = f(); err == nil { return nil }
    select {
    case <-ctx.Done(): return ctx.Err()
    case <-time.After(backoff):
      backoff *= 2
    }
  }
  return err
}
```

Idempotency:
- Ensure job side-effects can be retried safely (e.g., upserts, natural keys)

Ordering:
- A worker pool does not preserve submission order by default; if required, include sequence numbers and reorder at sink

---

## 6) Rate Limiting and Load Shedding

Token bucket via a buffered channel:
```go
// capacity C, refill every tick
func limiter(capacity int, tick time.Duration) <-chan struct{} {
  tokens := make(chan struct{}, capacity)
  // fill initial burst
  for i := 0; i < capacity; i++ { tokens <- struct{}{} }
  go func(){ t := time.NewTicker(tick); defer t.Stop();
    for range t.C { select { case tokens <- struct{}{}: default: } }
  }()
  return tokens
}

// Usage
allow := limiter(10, 100*time.Millisecond)
select { case <-allow: /* do work */ default: /* shed load */ }
```

Leaky bucket with time.Ticker is also common; choose based on requirements.

---

## 7) Common Mistakes and Gotchas

1) Closing a channel from the receiver side
```go
// Only the producer/sender should close the jobs channel.
```

2) Forgetting to close results after workers finish
```go
// Use a closer goroutine: go func(){ wg.Wait(); close(results) }()
```

3) Leaking workers on cancel
```go
// Workers should select on ctx.Done() and return.
```

4) Unbounded queues
```go
// Use bounded buffers and backpressure; monitor lengths.
```

5) Assuming in-order processing
```go
// Pool completion order is nondeterministic; reorder if needed.
```

---

## 8) Best Practices

- Prefer bounded queues; backpressure protects upstream systems
- Separate types for Job and Result; keep messages small
- Make tasks idempotent; design retries with exponential backoff + jitter
- Propagate context; cancel on first error when appropriate (errgroup)
- Instrument queue lengths, worker utilizations, and error rates

---

## 9) Performance Considerations

- Fewer, faster workers beat many contending workers; benchmark
- Use small buffers to reduce contention; avoid per-item goroutines
- Avoid boxing to interface in hot paths (heap escapes)
- Batch jobs/results when possible to amortize overhead
- Measure with `go test -bench`, `pprof`, and `-race` for correctness

---

## 10) Advanced Challenge Questions

1) How do you prevent producer overload when consumers are slow?
- Use a bounded channel (backpressure) or a dropping queue (shed load) with metrics.

2) How do you avoid leaks when cancelling mid-stream?
- Propagate context cancellation; ensure producers stop enqueuing and workers select on ctx.Done().

3) How to preserve ordering with a worker pool?
- Include sequence numbers and reorder at the sink, or use a single worker (at the cost of throughput).

4) Where to implement retries — producer, worker, or downstream?
- Prefer closest to the side-effect, where idempotency is known; avoid duplicate retries across layers.

5) How would you rate-limit job execution?
- Token bucket using a buffered channel or a time.Ticker; gate job dispatch on token availability.

