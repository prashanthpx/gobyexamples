# Go Goroutines: Advanced Developer Guide

## Table of Contents
1. Goroutine Fundamentals (what and why)
2. Scheduler Model (G-M-P), Stacks, Parallelism
3. Launching and Waiting (WaitGroup, done patterns)
4. Cancellation with context.Context
5. Synchronization: Mutex, Atomic, Channels (when to use which)
6. Blocking, Preemption, and Syscalls
7. Common Concurrency Patterns
8. Goroutine Leaks and How to Avoid Them
9. Best Practices
10. Performance Considerations
11. Advanced Challenge Questions


Run these examples
- Race detector: go run -race goroutine/race/race.go
- Benchmarks: go test -bench=. -cpu=1,4 -benchtime=200ms ./goroutine/bench

---

## 1) Goroutine Fundamentals (what and why)

Goroutines are lightweight concurrent execution units managed by the Go runtime. They:
- Are cheap to create (few KB initial stack, grows/shrinks dynamically)
- Multiplex onto OS threads by the scheduler
- Enable scalable concurrency with simple syntax

```go
package main
import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(3)
    for i := 1; i <= 3; i++ {
        go func(id int) {
            defer wg.Done()
            fmt.Println("hello from goroutine", id)
        }(i)
    }
    wg.Wait()
}
```

Why goroutines matter:
- Concurrency by default, easy to scale I/O bound work
- Pair naturally with channels for communication

---

## 2) Scheduler Model (G-M-P), Stacks, Parallelism

Concepts:
- G (goroutine), M (machine/OS thread), P (processor/run queue)
- GOMAXPROCS controls max P (degree of parallelism across cores)
- Goroutine stacks start small (~2KB) and grow; avoids large per-goroutine memory

```go
package main
import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println("CPUs:", runtime.NumCPU())
    prev := runtime.GOMAXPROCS(0) // read current
    fmt.Println("GOMAXPROCS:", prev)
}
```

Notes:
- CPU-bound tasks can run in parallel up to GOMAXPROCS
- I/O/syscalls park goroutines; scheduler runs others

---

## 3) Launching and Waiting (WaitGroup, done patterns)

Use sync.WaitGroup to wait for a set of goroutines.

```go
package main
import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    time.Sleep(100 * time.Millisecond)
    fmt.Println("done:", id)
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }
    wg.Wait()
}
```

A minimal "done" channel to stop a goroutine:
```go
package main
import (
    "fmt"
    "time"
)

func ticker(done <-chan struct{}) {
    t := time.NewTicker(50 * time.Millisecond)
    defer t.Stop()
    for {
        select {
        case <-t.C:
            fmt.Print("tick ")
        case <-done:
            fmt.Println("stop")
            return
        }
    }
}

func main() {
    done := make(chan struct{})
    go ticker(done)
    time.Sleep(200 * time.Millisecond)
    close(done)
    time.Sleep(50 * time.Millisecond)
}
```

---

## 4) Cancellation with context.Context

Run this example
- go run goroutine/examples/cancel_pipeline.go

Context propagates deadlines, timeouts, and cancellation.

```go
package main
import (
    "context"
    "fmt"
    "time"
)

func doWork(ctx context.Context) error {
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
        case <-time.After(50 * time.Millisecond):
            fmt.Print("work ")
        }
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
    defer cancel()
    _ = doWork(ctx)
    fmt.Println("ended")
}
```

Tips:
- Always accept ctx as first param: func(ctx context.Context, ...)
- Do not store ctx in struct; pass it down the call chain

---

## 5) Synchronization: Mutex, Atomic, Channels (when to use which)

Channels + Locks: when to combine
- Use channels to distribute work and transfer ownership of items (no shared mutation while processing)
- Use locks to protect shared aggregation structures (e.g., maps, counters) that multiple workers update
- Avoid forcing everything through channels if you’re really aggregating into shared state — a simple mutex is usually clearer and faster

Diagram
- Producer(s) -> jobs chan -> Workers -> (mu.Lock) shared map/counter (mu.Unlock)
- Cancellation via context propagated to producers/workers; workers exit, then close results

Run these examples
- Channel + Mutex (counts): go run goroutine/examples/chan_mutex_counts.go

- Mutex (sync.Mutex/RWMutex): protect shared mutable state
- Atomic (sync/atomic): low-level lock-free ops on single words (counters/flags)
- Channels: communicate ownership or data; avoid sharing state

```go
package main
import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    var mu sync.Mutex
    var counter int64

    // Mutex example
    mu.Lock(); counter++; mu.Unlock()

    // Atomic example
    atomic.AddInt64(&counter, 1)

    fmt.Println(counter)
}
```

Guideline: choose the simplest correct tool. Channels for communication; mutex/atomic for protected shared state.

---

## 6) Blocking, Preemption, and Syscalls

- Blocking operations (I/O, time.Sleep, channel send/recv) park the goroutine
- Go 1.14+ supports asynchronous preemption; long CPU loops are preemptible, but still honor cancellation points
- Syscalls may block OS threads; runtime manages handoff to keep P utilized

```go
// CPU-bound loop with cancellation
for {
    select { case <-ctx.Done(): return default: }
    // compute...
}
```

Use runtime.Gosched() rarely; the scheduler is generally sufficient.

---

## 7) Common Concurrency Patterns

Run these examples
- Channel + Mutex (counts): go run goroutine/examples/chan_mutex_counts.go
- Cancellable generator: go run goroutine/examples/generator_fixed.go
- Fan-in with context cancellation: go run goroutine/examples/fanin_context.go
- Rate limiter (token bucket): go run goroutine/examples/rate_limiter_ticker.go
- Worker pool with graceful shutdown: go run goroutine/examples/worker_pool_shutdown.go

Fan-out / Fan-in:
```go
package main
import (
    "fmt"
    "sync"
)

func main() {
    jobs := make(chan int)
    results := make(chan int)

    var wg sync.WaitGroup
    // Fan-out workers
    for w := 0; w < 3; w++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := range jobs {
                results <- j * j
            }
        }()
    }

    // Producer
    go func() {
        for i := 1; i <= 5; i++ { jobs <- i }
        close(jobs)
    }()

    // Fan-in closer
    go func() { wg.Wait(); close(results) }()

    for r := range results { fmt.Print(r, " ") }
}
```

Bounded worker pool:
```go
sem := make(chan struct{}, 10) // max 10 concurrent
for _, task := range tasks {
    sem <- struct{}{}
    go func(t Task){
        defer func(){ <-sem }()
        t.Run()
    }(task)
}
// Drain: ensure all slots released
for i := 0; i < cap(sem); i++ { sem <- struct{}{} }
```

Pipeline with cancellation:
```go
type item int
func stage1(out chan<- item, n int) { for i := 0; i < n; i++ { out <- item(i) } close(out) }
func stage2(in <-chan item, out chan<- item) { for v := range in { out <- v*v } close(out) }
```

---

## 8) Goroutine Leaks and How to Avoid Them

Context propagation anti-patterns and fixes:
```go
// ❌ Don’t store context in a struct for later use
// type Client struct { ctx context.Context } // avoid
// ✅ Accept ctx as first parameter and pass it through
func Fetch(ctx context.Context, url string) error { /* ... */ return nil }
```

Loop cancellation points:
```go
// ✅ Check ctx.Done in loops to exit promptly
for {
  select {
  case <-ctx.Done(): return
  default:
    // work
  }
}
```


Leak pattern: goroutine blocked forever on send/recv with no way to exit.

```go
// ❌ Leak: recv will never happen if main returns early
ch := make(chan int)
go func(){ ch <- 1 }()
// ... main exits -> goroutine blocked forever
```

Avoidance:
- Always ensure goroutines can exit: close channels, send stop signals, or use contexts
- Prefer select with ctx.Done() or a done channel when sending/receiving

```go
// ✅ Safe with context
ctx, cancel := context.WithCancel(context.Background())
go func(){
    defer fmt.Println("clean exit")
    select {
    case ch <- 1:
    case <-ctx.Done():
        return
    }
}()
cancel()
```

Missing receiver on unbuffered channel:
- Use buffered channel (cap 1) for notifications to avoid missed signals, or ensure receiver is ready before send

---

## 9) Best Practices

- Keep goroutines short-lived or clearly owned (who cancels/Waits?)
- Pass context for cancellation/timeouts; honor ctx.Done() in loops
- Avoid capturing loop variables incorrectly; pass as param to goroutine
- Prefer structured concurrency: tie lifetimes to parents; use WaitGroup
- Avoid unbounded goroutine creation; use pools/semaphores
- Document thread-safety expectations of your types

---

## 10) Performance Considerations

Profiling quickstart
- CPU profile: go run pprof/examples/cpu_profile.go -cpuprofile cpu.out; then go tool pprof cpu.out
- Heap profile: go run pprof/examples/heap_profile.go -memprofile mem.out; then go tool pprof mem.out
- In pprof: top, list <func>, web (requires graphviz)


- Goroutines are cheap but not free: each adds stack + scheduler overhead
- Reduce allocations in hot paths; avoid boxing to interface when possible
- Tune GOMAXPROCS for CPU-bound workloads (usually defaults to NumCPU)
- Prefer batch operations to reduce contention on channels/mutexes
- Measure with benchmarks and the race detector: `go test -race`, `pprof`


Microbenchmarks: Mutex vs Channel increment
````go
// Save as counter_test.go in your package and run: go test -bench=. -cpu=1,4 -benchtime=200ms
package counter

import (
  "sync"
  "testing"
)

func BenchmarkMutex(b *testing.B) {
  var mu sync.Mutex
  var n int
  b.RunParallel(func(pb *testing.PB) {
    for pb.Next() {
      mu.Lock(); n++; mu.Unlock()
    }
  })
}

func BenchmarkChannel(b *testing.B) {
  ch := make(chan int, 1024)
  done := make(chan struct{})
  go func(){
    for {
      select {
      case <-done: return
      case <-ch:
      }
    }
  }()
  b.RunParallel(func(pb *testing.PB) {
    for pb.Next() { ch <- 1 }
  })
  close(done)
}
````

Race detector quickcheck (see runnable: goroutine/race/race.go)
````go
// Save as race.go and run: go run -race race.go
package main
import (
  "fmt"
  "sync"
)

func main(){
  var x int
  var wg sync.WaitGroup
  for i := 0; i < 1000; i++ {
    wg.Add(1)
    go func(){ x++; wg.Done() }()
  }
  wg.Wait()
  fmt.Println(x)
}
````

---

## 11) Advanced Challenge Questions

(See also: Section 8 for leak patterns and context propagation)


1) Why can massive numbers of goroutines still be efficient?
- Small stacks that grow, multiplexing onto a few OS threads, and cooperative blocking.

2) How do you avoid goroutine leaks with channel-based APIs?
- Provide cancellation (ctx or done), close channels when producers exit, and ensure consumers drain.

3) When would you choose mutex/atomic over channels?
- When coordinating access to shared data (stateful object) is simpler than message passing, or for single-word counters/flags.

4) What does GOMAXPROCS control and when would you change it?
- Degree of parallelism (number of P). Adjust for CPU-bound workloads or when embedding Go in constrained environments.

5) How do syscalls affect scheduling?
- A syscall can block the current M; the runtime parks it and schedules runnable Gs on other Ms to keep Ps busy.

