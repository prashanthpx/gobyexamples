# Go Channels: Advanced Developer Guide

## Table of Contents
1. [Channel Fundamentals (what/why)](#toc-1-fundamentals)
2. [Buffered vs Unbuffered Semantics](#toc-2-buffering)
3. [Send/Receive, Close, and Range](#toc-3-send-recv-close-range)
4. [Channel Directions (chan<- / <-chan)](#toc-4-directions)
5. [Select, Default, Timeouts, and Tickers](#toc-5-select-timeouts-tickers)
6. [Nil Channels and Disabling Cases](#toc-6-nil-channels)
7. [Fan-in/Fan-out, Pipelines, Worker Pools](#toc-7-fanin-fanout)
8. [Cancellation Patterns (done vs context)](#toc-8-cancellation)
9. [Common Mistakes and Gotchas](#toc-9-mistakes)
10. [Best Practices](#toc-10-best-practices)
11. [Performance Considerations](#toc-11-performance)
12. [Advanced Challenge Questions](#toc-12-advanced-questions)
13. [Channels as Reference Types and When to Use *chan](#toc-13-channel-ref-and-ptr)
14. Two-Value Patterns Cheat Sheet (comma-ok and range) — see Operators guide: [link](../Operators/OperatorsGuide.md#toc-7-2value)


Run these examples
- Directional API: go run channels/examples/directional.go
- time.After leak vs Ticker: go run channels/mistakes/time_after_loop.go

---

<a id="toc-1-fundamentals"></a>

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

<a id="toc-2-buffering"></a>

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

### Q: Is `ch := make(chan int)` the same as `ch := make(chan int, 1)`?

A: Not the same — both create `chan int`, but buffering differs.

1) `ch := make(chan int)`
- Unbuffered (capacity = 0)
- Send blocks until a receiver is ready; receive blocks until a sender arrives
- Forces synchronization between sender and receiver

Example:
```go
ch := make(chan int)

go func() {
    ch <- 42 // blocks until someone receives
}()

v := <-ch // will unblock sender
fmt.Println(v) // 42
```

2) `ch := make(chan int, 1)`
- Buffered with capacity 1
- First send does NOT block (goes into buffer); second send blocks until a receive occurs
- Receive blocks only when buffer is empty

Example:
```go
ch := make(chan int, 1)

ch <- 42 // does NOT block (buffer has room)
fmt.Println("sent 42")

v := <-ch // removes from buffer
fmt.Println(v) // 42
```

3) General difference
- Unbuffered: communication = synchronization (rendezvous)
- Buffered: communication = queueing (up to capacity N, then blocks)

So:
- `make(chan int)` → capacity = 0 (unbuffered)
- `make(chan int, N)` → capacity = N (buffered)

✅ Answer: They are not the same. The first is unbuffered (blocking handoff), the second is buffered with capacity 1 (allows one value to sit in the channel without a receiver).


---

<a id="toc-3-send-recv-close-range"></a>

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

<a id="toc-4-directions"></a>

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

<a id="toc-5-select-timeouts-tickers"></a>

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

### How select chooses a case (and why it never blocks after picking)

Rules the runtime follows each time select runs:
- It inspects all cases and builds a list of those that are ready now (non-blocking)
- If none are ready and there is no `default`, it blocks until one becomes ready
- If exactly one is ready, it executes that one
- If multiple are ready, it picks exactly one pseudo-randomly among them (order in code gives no priority)

Therefore: a selected case is guaranteed to be immediately executable. The runtime never “picks first, then blocks”.

Example pattern with Ticker + cancellation:
```go
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
- If `t.C` isn’t ready but `ctx.Done()` is, select will choose the `ctx.Done()` case
- If both are ready, it picks one randomly; the loop then iterates and re-evaluates

Notes on Ticker behavior:
- Ticker does not accumulate unlimited ticks; if `work()` is slow you can miss intermediate ticks and just get the most recent tick
- If you want slightly more eager cancellation, you can check `ctx.Err()` before the select


---

<a id="toc-6-nil-channels"></a>

## 6) Nil Channels and Disabling Cases

A nil channel blocks forever on send and receive. This is useful to disable select cases without extra flags.

### Q: Is `var c chan int` the same as `c := make(chan int)`?

A: Not the same — they differ in initialization and usability.

1) `var c chan int`
- Declares a channel variable of type `chan int`
- Zero value is `nil`; no backing runtime structure
- Sending (`c <- 1`) or receiving (`<-c`) blocks forever (deadlock)
- You must call `make` later to allocate the channel

Example:
```go
var c chan int
fmt.Println(c == nil) // true
// c <- 1 // would deadlock forever
```

2) `c := make(chan int)`
- Allocates and initializes a real channel (unbuffered by default, capacity 0)
- Ready to use immediately; send/receive follow normal blocking rules

Example:
```go
c := make(chan int)
go func() { c <- 42 }()
fmt.Println(<-c) // 42
```

3) Key differences
- `var c chan int` → declares only; unusable until `make`; nil channel blocks forever on send/receive
- `c := make(chan int)` → declares + initializes; usable immediately

✅ Summary:
- `var c chan int` declares a nil channel variable; call `make` before use
- `c := make(chan int)` declares and initializes the channel so it’s ready to use


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

<a id="toc-7-fanin-fanout"></a>

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

<a id="toc-8-cancellation"></a>

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

<a id="toc-9-mistakes"></a>

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

<a id="toc-10-best-practices"></a>

## 10) Best Practices

- One-way channels in APIs to document intent (chan<- / <-chan)
- Close channels from the sender; receivers should range/comma-ok
- Prefer context over ad-hoc done channels for request lifetimes
- Buffer sizes: small buffers (1–N) unblock senders; benchmark to choose
- Avoid per-item goroutine creation; use worker pools or bounded semaphores
- Use select with default sparingly; avoid busy-waiting
- Send structs, not pointers to short-lived values, to avoid races

---

<a id="toc-11-performance"></a>

## 11) Performance Considerations

- Buffered channels reduce contention; try small powers of two, measure
- Large messages copy cost; consider pointers for large immutable data
- Avoid channel-of-interface hot paths (boxing causes heap escapes)
- Batch sends/receives to amortize sync overhead
- The race detector helps verify correctness: `go test -race`

---

<a id="toc-12-advanced-questions"></a>

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



---

<a id="toc-13-channel-ref-and-ptr"></a>

## 13) Channels as Reference Types and When to Use *chan

### 1) Are channels already references in Go?

Channels are reference types (like slices, maps, functions). Passing `chan T` to a function passes a descriptor that refers to the same underlying channel. Sender and receiver share it.

Example:
```go
func producer(ch chan int) {
    ch <- 42
}

func main() {
    ch := make(chan int)
    go producer(ch)
    fmt.Println(<-ch) // prints 42
}
```
Here, `ch` is passed by value, but the value itself is a reference to the same channel.

### 2) Can you pass a pointer to a channel?

Yes, you can pass `*chan T`, but it's rarely needed. It's only useful when the function must reassign which channel variable the caller sees:
```go
func modifyChannel(ch *chan int) {
    *ch = make(chan int, 10) // reassign caller's channel variable
}

func main() {
    var ch chan int // nil
    modifyChannel(&ch) // pass pointer so reassignment is visible to caller
    ch <- 5
    fmt.Println(<-ch) // prints 5
}
```

### 3) When should you use a pointer to a channel?

- Reassignment inside a function: create or swap out the channel and reflect that change in the caller
- Rare indirection patterns: APIs/data structures that conditionally replace a channel

Most of the time, just pass `chan T` directly.

### 4) Best practice

- Use `chan T` normally — it already behaves like a reference
- Use `*chan T` only if you truly need to change which channel the caller's variable refers to (uncommon; may hint at redesign)

✅ Summary
- Channels are reference types in Go; passing `chan T` is enough in 99% of cases
- Passing `*chan T` is only useful if the function needs to reassign the channel itself
