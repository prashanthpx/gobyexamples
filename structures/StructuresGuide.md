# Go Structs: Advanced Developer Guide

## Table of Contents
1. [What Are Structs and Why Use Them](#toc-1-why-structs)
2. [Composition, Embedding, and Inheritance (not!)](#toc-2-composition-embedding)
3. [Initialization Patterns and Zero Values](#toc-3-initialization)
4. [Methods on Structs (value vs pointer)](#toc-4-methods)
5. [Tags, Encoding/Decoding (JSON/YAML)](#toc-5-tags-encoding)
6. [Memory Layout, Alignment, and Packing](#toc-6-memory-layout)
7. [Equality, Comparability, and Maps/Sets](#toc-7-equality)
8. [Common Mistakes and Gotchas](#toc-8-mistakes)
9. [Best Practices](#toc-9-best)
10. [Advanced Challenge Questions](#toc-10-advanced)


Run these examples
- Tag pitfalls (see notes); runnable JSON example: maps/examples/iterate_order.go shows deterministic iteration pattern; for JSON, rely on your own examples directory.

---

<a id="toc-1-why-structs"></a>

## 1) What Are Structs and Why Use Them

Structs are the primary aggregation type in Go. They:
- Group related fields into a custom type
- Are value types (assignments copy the struct)
- Support methods and interface satisfaction indirectly
- Avoid the pitfalls of class-based inheritance by favoring composition

```go
package main
import "fmt"

type User struct {
    ID   int
    Name string
    Admin bool
}

func main() {
    u := User{ID: 1, Name: "Alice", Admin: true}
    fmt.Printf("%+v\n", u) // {ID:1 Name:Alice Admin:true}
}
```

Why structs matter:
- Clear, explicit data models
- Enable stable API design with zero values that are usable
- Pair naturally with methods, interfaces, and encoding packages

---

<a id="toc-2-composition-embedding"></a>

## 2) Composition, Embedding, and Inheritance (not!)

Go has no inheritance. Prefer composition and embedding.

### Composition (named field)
```go
type Address struct { City, Country string }

type Customer struct {
    User    // embedded (see below)
    Addr Address // composed (named field)
}
```

### Embedding (anonymous field)
- Promotes the embedded type’s methods/fields to the outer type’s method set
- Not inheritance; just syntactic convenience

```go
type Logger struct{ Prefix string }
func (l Logger) Log(msg string) { fmt.Println(l.Prefix + msg) }

type Service struct {
    Logger // embedded
    Name string
}

func main() {
    s := Service{Logger: Logger{Prefix: "[SVC] "}, Name: "Billing"}
    s.Log("started") // promoted: calls Logger.Log
}
```

Rules to remember:
- Method promotion happens at compile time; ambiguity must be resolved explicitly
- Embedding a pointer vs a value affects the method set exposed

---

<a id="toc-3-initialization"></a>

## 3) Initialization Patterns and Zero Values

Zero values in Go should be useful whenever possible.

```go
type Counter struct{ N int }

func (c *Counter) Inc() { c.N++ }

func main() {
    var c Counter // zero value is usable
    c.Inc()
}
```

Initialization patterns:
```go
u1 := User{ID: 1, Name: "Alice"}       // keyed literal (preferred for clarity)
u2 := User{2, "Bob", true}             // field-order literal (fragile)
u3 := new(User)                         // *User with zero values
u4 := &User{ID: 3}                       // pointer to literal
```

Functional options (for complex construction):
```go
type Server struct { Host string; Port int; TLS bool }

type Option func(*Server)

func WithTLS(on bool) Option { return func(s *Server) { s.TLS = on } }

func NewServer(host string, port int, opts ...Option) *Server {
    s := &Server{Host: host, Port: port}
    for _, opt := range opts { opt(s) }
    return s
}

_ = NewServer("localhost", 8080, WithTLS(true))
```

---

<a id="toc-4-methods"></a>

## 4) Methods on Structs (value vs pointer)

```go
type Acc struct { Balance int }

func (a Acc) Snapshot() int { return a.Balance }        // value receiver
func (a *Acc) Deposit(n int) { a.Balance += n }         // pointer receiver
```

Guidelines:
- Use pointer receivers when the method mutates or the struct is large
- Be consistent across the type to avoid surprising interface behavior
- Values can call pointer methods (address-taken) and vice versa when addressable

---

<a id="toc-5-tags-encoding"></a>

## 5) Tags, Encoding/Decoding (JSON/YAML)

Run this example
- omitempty behavior: go run structures/examples/json_omitempty.go

Tags are metadata strings associated with fields.

```go
package main
import (
    "encoding/json"
    "fmt"
)

type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price,omitempty"`
}

func main() {
    p := Product{ID: 10, Name: "Book"}
    b, _ := json.Marshal(p)
    fmt.Println(string(b)) // {"id":10,"name":"Book"}
}
```

Notes:
- `omitempty` skips zero values (e.g., 0 for ints, empty string, nil slices/maps)
- Tag format is a single string; convention is space-separated key:"value"
- Mis-typed tags are ignored silently

Common pitfalls:
```go
// 1) Unexported fields are ignored by encoders
// type user struct { name string `json:"name"` } // name won’t be marshaled

// 2) omitempty hides zero values
// type Item struct { Count int `json:"count,omitempty"` } // 0 is omitted

// 3) Time formats
// Use time.Time with proper layout or custom marshalers for non-RFC3339 formats

// 4) Custom (Un)Marshal
// Implement json.Marshaler / json.Unmarshaler for custom behavior
```

---

<a id="toc-6-memory-layout"></a>

## 6) Memory Layout, Alignment, and Packing

Field order affects padding and size.

```go
package main
import (
  "fmt"
  "unsafe"
)

type A struct { // suboptimal
  A bool
  B int64
  C bool
}

type B struct { // better: minimize padding
  B int64
  A bool
  C bool
}

func main() {
  fmt.Println(unsafe.Sizeof(A{})) // e.g., 24
  fmt.Println(unsafe.Sizeof(B{})) // e.g., 16
}
```

Tips:
- Group larger-alignment fields first
- Avoid unnecessary pointers inside hot structs (GC scanning)

---

<a id="toc-7-equality"></a>

## 7) Equality, Comparability, and Maps/Sets

- Structs are comparable if all fields are comparable
- Use struct keys in maps to avoid pointer identity bugs

```go
type Key struct { A, B int }
m := map[Key]string{}
m[Key{1,2}] = "x"
_, ok := m[Key{1,2}] // ok is true
```

Not comparable:
- Any field that is a slice, map, or function makes the struct non-comparable

---

<a id="toc-8-mistakes"></a>

## 8) Common Mistakes and Gotchas

- Relying on field-order literals: breaks when fields rearranged
- Embedding confusion: promoted names clash; specify explicitly (s.Logger.Log())
- JSON tags with wrong keys/typos: silently ignored
- Using pointers for tiny structs: worse cache locality and GC pressure
- Mutating time.Time by value (it’s a struct): copy vs pointer semantics surprise

Practice-ready explanation:
- “Structs are values; copying occurs on assignment/pass. Pointer receivers are chosen for mutation and performance, but I prefer consistent receiver types to satisfy interfaces predictably. For encoding, I use keyed literals and careful tag validation. I order fields to reduce padding and avoid unnecessary heap pointers in hot paths.”

---

<a id="toc-9-best"></a>

## 9) Best Practices

- Prefer keyed literals for clarity and stability
- Choose pointer receivers consistently when any method needs it
- Validate and lint struct tags (json, yaml, db)
- Optimize field order in hot structs; avoid pointer-rich designs in hot paths
- Favor composition/embedding over inheritance; avoid ambiguous promotion

---

<a id="toc-10-advanced"></a>

## 10) Advanced Challenge Questions

1) How does embedding differ from inheritance?
- Embedding promotes methods/fields; there is no polymorphic substitution like subclassing. It’s composition with syntactic sugar.

2) Are structs comparable? When not?
- Comparable if all fields are comparable; not if any field is a slice/map/func.

3) Why reorder fields? What’s the effect?
- Reduce padding due to alignment; improves memory footprint and cache behavior.

4) How do tags work? What’s `omitempty`?
- Tags are strings parsed by packages like encoding/json. `omitempty` omits zero-value fields during marshaling.

5) Value vs pointer receivers — how do they affect interface satisfaction?
- Method sets differ; only *T has pointer receiver methods. If an interface includes a pointer-receiver method, only *T implements it.

