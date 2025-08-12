# Conditionals in Go: Advanced Developer Guide

## Table of Contents
1. if Statements and Short Declarations
2. switch (Values, Expressions, No-Condition)
3. Type Switches
4. Labeled break/continue and fallthrough
5. Boolean Logic and Short-Circuiting
6. Error Handling Patterns with Conditionals
7. Common Mistakes and Gotchas
8. Best Practices
9. Performance Considerations
10. Advanced Challenge Questions

---

## 1) if Statements and Short Declarations

Go’s `if` supports an optional short statement, scoped to the `if` and `else` blocks. This is idiomatic for error checks and temporary variables.

```go
package main
import (
  "errors"
  "fmt"
)

func div(a, b int) (int, error) {
  if b == 0 { return 0, errors.New("division by zero") }
  return a / b, nil
}

func main() {
  if q, err := div(10, 2); err != nil {
    fmt.Println("err:", err)
  } else {
    fmt.Println("quotient:", q)
  }
  // q, err are not in scope here
}
```

Notes:
- The short statement executes before the condition and is in scope for both if and else blocks only
- Prefer early-returns on errors for clarity; use short `if` when you need values in both arms

---

## 2) switch (Values, Expressions, No-Condition)

switch compares a value against cases; it can also switch on arbitrary expressions, or omit a condition entirely (switch true).

```go
package main
import "fmt"

func main() {
  // Value switch
  day := 3
  switch day {
  case 1, 7:
    fmt.Println("weekend")
  case 2, 3, 4, 5, 6:
    fmt.Println("weekday")
  default:
    fmt.Println("unknown")
  }

  // Expression switch
  n := 42
  switch {
  case n%2 == 0 && n%3 == 0:
    fmt.Println("divisible by 6")
  case n%2 == 0:
    fmt.Println("even")
  case n%3 == 0:
    fmt.Println("divisible by 3")
  default:
    fmt.Println("other")
  }
}
```

Notes:
- Cases evaluate top‑down; the first matching case runs
- No implicit fallthrough (unlike C). Use `fallthrough` explicitly if needed

---

## 3) Type Switches

A type switch branches on the dynamic type of an interface value. It does not perform allocations.

```go
package main
import (
  "fmt"
)

func describe(i interface{}) {
  switch v := i.(type) {
  case nil:
    fmt.Println("nil")
  case int:
    fmt.Println("int:", v)
  case string:
    fmt.Println("string:", v)
  case fmt.Stringer:
    fmt.Println("stringer:", v.String())
  default:
    fmt.Printf("unknown %T\n", v)
  }
}

func main() {
  describe(nil)
  describe(10)
  describe("go")
}
```

Notes:
- The `v` inside a type switch case is the asserted concrete type
- The zero case `case nil` is valid and often useful

---

## 4) Labeled break/continue and fallthrough

Use labels to break/continue outer loops unambiguously. Use `fallthrough` sparingly to chain cases.

```go
package main
import "fmt"

func main() {
Outer:
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      if i*j > 2 { fmt.Println("break outer"); break Outer }
      if i == j { fmt.Println("continue outer"); continue Outer }
      fmt.Println("i,j:", i, j)
    }
  }

  // fallthrough example
  x := 2
  switch x {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("two")
    fallthrough
  case 3:
    fmt.Println("three or after two")
  default:
    fmt.Println("other")
  }
}
```

Notes:
- `fallthrough` does not re-check the next case condition; it blindly continues to the next case’s body

---

## 5) Boolean Logic and Short-Circuiting

Go uses short-circuit evaluation for `&&` and `||`.

```go
package main
import "fmt"

func A() bool { fmt.Print("A "); return false }
func B() bool { fmt.Print("B "); return true }

func main() {
  if A() && B() { fmt.Println("X") } else { fmt.Println("Y") } // prints: A Y
  if A() || B() { fmt.Println("X") } else { fmt.Println("Y") } // prints: A B X
}
```

Notes:
- Use short-circuit behavior to guard nil checks or expensive calls

---

## 6) Error Handling Patterns with Conditionals

Idiomatic early return:
```go
if err := do(); err != nil { return err }
```

Wrapping with context:
```go
if err := do(); err != nil { return fmt.Errorf("do failed: %w", err) }
```

Branching on sentinel/typed errors (Go 1.13+):
```go
if errors.Is(err, os.ErrNotExist) { /* handle not-exist */ }
var e *fs.PathError
if errors.As(err, &e) { /* use fields on e */ }
```

---

## 7) Common Mistakes and Gotchas

1) Shadowing with short declarations
```go
// ❌ Shadows outer err; later checks see the wrong variable
if data, err := load(); err != nil { return err }
// ...
if err != nil { /* outer err here may be stale/unused */ }

// ✅ Name inner vars differently or limit scope tightly
if d, err := load(); err != nil { return err } else { _ = d }
```

2) Missing braces around if bodies (Go requires braces)
```go
// ❌ Not allowed in Go (unlike some C variants)
// if cond
//     stmt
```

3) Misusing fallthrough
```go
// ❌ fallthrough skips re-checking conditions; can create subtle bugs
// ✅ Prefer distinct cases or nested switch/if
```

4) Comparing errors directly when wrapping is used
```go
// ❌ err == os.ErrNotExist fails when err is wrapped
// ✅ use errors.Is/As
```

5) Type assertions without ok-check
```go
// ❌ panics on unexpected type
s := i.(string)
// ✅
if s, ok := i.(string); ok { _ = s } else { /* handle */ }
```

---

## 8) Best Practices

- Prefer early returns over deep nesting
- Use if short statements for narrow-scope temps (e.g., parsing, err)
- Use `switch {}` for multi-branch boolean conditions; it reads better
- Avoid `fallthrough` unless deliberately modeling tiered behavior
- Keep condition expressions side-effect free and testable
- Use errors.Is/As for robust error decisions

---

## 9) Performance Considerations

- Branch prediction: keep hot, likely branches first for readability (the compiler may not reorder)
- Avoid allocating in conditions; hoist work out of conditionals where possible
- Type switches and assertions are efficient; they do not allocate by themselves

---

## 10) Advanced Challenge Questions

1) How does the `if short statement` scope work across if/else blocks?
- The identifiers exist in both arms but not outside the entire if/else statement.

2) When would you use a `switch` with no condition?
- When you want a tidy multi-branch boolean decision tree; it avoids repeated `case true` or nested ifs.

3) What’s the behavior of `fallthrough` in Go?
- It executes the next case’s body unconditionally without re-checking its condition.

4) How do `errors.Is` and `errors.As` improve conditional error handling?
- They work with wrapped errors and let you branch on sentinel or concrete error types reliably.

5) What pitfalls exist with variable shadowing in conditionals, and how to avoid them?
- Short declarations can shadow outer vars (commonly `err`). Use different names, narrower scopes, or avoid short declaration when reusing an outer variable.

