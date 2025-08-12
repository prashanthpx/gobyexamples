# Go Operators: Advanced Developer Guide

## Table of Contents
1. Operator Overview and Precedence
2. Arithmetic, Assignment, and Increment/Decrement
3. Comparison and Equality (values vs references)
4. Logical Operators and Short-Circuiting
5. Bitwise and Shift Operators (with signed/unsigned notes)
6. Address-of, Dereference, Indexing, Slicing
7. Map/Channel Operators (comma-ok, send/recv)
8. Type Conversion vs Casting (there’s no cast)
9. Comparability Rules (structs, arrays, slices, maps)
10. Common Mistakes and Gotchas
11. Best Practices
12. Advanced Challenge Questions

---

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

## 7) Map/Channel Operators (comma-ok, send/recv)

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

## 9) Comparability Rules (structs, arrays, slices, maps)

- Comparable: booleans, numbers, strings, pointers, channels, arrays, structs (if all fields comparable)
- Not comparable: slices, maps, functions; interfaces may be comparable depending on dynamic value
- Use `reflect.DeepEqual` only for tests/tools; not on hot paths

---

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

## 11) Best Practices

- Use parentheses for clarity in complex expressions
- Prefer unsigned types for bit manipulation
- Be explicit with conversions; avoid unsafe unless absolutely required
- Compare composite types by value semantics where possible; design types to be comparable when useful
- Avoid DeepEqual in production logic; write domain-specific equality

---

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

