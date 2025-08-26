# Validation in Go: Advanced Developer Guide

## Table of Contents
1. [Validation Philosophy (where and how)](#toc-1-philosophy)
2. [Validating Structs (standard library only)](#toc-2-structs)
3. [Accumulating Multiple Errors](#toc-3-multi-errors)
4. [Contextual Validation (cross-field, external checks)](#toc-4-contextual)
5. [Custom Types and Value Objects](#toc-5-value-objects)
6. [JSON/YAML Tag Interactions and Zero Values](#toc-6-tags-zero)
7. [Common Mistakes and Gotchas](#toc-7-mistakes)
8. [Best Practices](#toc-8-best)
9. [Performance Considerations](#toc-9-performance)
10. [Advanced Challenge Questions](#toc-10-advanced)

---

<a id="toc-1-philosophy"></a>

## 1) Validation Philosophy (where and how)

- Keep validation close to the data model or at boundaries (API/CLI/Config)
- Return specific, actionable errors; avoid generic "invalid input"
- Prefer explicit functions or methods; avoid hidden validation in init()
- For batch validation, accumulate errors rather than failing fast

---

<a id="toc-2-structs"></a>

## 2) Validating Structs (standard library only)

You don’t need a third‑party library for many cases. Start simple and explicit.

```go
package main
import (
  "errors"
  "fmt"
  "strings"
)

type User struct {
  Name  string
  Email string
  Age   int
}

func (u User) Validate() error {
  if strings.TrimSpace(u.Name) == "" {
    return errors.New("name is required")
  }
  if !strings.Contains(u.Email, "@") {
    return errors.New("email must contain @")
  }
  if u.Age < 0 || u.Age > 150 {
    return errors.New("age out of range")
  }
  return nil
}

func main(){
  u := User{Name:"", Email:"x", Age:-1}
  fmt.Println(u.Validate())
}
```

---

<a id="toc-3-multi-errors"></a>

## 3) Accumulating Multiple Errors

Batch callers appreciate getting all issues at once.

```go
type MultiError []error
func (m MultiError) Error() string {
  var b strings.Builder
  for i, e := range m { if i>0 { b.WriteString("; ") }; b.WriteString(e.Error()) }
  return b.String()
}

func (u User) ValidateAll() error {
  var errs MultiError
  if strings.TrimSpace(u.Name) == "" { errs = append(errs, errors.New("name required")) }
  if !strings.Contains(u.Email, "@") { errs = append(errs, errors.New("email invalid")) }
  if u.Age < 0 || u.Age > 150 { errs = append(errs, errors.New("age out of range")) }
  if len(errs) > 0 { return errs }
  return nil
}
```

Tip: Consider returning a typed error that exposes fields for programmatic handling.

---

<a id="toc-4-contextual"></a>

## 4) Contextual Validation (cross-field, external checks)

Some validations need context (e.g., uniqueness, cross-field coherence).

```go
type Registration struct { User User; Password string; Confirm string }

type UserStore interface{ Exists(email string) (bool, error) }

func (r Registration) ValidateWith(ctx context.Context, store UserStore) error {
  if err := r.User.Validate(); err != nil { return err }
  if r.Password == "" || len(r.Password) < 8 { return errors.New("password too short") }
  if r.Password != r.Confirm { return errors.New("passwords do not match") }
  ok, err := store.Exists(r.User.Email)
  if err != nil { return fmt.Errorf("lookup: %w", err) }
  if ok { return errors.New("email already registered") }
  return nil
}
```

Guidelines:
- Accept context.Context as the first parameter when making I/O or long checks
- Keep cross-field logic cohesive and testable

---

<a id="toc-5-value-objects"></a>

## 5) Custom Types and Value Objects

Wrap primitives to encode invariants and centralize validation.

```go
type Email string

func NewEmail(s string) (Email, error) {
  s = strings.TrimSpace(s)
  if !strings.Contains(s, "@") { return "", errors.New("invalid email") }
  return Email(s), nil
}

type Age int
func NewAge(n int) (Age, error) {
  if n < 0 || n > 150 { return 0, errors.New("age out of range") }
  return Age(n), nil
}
```

Advantages:
- Invariants enforced at construction
- Functions can accept Email/Age types and assume validity

---

<a id="toc-6-tags-zero"></a>

## 6) JSON/YAML Tag Interactions and Zero Values

Tags affect marshaling, not validation — but they influence what you receive.

```go
type Item struct {
  Name  string   `json:"name" yaml:"name"`
  Count int      `json:"count,omitempty" yaml:"count,omitempty"`
  Tags  []string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
```

Notes:
- omitempty hides zero values (Count=0 won’t appear); your validator must define whether 0 is allowed
- Unexported fields are ignored by encoders/decoders; validate exported inputs only
- For YAML/JSON, unknown fields may be accepted unless you opt into strict decoding

---

<a id="toc-7-mistakes"></a>

## 7) Common Mistakes and Gotchas

1) Only validating at the edge, not the model
```go
// Put Validate() on the type so all code paths can use it
```

2) Returning generic errors
```go
// Provide actionable messages; prefer typed errors for programmatic checks
```

3) Coupling validation with global singletons
```go
// Pass dependencies (stores, clocks) explicitly; avoid package-level state
```

4) Ignoring context cancellation in external checks
```go
// Accept ctx and honor ctx.Done() in validators that call I/O
```

5) Over-relying on regex for complex formats
```go
// Prefer proper parsers for emails/URLs (net/mail, net/url)
```

---

<a id="toc-8-best"></a>

## 8) Best Practices

- Keep validation deterministic and side-effect free (except deliberate existence checks)
- Co-locate validation with data types; expose Validate() or constructor functions
- Accumulate errors for batch inputs; return first error for request paths as appropriate
- Prefer typed errors with fields or helper functions (Is/As) for matching
- Make validation messages clear, consistent, and user-friendly

---

<a id="toc-9-performance"></a>

## 9) Performance Considerations

- Validation is usually I/O-bound (lookups); micro-optimizations rarely needed
- Avoid repeated allocations in hot paths; reuse buffers/builders for message composition
- Cache heavy metadata (e.g., compiled regex) in package-level vars

---

<a id="toc-10-advanced"></a>

## 10) Advanced Challenge Questions

1) Where should validation live in a layered architecture, and why?
- Close to the domain types to ensure invariants everywhere, with adapters validating at boundaries too.

2) How do tags like `omitempty` affect validation decisions?
- They change what fields arrive; validators must define required vs optional and acceptable zero values.

3) When would you accumulate errors vs fail fast?
- Accumulate for batch/admin tools; fail fast for request paths where user iteration is immediate.

4) How do you validate user-visible text with Unicode considerations?
- Normalize if needed, consider rune counts rather than byte length, and avoid invalid UTF‑8.

