# Concurrency Coordination: Advanced Developer Guide

## Table of Contents
1. Coordination Fundamentals (what/why)
2. sync.WaitGroup Deep Dive
3. sync.Once and Idempotent Initialization
4. sync.Cond (Condition Variables) and Signaling
5. errgroup With Context for Cancellation
6. Common Mistakes and Gotchas
7. Best Practices
8. Performance Considerations
9. Advanced Challenge Questions

---

## 1) Coordination Fundamentals (what/why)

Goroutines are cheap, but you must coordinate their lifecycles:
- Wait for completion (fan-out/fan-in)
- Signal state changes (ready, resource available)
- Perform one-time initialization safely
- Propagate cancellation and collect errors

---

## 2) sync.WaitGroup Deep Dive

WaitGroup waits for a collection of goroutines to finish.

Key rules:
- Call `Add(n)` before launching the n goroutines
- Each goroutine must call `Done()` exactly once
- `Wait()` blocks until the counter reaches zero
- Zero value of WaitGroup is ready to use; do not copy after first use

```go
package main
import (
  "fmt"
  "sync"
  "time"
)

func worker(id int, wg *sync.WaitGroup) {
  defer wg.Done()
  time.Sleep(50 * time.Millisecond)
  fmt.Println("done:", id)
}

func main() {
  var wg sync.WaitGroup
  const n = 3
  wg.Add(n)               // Add BEFORE launching
  for i := 0; i < n; i++ {
    go worker(i, &wg)
  }
  wg.Wait()
}
```

Pitfalls:
- Adding after launching can race (counter may hit zero early)
- Forgetting `Done()` leads to permanent `Wait()` block
- Reusing a WaitGroup concurrently across phases can cause panics

---

## 3) sync.Once and Idempotent Initialization

Use Once to ensure a function runs at most once across goroutines.

```go
package main
import (
  "fmt"
  "sync"
)

var (
  once sync.Once
  initVal int
)

func initHeavy() { initVal = 42 }

func main() {
  var wg sync.WaitGroup
  for i := 0; i < 5; i++ {
    wg.Add(1)
    go func() {
      defer wg.Done()
      once.Do(initHeavy)
    }()
  }
  wg.Wait()
  fmt.Println(initVal) // 42
}
```

Notes:
- `Once` guarantees the function body runs only once, even if `Do` is called concurrently
- If the function panics, it is considered not done, and a later call may run it again

---

## 4) sync.Cond (Condition Variables) and Signaling

`sync.Cond` coordinates goroutines waiting for a condition while holding a lock.

```go
package main
import (
  "fmt"
  "sync"
)

func main() {
  mu := &sync.Mutex{}
  cond := sync.NewCond(mu)
  ready := false

  // Waiter
  var wg sync.WaitGroup
  wg.Add(1)
  go func() {
    defer wg.Done()
    cond.L.Lock()
    for !ready {         // use a loop to guard spurious wakeups
      cond.Wait()
    }
    fmt.Println("proceed")
    cond.L.Unlock()
  }()

  // Notifier
  cond.L.Lock()
  ready = true
  cond.Signal()          // wake one waiter (Broadcast() to wake all)
  cond.L.Unlock()

  wg.Wait()
}
```

When to use Cond:
- Queue-like coordination (producer/consumer) when channels are not a good fit
- Complex predicates under a mutex; always check the predicate in a loop

---

## 5) errgroup With Context for Cancellation

Run this example
- go run wait_group/examples/errgroup_pipeline.go


`errgroup` (golang.org/x/sync/errgroup) runs multiple functions, cancels on first error, and aggregates errors.

```go
package main
import (
  "context"
  "fmt"
  "net/http"
  "time"
  "golang.org/x/sync/errgroup"
)

func main() {
  ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
  defer cancel()

  g, ctx := errgroup.WithContext(ctx)

  urls := []string{"https://example.com", "https://example.org"}
  for _, u := range urls {
    u := u // capture
    g.Go(func() error {
      req, _ := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
      resp, err := http.DefaultClient.Do(req)
      if err != nil { return err }
      defer resp.Body.Close()
      return nil
    })
  }

  if err := g.Wait(); err != nil {
    fmt.Println("failed:", err)
  } else {
    fmt.Println("all ok")
  }
}
```

Notes:
- `WithContext` returns a derived context that is cancelled on first error
- Always ensure goroutines honor `ctx.Done()` or use context-aware APIs

---

## 6) Common Mistakes and Gotchas

6) One-shot select processes only one job
```go
// ❌ Processes a single job then exits select
func worker(in <-chan Job) {
  select {
  case j := <-in:
    _ = j.Run()
  }
}

// ✅ Loop or range over the channel
func workerFixed(in <-chan Job) {
  for j := range in { _ = j.Run() }
}
```


1) Using WaitGroup incorrectly
```go
// ❌ Adding inside goroutine can race with Wait
// ✅ Call Add(n) BEFORE launching n goroutines
```

2) Copying a WaitGroup
```go
// ❌ Copying after use causes panic
// Pass *sync.WaitGroup into goroutines
```

3) Forgetting to loop on Cond.Wait
```go
// ❌ Using if instead of for leads to missed wake-ups
// ✅ Always check predicate in a for-loop
```

4) Leaking goroutines on error paths
```go
// ✅ Use errgroup.WithContext or explicit done channels to cancel others
```

5) Data races on shared variables
```go
// ✅ Guard with sync.Mutex/RWMutex or use channel ownership
```

---

## 7) Best Practices

- Prefer channels and errgroup for pipelines; WaitGroup for basic joins
- Always call `Add` before launching; use `defer wg.Done()` in goroutines
- Use `sync.Once` for global init and singletons; keep function small and fast
- For Cond, keep critical sections small and always check predicates in a loop
- Propagate context through APIs that might block or take time

---

## 8) Performance Considerations

- WaitGroup and Once are very cheap; the dominant cost is the work itself
- Cond can reduce busy-waiting compared to polling; design predicates carefully
- errgroup reduces boilerplate and prevents wasted work after first failure
- Prefer batching to reduce synchronization overhead in hot loops

---

## 9) Advanced Challenge Questions

1) Why must `wg.Add(n)` be called before starting goroutines?
- To prevent a race where `Wait()` observes zero and returns before `Add` increments the counter.

2) When would you choose `sync.Cond` over channels?
- When multiple conditions and a shared predicate under a mutex must be coordinated, or when wake-one semantics are needed efficiently.

3) What happens if the function passed to `once.Do` panics?
- The once is considered not done; a subsequent call may run the function again.

4) How does `errgroup.WithContext` cancel remaining work?
- It cancels the derived context; goroutines must check `ctx.Done()` or use context-aware calls.

5) Why is copying a `WaitGroup` after use a problem?
- Internal state becomes inconsistent; the runtime will detect and panic.

