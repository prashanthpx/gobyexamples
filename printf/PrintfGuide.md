# Formatting and Printing in Go (fmt): Advanced Developer Guide

## Table of Contents
1. Print Families (Print, Fprint, Sprint)
2. Verbs and Flags Overview
3. Formatting Numbers (int, uint, float, complex)
4. Strings, Bytes, Runes
5. Structs, Maps, Slices, Pointers
6. Width, Precision, and Alignment
7. Custom Formatting (fmt.Stringer, fmt.GoStringer, fmt.Formatter)
8. Errors and %w Wrapping
9. Logging vs fmt: Performance Notes
10. Common Mistakes and Gotchas
11. Best Practices
12. Advanced Challenge Questions


Run these examples
- Vet format strings: see section below; try the fmtbad.go snippet

---

## 1) Print Families (Print, Fprint, Sprint)

- Print… writes to standard output
- Fprint… writes to an io.Writer (files, buffers, sockets)
- Sprint… returns a string (allocates)

```go
package main
import (
  "bytes"
  "fmt"
  "os"
)

func main() {
  fmt.Println("hello", 42)           // Println -> stdout + newline
  fmt.Printf("%s %d\n", "hello", 42) // Printf -> formatted to stdout

  var buf bytes.Buffer
  fmt.Fprintf(&buf, "id=%d", 7)     // to any io.Writer
  s := fmt.Sprintf("%q", "go")       // to string
  os.Stdout.Write(buf.Bytes())
  fmt.Println(" ", s)
}
```

---

## 2) Verbs and Flags Overview

- Generic: %v (value), %+v (include field names), %#v (Go-syntax), %T (type)
- Booleans: %t
- Integers: %d (decimal), %b (binary), %o (octal), %x/%X (hex), %c (rune)
- Floats/complex: %f, %e/%E, %g/%G
- Strings/bytes: %s, %q (quoted), %x (hex)
- Pointers: %p

Flags:
- + (always show sign), space (pad sign), 0 (zero-pad), - (left align), # (alternate form)

```go
fmt.Printf("%#v %T\n", []int{1,2,3}, []int{}) // []int{1, 2, 3} []int
fmt.Printf("%+v\n", struct{A int; B string}{1, "x"}) // {A:1 B:x}
fmt.Printf("%#x %#X\n", 255, 255) // 0xff 0XFF (with # alternate form)
```

---

## 3) Formatting Numbers (int, uint, float, complex)

```go
package main
import "fmt"

func main() {
  n := -42
  fmt.Printf("%d %b %o %x %X\n", n, n, n, n, n)
  fmt.Printf("%+d % d %06d\n", n, n, 7) // sign, space-sign, zero-pad width 6

  f := 3.14159265
  fmt.Printf("%.2f %.4g %e\n", f, f, f)

  c := 2 + 3i
  fmt.Printf("%v %f\n", c, c) // %f prints real and imag as "(a+bi)"
}
```

Notes:
- For floats: %.N sets precision; %g chooses compact form
- For complex: %f formats as (real+imagi)

---

## 4) Strings, Bytes, Runes

```go
package main
import "fmt"

func main() {
  s := "Go\n"
  fmt.Printf("%s %q %x\n", s, s, s) // plain, quoted (escapes), hex

  b := []byte("Hi")
  r := 'A' // rune
  fmt.Printf("bytes:% x rune:%c (%U)\n", b, r, r) // spaced-hex, char, Unicode
}
```

Notes:
- %q for strings/runes prints Go-escaped representation
- % x inserts spaces between bytes when formatting slices

---

## 5) Structs, Maps, Slices, Pointers

```go
package main
import "fmt"

type User struct { ID int; Name string }

func main() {
  u := User{ID:1, Name:"Alice"}
  fmt.Printf("%v\n%+v\n%#v\n", u, u, u)
  m := map[string]int{"a":1}
  p := &u
  fmt.Printf("map:%v ptr:%p type:%T\n", m, p, p)
}
```

Notes:
- %+v includes field names; %#v prints a Go-syntax literal (handy for debugging)
- %p prints the pointer value (address)

---

## 6) Width, Precision, and Alignment

`%[index][flags][width][.prec][verb]`

```go
fmt.Printf("|%8s|%-8s|\n", "right", "left")    // pad to width 8
fmt.Printf("|%08d|%.3f|\n", 42, 3.1)          // zero-pad; precision for floats
fmt.Printf("|%10.10s|\n", "longlonglong")      // precision truncates strings

// Argument indexing (useful in i18n templates)
fmt.Printf("%[2]d %[1]s\n", "idx1", 2) // prints: 2 idx1
```

---

## 7) Custom Formatting (fmt.Stringer, fmt.GoStringer, fmt.Formatter)

Implement `String() string` to control %s/%v.

```go
package main
import "fmt"

type Point struct{ X, Y int }
func (p Point) String() string { return fmt.Sprintf("(%d,%d)", p.X, p.Y) }

func main() {
  fmt.Println(Point{1,2}) // (1,2)
}
```

Implement `GoString() string` to control %#v.

`fmt.Formatter` gives full control over verbs/flags:
```go
type Money int
func (m Money) Format(f fmt.State, c rune) {
  switch c {
  case 'v', 's':
    fmt.Fprintf(f, "$%d.%02d", int(m)/100, int(m)%100)
  default:
    fmt.Fprintf(f, "%%!%c(Money=%d)", c, int(m))
  }
}
```

---

## 8) Errors and %w Wrapping

Use `%w` in `fmt.Errorf` to wrap errors (Go 1.13+). Unwrap with `errors.Is/As`.

```go
package main
import (
  "errors"
  "fmt"
  "os"
)

func main() {
  _, err := os.Open("missing.txt")
  if err != nil {
    err = fmt.Errorf("open failed: %w", err)
    if errors.Is(err, os.ErrNotExist) {
      fmt.Println("not found")
    }
  }
}
```

---

## 9) Logging vs fmt: Performance Notes

- fmt.Sprint/Sprintf allocate strings; prefer io.Writer (Fprint) to stream
- For hot paths, prefer bufio.Writer or bytes.Buffer over repeated fmt
- Avoid fmt in tight loops when constructing simple text; use WriteString
- `fmt` is fine for most CLIs and tutorials; profile before optimizing

### Vetting format strings
```bash
# go vet flags type/arg mismatches at compile time
cat > fmtbad.go <<'EOF'
package main
import "fmt"
func main(){ fmt.Printf("%d", "x") }
EOF
go vet fmtbad.go # reports: Printf format %d has arg "x" of wrong type string
```

---

## 10) Common Mistakes and Gotchas

1) Mismatched verbs and arguments
```go
fmt.Printf("%d", "x") // output: %!d(string=x)
fmt.Printf("%s", 10)  // output: %!s(int=10)
fmt.Printf("%d %d", 1) // output: %!d(MISSING)
```

2) Forgetting newlines with Printf/Print
```go
fmt.Printf("hello") // no newline; use \n or Println
```

3) Assuming maps print in key order
```go
fmt.Printf("%v", map[int]int{2:2,1:1}) // order not guaranteed
```

4) Using %v for user-facing strings when exact format is required
```go
// Prefer explicit verbs and width/precision for stable output
```

5) Printing large data structures with %#v in production logs
```go
// Expensive and verbose; use targeted fields or custom String/Formatter
```

---

## 11) Best Practices

- Use Fprint family to decouple formatting from destinations
- Implement String() for types with natural textual form
- Use `%q` for strings in logs to make whitespace visible
- Prefer explicit width/precision for tabular outputs
- Use `%T` during debugging to confirm dynamic types
- Wrap errors with `%w` to preserve cause chain

---

## 12) Advanced Challenge Questions

1) When do you use `%#v` vs `%+v`?
- `%#v` prints a Go-syntax representation (useful for recreating values). `%+v` includes field names for structs.

2) How do you implement custom formatting for a type across multiple verbs?
- Implement `fmt.Formatter` and branch on the rune verb to support `%v`, `%s`, etc.

3) What are the performance implications of Sprintf vs Fprintf?
- Sprintf allocates a new string; Fprintf writes to a provided Writer, avoiding intermediate allocation.

4) What does `%w` do in `fmt.Errorf` and how do you unwrap?
- Wraps an error for later inspection. Use `errors.Is/As` to check/unwrap.

5) Why might you see output like `%!d(string=x)` or `%!d(MISSING)`?
- Verb/type mismatch or insufficient arguments; go vet can catch many of these.

