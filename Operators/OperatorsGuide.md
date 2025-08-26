# Go Operators: Advanced Developer Guide

## Table of Contents
1. [Operator Overview and Precedence](#toc-1-overview)
2. [Arithmetic, Assignment, and Increment/Decrement](#toc-2-arithmetic)
3. [Comparison and Equality (values vs references)](#toc-3-comparison)
4. [Logical Operators and Short-Circuiting](#toc-4-logical)
5. [Bitwise and Shift Operators (with signed/unsigned notes)](#toc-5-bitwise)
6. [Address-of, Dereference, Indexing, Slicing](#toc-6-address-indexing)
7. [Map/Channel Operators (comma-ok, send/recv)](#toc-7-map-channel)
8. [Type Conversion vs Casting (there’s no cast)](#toc-8-conversion)
9. [Comparability Rules (structs, arrays, slices, maps)](#toc-9-comparability)
10. [Common Mistakes and Gotchas](#toc-10-mistakes)
11. [Best Practices](#toc-11-best)
12. [Advanced Challenge Questions](#toc-12-advanced)

---

<a id="toc-1-overview"></a>

## 1) Operator Overview and Precedence

From high to low (simplified):
- Unary: `+ - ! ^ * & <-`
- Multiplicative: `* / % << >> & &^`
- Additive: `+ - | ^`
- Comparison: `== != < <= > >=`
- Logical AND: `&&`
- Logical OR: `||`

Parentheses `()` override precedence. Use them for clarity.

---

<a id="toc-2-arithmetic"></a>

## 2) Arithmetic, Assignment, and Increment/Decrement

```go
package main
import "fmt"

func main() {
  a, b := 7, 3
  fmt.Println(a+b, a-b, a*b, a/b, a%b) // 10 4 21 2 1

  a += 2; b *= 5
  fmt.Println(a, b) // 9 15

  // ++ and -- are statements, not expressions
  a++ // ok
  // x := a++ // compile error
}
```

Notes:
- Integer division truncates toward zero
- `%` works for integers; for floats use math.Mod
- `++` and `--` can’t be used in expressions

---

<a id="toc-3-comparison"></a>

## 3) Comparison and Equality (values vs references)

```go
package main
import (
  "fmt"
  "reflect"
)

type P struct{ X, Y int }

func main() {
  // Values
  fmt.Println(3 == 3.0)       // true (untyped const rules)
  fmt.Println(P{1,2} == P{1,2}) // true (all fields comparable)

  // Slices and maps are not comparable (except to nil)
  var s1, s2 []int
  fmt.Println(s1 == nil) // true
  // fmt.Println(s1 == s2) // compile error
  fmt.Println(reflect.DeepEqual([]int{1}, []int{1})) // true (slow path)
}
```

Notes:
- Arrays and structs comparable if all elements/fields are comparable
- Slices, maps, and functions are not comparable (except with nil)
- Interface equality compares dynamic type and value

---

<a id="toc-4-logical"></a>

## 4) Logical Operators and Short-Circuiting

```go
package main
import "fmt"

func A() bool { fmt.Print("A "); return false }
func B() bool { fmt.Print("B "); return true }

func main() {
  if A() && B() { fmt.Println("X") } else { fmt.Println("Y") } // A Y
  if A() || B() { fmt.Println("X") } else { fmt.Println("Y") } // A B X
}
```

Use short-circuiting to guard expensive or unsafe operations.

---

<a id="toc-5-bitwise"></a>

## 5) Bitwise and Shift Operators (with signed/unsigned notes)

```go
package main
import "fmt"

func main() {
  var x uint8 = 0b1010 // 10
  var y uint8 = 0b1100 // 12
  fmt.Printf("%08b\n", x&y)   // 00001000
  fmt.Printf("%08b\n", x|y)   // 00001110
  fmt.Printf("%08b\n", x^y)   // 00000110
  fmt.Printf("%08b\n", x&^y)  // 00000010 (AND NOT)

  fmt.Println(1<<3, 16>>2) // 8 4

  // Signed right shift keeps sign bit (implementation-defined as arithmetic shift in Go)
  var n int8 = -8
  fmt.Println(n>>1) // -4 on two's complement
}
```

Notes:
- Shift count’s type is not restricted; negative shift is a compile error
- Use unsigned for pure bit patterns; signed right shift preserves sign

---

<a id="toc-6-address-indexing"></a>

## 6) Address-of, Dereference, Indexing, Slicing

```go
package main
import "fmt"

func main() {
  v := 10
  p := &v      // address-of
  *p = 20      // dereference write
  fmt.Println(v) // 20

  a := [5]int{0,1,2,3,4}
  s := a[1:4]      // [1 2 3]
  s = s[:cap(s)]   // extend within capacity
  fmt.Println(s)
}
```

---

<a id="toc-7-map-channel"></a>

## 7) Map/Channel Operators (comma-ok, send/recv)

### Map Declaration and Initialization

#### 1. Declaring with just the type
```go
var m map[string]string
```

This declares a nil map variable.
- `m` is of type `map[string]string`, but the value is `nil`
- You can read from it safely (returns the zero value for the value type)
- But you cannot write (insert/update), it will panic:

```go
fmt.Println(m == nil)       // true
fmt.Println(m["foo"])       // "" (zero value of string)
m["foo"] = "bar"            // panic: assignment to entry in nil map
```

#### 2. Declaring with make
```go
m := make(map[string]string)
```

This creates and initializes an empty but usable map.

Now you can both read and write:
```go
fmt.Println(m == nil)   // false
m["foo"] = "bar"        // works fine
fmt.Println(m["foo"])   // "bar"
```

#### 3. Why the difference?

A map type (`map[string]string`) is a reference type.
- Just declaring `var m map[...]...` gives you a nil reference
- You need `make` (or a literal like `map[string]string{"foo":"bar"}`) to allocate the underlying map header and buckets

This is very much like slices:
- `var s []int` → nil slice (len=0, cap=0)
- `make([]int, 0)` → empty slice with underlying array allocated (len=0, cap=0 but usable)

#### 4. Quick comparison table

| Form | Example | Usable for reads? | Usable for writes? | Is nil? |
|------|---------|-------------------|-------------------|---------|
| `var m map[K]V` | `var m map[string]string` | ✅ returns zero value | ❌ panics | ✅ |
| `make(map[K]V)` | `make(map[string]string)` | ✅ | ✅ | ❌ |
| Map literal | `map[string]string{"k":"v"}` | ✅ | ✅ | ❌ |

### Channel Buffering

#### make(chan int, N) in Go

The first argument is the channel's element type (`int` here).
The second argument is the channel's buffer capacity (`N`).

So:
- `make(chan int)` → creates an unbuffered channel (capacity = 0)
- `make(chan int, 1)` → creates a buffered channel with capacity = 1
- `make(chan int, 2)` → creates a buffered channel with capacity = 2

#### What buffering means

- Capacity = 1 → at most one value can sit in the channel without being received
- Capacity = 2 → at most two values can sit in the channel before a sender blocks

Example:
```go
ch := make(chan int, 2)

ch <- 10   // ✅ doesn't block
ch <- 20   // ✅ doesn't block

// At this point, the buffer is full.
// Next send would block until someone receives.
go func() {
    ch <- 30 // this will block until a receiver takes one out
}()
fmt.Println(<-ch) // 10
fmt.Println(<-ch) // 20
fmt.Println(<-ch) // 30
```

#### Key rules

- `cap(ch)` → returns the buffer capacity
- `len(ch)` → returns the current number of elements waiting in the buffer
- If the buffer is full, a sender blocks until space is freed
- If the buffer is empty, a receiver blocks until a value is sent

### Map and Channel Operations

```go
package main
import "fmt"

func main() {
  m := map[string]int{"zero":0}
  if v, ok := m["zero"]; ok { fmt.Println(v) }
  if _, ok := m["missing"]; !ok { fmt.Println("no key") }

  ch := make(chan int, 1)
  ch <- 42              // send
  if v, ok := <-ch; ok { fmt.Println(v) } // recv, ok=true before close
  close(ch)
  v, ok := <-ch         // zero, false after close
  fmt.Println(v, ok)
}
```

---

<a id="toc-8-conversion"></a>

## 8) Type Conversion vs Casting (there’s no cast)

Go has explicit conversions, not C-style casts.

```go
package main
import "fmt"

func main() {
  var i int = 65
  var b byte = byte(i)   // conversion
  s := string(b)         // "A"

  f := float64(i)        // 65.0
  // var p *int = (*int)(&f) // compile error: unrelated types
}
```

Notes:
- Conversions change the representation where defined; not all types are convertible
- Use unsafe.Pointer for low-level reinterpretation (avoid unless necessary)

---

<a id="toc-9-comparability"></a>

## 9) Comparability Rules (structs, arrays, slices, maps)

- Comparable: booleans, numbers, strings, pointers, channels, arrays, structs (if all fields comparable)
- Not comparable: slices, maps, functions; interfaces may be comparable depending on dynamic value
- Use `reflect.DeepEqual` only for tests/tools; not on hot paths

---

<a id="toc-10-mistakes"></a>

## 10) Common Mistakes and Gotchas

1) Using ++/-- in expressions
```go
// ❌ x := a++
// ✅ a++; x := a
```

2) Expecting slice/map equality with ==
```go
// ❌ s1 == s2 // compile error
// ✅ use DeepEqual or compare elements/keys directly
```

3) Shifts with negative counts or mixing signed/unsigned unintentionally
```go
// ❌ n << -1 // compile error
```

4) Assuming interface == compares pointers only
```go
// It compares dynamic type + value; pointers compare by address values.
```

5) Confusing XOR ^ with bit clear &^
```go
// ^ is XOR (binary) or bitwise NOT (unary); &^ is AND NOT
```

---

<a id="toc-11-best"></a>

## 11) Best Practices

- Use parentheses for clarity in complex expressions
- Prefer unsigned types for bit manipulation
- Be explicit with conversions; avoid unsafe unless absolutely required
- Compare composite types by value semantics where possible; design types to be comparable when useful
- Avoid DeepEqual in production logic; write domain-specific equality

---

<a id="toc-12-advanced"></a>

## 12) Advanced Challenge Questions

1) Why are slices not comparable, and how do you compare them?
- They are headers pointing to backing arrays; equality is not well-defined for content. Compare lengths and elements, or use DeepEqual for tests.

2) What does `&^` do and when is it useful?
- Bit clear (AND NOT): clears bits present in the right operand. Useful in masks.

3) Why can’t you convert `*float64` to `*int` directly?
- Pointer conversions require identical underlying types (barring unsafe). The memory layout differs; use value conversion via an intermediate.

4) How does interface equality work for two interface variables?
- Equal only if dynamic types are identical and dynamic values compare equal.

5) What’s the difference between unary `^x` and binary `x ^ y`?
- Unary is bitwise NOT (flip bits); binary is XOR.

