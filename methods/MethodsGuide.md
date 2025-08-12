# Go Methods: Advanced Developer Guide

## Table of Contents
1. Method Fundamentals (what they are and why they matter)
2. Receiver Semantics (value vs pointer)
3. Method Sets and Interface Satisfaction
4. Method Values vs Method Expressions
5. Embedding and Promoted Methods
6. Concurrency and Receiver Design
7. Common Mistakes and Gotchas
8. Best Practices
9. Performance Considerations
10. Advanced Challenge Questions

---

## 1) Method Fundamentals (what they are and why they matter)

Methods in Go are functions with a special receiver argument. They:
- Attach behavior to types (including struct and named non-struct types)
- Enable polymorphism via interfaces without inheritance
- Let you control mutation semantics via receiver choice

```go
package main
import "fmt"

type Counter int

func (c Counter) Snapshot() int { // value receiver
    return int(c)
}
func (c *Counter) Inc() { // pointer receiver
    *c++
}

func main() {
    var c Counter
    c.Inc()
    fmt.Println(c.Snapshot()) // 1
}
```

Why methods matter:
- Clarify API semantics (mutating vs non-mutating)
- Enable interface satisfaction and decoupling
- Improve readability/cohesion by colocating behavior with data

---

## 2) Receiver Semantics (value vs pointer)

- Value receiver: method gets a copy of the value
- Pointer receiver: method receives address; can mutate the original

Guidelines:
- Use pointer receivers when:
  - The method mutates the receiver
  - The receiver is large (avoid copying)
  - Consistency with other methods requires it
- Use value receivers when:
  - The receiver is small/immutable
  - You want copy semantics by design

```go
package main
import "fmt"

type Wallet struct{ Cents int }

func (w Wallet) Balance() int { return w.Cents }      // value
func (w *Wallet) Deposit(n int) { w.Cents += n }      // pointer

func main() {
    w := Wallet{}
    w.Deposit(150)
    fmt.Println(w.Balance()) // 150
}
```

Edge case: value methods are callable on pointers and vice versa (when addressable), because the compiler implicitly takes/dereferences addresses to make method calls valid.

---

## 3) Method Sets and Interface Satisfaction

Method sets determine which methods are available on a type and whether it satisfies an interface.

Rules (simplified):
- The method set of T (non-pointer) includes all methods with receiver type T
- The method set of *T (pointer) includes all methods with receiver type T and *T
- Only *T has methods with pointer receivers in its method set

```go
package main
import "fmt"

type T struct{ n int }

func (t T) Get() int   { return t.n }
func (t *T) Set(v int) { t.n = v }

type Getter interface{ Get() int }
type Setter interface{ Set(int) }

func main() {
    var t T
    var p *T = &t

    var g Getter = t  // ok: T has Get
    _ = g
    // var s Setter = t // compile error: T lacks Set in its method set
    var s Setter = p   // ok: *T has Set
    s.Set(10)
    fmt.Println(p.Get()) // 10
}
```

Takeaway: if any method needs a pointer receiver, prefer using *T consistently when passing to interfaces that expect those methods.

---

## 4) Method Values vs Method Expressions

- Method value: binds the receiver at the time you take the value
- Method expression: does not bind the receiver; you pass it later

```go
package main
import "fmt"

type Counter struct{ n int }
func (c *Counter) Inc() { c.n++ }

func main() {
    c := &Counter{}

    // Method value: receiver bound
    incVal := c.Inc
    incVal()

    // Method expression: pass receiver explicitly
    incExpr := (*Counter).Inc
    incExpr(c)

    fmt.Println(c.n) // 2
}
```

When to use:
- Method values are convenient callbacks when the receiver is known
- Method expressions are useful for higher-order functions that accept operations over many receivers

---

## 5) Embedding and Promoted Methods

Embedding promotes methods into the outer type’s method set (subject to rules of ambiguity).

```go
package main
import "fmt"

type Logger struct{ Prefix string }
func (l Logger) Log(msg string) { fmt.Println(l.Prefix + msg) }

type Service struct {
    Logger // embedded
    Name string
}

func main() {
    s := Service{Logger: Logger{Prefix: "[SVC] "}, Name: "Payments"}
    s.Log("up") // promoted method from Logger
}
```

Notes:
- If multiple embedded types expose the same method, you must qualify explicitly (s.Logger.Log)
- Embedding a *Logger vs Logger changes which methods are promoted onto Service (pointer-receiver methods require *Logger)

---

## 6) Concurrency and Receiver Design

Receiver choice affects concurrency expectations:
- Value receiver methods on small immutable types are naturally safe
- Pointer receiver methods that mutate shared state need external synchronization (e.g., mutex)

```go
package main
import (
  "fmt"
  "sync"
)

type SafeCounter struct {
  mu sync.Mutex
  n  int
}
func (s *SafeCounter) Inc() { s.mu.Lock(); s.n++; s.mu.Unlock() }
func (s *SafeCounter) Get() int { s.mu.Lock(); defer s.mu.Unlock(); return s.n }

func main() {
  s := &SafeCounter{}
  var wg sync.WaitGroup
  for i := 0; i < 10; i++ { wg.Add(1); go func(){ s.Inc(); wg.Done() }() }
  wg.Wait()
  fmt.Println(s.Get()) // 10
}
```

---

## 7) Common Mistakes and Gotchas

1) Expecting T to satisfy interfaces with pointer-receiver methods
```go
// fix: pass *T instead of T when the interface requires pointer-receiver methods
```

2) Unintended copies with value receivers on large structs
```go
// fix: use pointer receivers for large data or when mutation is expected
```

3) Capturing method values of temporary receivers
```go
// fix: ensure the receiver outlives the method value’s use (store pointer)
```

4) Relying on promoted methods with ambiguous embedding
```go
// fix: call explicitly (s.Logger.Log)
```

---

## 8) Best Practices

- Choose receiver type based on semantics (mutate vs read-only) and size
- Be consistent across a type; if any method needs *T, consider making all methods pointer receivers
- Keep methods small and cohesive; defer complex work to helper functions
- Document concurrency expectations (thread-safety) for pointer-receiver methods
- Avoid exposing method sets that make interface satisfaction surprising

---

## 9) Performance Considerations

- Pointer receivers avoid copying large structs; value receivers can enable better cache locality for small types
- Methods may be inlined; small, simple methods often are (inspect with `go build -gcflags="-m"`)
- Method values allocate when captured by closures; consider method expressions when appropriate

---

## 10) Advanced Challenge Questions

1) Why can a value of type T call a method with pointer receiver in many cases?
- Because the compiler can take the address of addressable values to match the method set of *T.

2) How do method sets affect interface satisfaction for T vs *T?
- T’s method set excludes pointer-receiver methods; *T’s includes both. Interfaces requiring pointer-receiver methods are only satisfied by *T.

3) What happens if two embedded types promote the same method name?
- The promotion is ambiguous; you must qualify the call (e.g., s.A.Log vs s.B.Log), or redesign to avoid collision.

4) When do method values vs expressions make a practical difference?
- Method values capture a receiver at creation time (useful for callbacks). Method expressions defer binding and require the receiver at call time (useful for generic pipelines).

5) What concurrency pitfalls arise from pointer-receiver methods?
- Data races if the same receiver is mutated concurrently; guard with sync primitives or design for immutability.

