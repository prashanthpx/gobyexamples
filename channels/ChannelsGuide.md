# Go Channels: Advanced Developer Guide

## Table of Contents
1. Channel Fundamentals (what/why)
2. Buffered vs Unbuffered Semantics
3. Send/Receive, Close, and Range
4. Channel Directions (chan<- / <-chan)
5. Select, Default, Timeouts, and Tickers
6. Nil Channels and Disabling Cases
7. Fan-in/Fan-out, Pipelines, Worker Pools
8. Cancellation Patterns (done vs context)
9. Common Mistakes and Gotchas
10. Best Practices
11. Performance Considerations
12. Advanced Challenge Questions


Run these examples
- Directional API: go run channels/examples/directional.go
- time.After leak vs Ticker: go run channels/mistakes/time_after_loop.go

---

## 1) Channel Fundamentals (what/why)

Channels connect concurrent goroutines by sending and receiving typed values.
- Synchronization: unbuffered channels synchronize sender and receiver
- Communication: pass ownership of data instead of sharing memory

```go
package main
import "fmt"

func main() {
    ch := make(chan int) // unbuffered
    go func() { ch <- 42 }()
    v := <-ch
    fmt.Println(v) // 42
}
```

Why channels matter:
- Eliminate explicit locking for many patterns
- Make concurrency easier to reason about (structured handoffs)

---

## 2) Buffered vs Unbuffered Semantics

- Unbuffered: send blocks until a receiver is ready; recv blocks until a value is available
- Buffered: capacity N; send blocks only when full, recv blocks only when empty

```go
package main
import "fmt"

func main() {
    ch := make(chan string, 2) // buffered
    ch <- "a"
    ch <- "b" // still not blocking
    // ch <- "c" // would block until a receiver runs
    fmt.Println(<-ch, <-ch) // a b
}
```

Use buffering to decouple producer/consumer speeds and reduce contention.

---

## 3) Send/Receive, Close, and Range

Closing signals that no more values will be sent.
- Only the sender should close
- Receiving from a closed channel yields zero value with ok=false
- Sending on a closed channel panics

```go
package main
import "fmt"

func main() {
    ch := make(chan int)
    go func() {
        for i := 0; i < 3; i++ { ch <- i }
        close(ch)
    }()
    for v := range ch { fmt.Print(v, " ") } // 0 1 2
}
```

Manual receive with comma-ok:
```go
v, ok := <-ch
if !ok { /* channel closed */ }
```

---

## 4) Channel Directions (chan<- / <-chan)

Restrict direction in signatures to document intent and prevent mis-use.

```go
package main
import "fmt"

func producer(out chan<- int) { for i := 0; i < 3; i++ { out <- i }; close(out) }
func consumer(in <-chan int)  { for v := range in { fmt.Print(v, " ") } }

func main() {
    ch := make(chan int, 3)
    go producer(ch)
    consumer(ch)
}
```

---

## 5) Select, Default, Timeouts, and Tickers

Select waits on multiple channel operations.

```go
package main
import (
  "fmt"
  "time"
)

func main() {
  ch := make(chan int)
  go func(){ time.Sleep(50*time.Millisecond); ch <- 1 }()

  select {
  case v := <-ch:
    fmt.Println("got", v)
  case <-time.After(10*time.Millisecond):
    fmt.Println("timeout")
  }
}
```

Default for non-blocking operations:
```go
select {
case ch <- 1:
    // sent
default:
    // would block; take alternate path
}
```

Tickers:
```go
t := time.NewTicker(100*time.Millisecond)
defer t.Stop()
for i := 0; i < 3; i++ {
    <-t.C
    fmt.Print("tick ")
}
```

---

## 6) Nil Channels and Disabling Cases

A nil channel blocks forever on send and receive. This is useful to disable select cases without extra flags.

```go
package main
import "time"

func main() {
    var ch <-chan int // nil
    timeout := time.After(50 * time.Millisecond)
    for {
        select {
        case <-ch: // never fires
        case <-timeout:
            return // exit after timeout
        }
    }
}
```

To disable a case dynamically, set the channel to nil.

---

## 7) Fan-in/Fan-out, Pipelines, Worker Pools

Fan-out workers, fan-in results:
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

  for w := 0; w < 3; w++ {
    wg.Add(1)
    go func(){ defer wg.Done(); for j := range jobs { results <- j*j } }()
  }

  go func(){ for i:=1;i<=5;i++{ jobs<-i }; close(jobs) }()
  go func(){ wg.Wait(); close(results) }()

  for r := range results { fmt.Print(r, " ") }
}
```

Pipeline stages:
```go
type item int
func gen(n int) <-chan item { ch := make(chan item); go func(){ for i:=0;i<n;i++{ ch<-item(i) }; close(ch) }(); return ch }
func square(in <-chan item) <-chan item { out := make(chan item); go func(){ for v := range in { out<-v*v }; close(out) }(); return out }
```

---

## 8) Cancellation Patterns (done vs context)

Done channel:
```go
done := make(chan struct{})
go func(){
  select { case <-done: return }
}()
close(done)
```

Prefer context for request-scoped lifetimes and timeouts:
```go
ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
defer cancel()
select {
case <-ctx.Done():
    // cancelled or timed out
}
```

Ensure all goroutines exit on cancellation and that senders close their channels when done.

---

## 9) Common Mistakes and Gotchas

8) time.After leak in loops
```go
// ❌ A new Timer is created on every iteration
for {
  select {
  case <-time.After(100 * time.Millisecond):
    work()
  case <-ctx.Done():
    return
  }
}

// ✅ Use a single Ticker (or reuse a Timer)
t := time.NewTicker(100 * time.Millisecond)
defer t.Stop()
for {
  select {
  case <-t.C:
    work()
  case <-ctx.Done():
    return
  }
}
```

9) Busy loop with select default
```go
// ❌ Burns CPU in a tight loop when nothing is ready
select {
case ch <- v:
  // sent
default:
  // spins; consider backoff or blocking path
}

// ✅ Add a backoff or proper wait path
select {
case ch <- v:
case <-time.After(1 * time.Millisecond):
  // backoff
}
```

1) Deadlock with unbuffered send and no receiver
```go
ch := make(chan int)
ch <- 1 // blocks forever if no receiver
```

2) Sending on a closed channel (panic)
```go
close(ch)
ch <- 1 // panic: send on closed channel
```

3) Reading from closed channel and ignoring ok
```go
v, ok := <-ch
if !ok { /* handle close */ }
```

4) Leaking goroutines waiting on channels
```go
// Always provide a cancellation path (ctx or done) in selects
```

5) Closing from receiver side
```go
// Only the sender should close the channel.
```

6) Assuming map or slice values are safe across goroutines
```go
// Use channels to transfer ownership, or protect shared structures with sync.
```

7) Using nil channels unintentionally (blocks forever)
```go
var ch chan int // nil; both send and recv block
```

---

## 10) Best Practices

- One-way channels in APIs to document intent (chan<- / <-chan)
- Close channels from the sender; receivers should range/comma-ok
- Prefer context over ad-hoc done channels for request lifetimes
- Buffer sizes: small buffers (1–N) unblock senders; benchmark to choose
- Avoid per-item goroutine creation; use worker pools or bounded semaphores
- Use select with default sparingly; avoid busy-waiting
- Send structs, not pointers to short-lived values, to avoid races

---

## 11) Performance Considerations

- Buffered channels reduce contention; try small powers of two, measure
- Large messages copy cost; consider pointers for large immutable data
- Avoid channel-of-interface hot paths (boxing causes heap escapes)
- Batch sends/receives to amortize sync overhead
- The race detector helps verify correctness: `go test -race`

---

## 12) Advanced Challenge Questions

(See also: Mistakes #8 and #9 above for time.After leaks and busy default)


1) What does receiving from a closed channel return and how do you detect it?
- Zero value and ok=false via `v, ok := <-ch`.

2) How can you disable a select case at runtime?
- Set that case’s channel to nil so it never fires.

3) When do you prefer mutex over channels?
- When protecting in-place shared state is simpler than message passing.

4) Why should only the sender close a channel?
- Receivers cannot know whether more senders exist; closing from receiver can race.

5) How do you avoid goroutine leaks in a pipeline?
- Propagate ctx.Done() through all stages, ensure producers close outputs, and consumers drain inputs when cancelling.

