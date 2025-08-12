# Error Handling in Go: Advanced Developer Guide

## Table of Contents
1. Return Values, Not Exceptions
2. Wrapping with %w and Unwrapping with errors.Is/As
3. Sentinel vs Typed Errors
4. Typed-nil Interfaces and Pitfalls
5. Error Values and Context (fmt.Errorf, %w)
6. Common Mistakes and Gotchas
7. Best Practices
8. Performance Considerations
9. Advanced Challenge Questions


Run these examples
- Wrap + Is/As: go run testerrors/examples/wrap_is_as.go

---

## 1) Return Values, Not Exceptions

Go uses explicit error return values. Favor early returns and clear semantics.

```go
func readConfig(path string) (*Config, error) {
  b, err := os.ReadFile(path)
  if err != nil { return nil, fmt.Errorf("read: %w", err) }
  var c Config
  if err := json.Unmarshal(b, &c); err != nil { return nil, fmt.Errorf("parse: %w", err) }
  return &c, nil
}
```

---

## 2) Wrapping with %w and Unwrapping with errors.Is/As

```go
_, err := os.Open("missing.txt")
if err != nil {
  err = fmt.Errorf("open failed: %w", err)
  if errors.Is(err, os.ErrNotExist) {
    // handle not found
  }
}

var pErr *fs.PathError
if errors.As(err, &pErr) {
  // inspect pErr.Path, pErr.Op
}
```

Notes:
- Use %w exactly once in fmt.Errorf to wrap; multiple %w are invalid
- errors.Is walks the chain to match sentinels; As finds the first assignable target

---

## 3) Sentinel vs Typed Errors

- Sentinel: package-level var (e.g., io.EOF). Simple but couples callers to exact values
- Typed: custom types (struct or alias) with extra fields; more flexible

```go
var ErrInvalid = errors.New("invalid")

type ErrCode struct {
  Code int
  Msg  string
}
func (e *ErrCode) Error() string { return fmt.Sprintf("%d: %s", e.Code, e.Msg) }
```

Guidance:
- Prefer typed errors for rich context; expose helpers like Is/As or methods
- For cross-package contracts, document error semantics carefully

---

## 4) Typed-nil Interfaces and Pitfalls

```go
type Reader interface { Read([]byte) (int, error) }

func NewReader() Reader {
  var r *bytes.Reader = nil
  return r // non-nil interface holding typed-nil -> callers see i != nil
}
```

Rule:
- Return real nil interface when signalling absence; avoid returning typed-nil inside interface
- Callers should check both interface and, when asserted, the underlying value

---

## 5) Error Values and Context (fmt.Errorf, %w)

- Add context as you bubble up: function/op names, key values
- Keep messages user-friendly at the edges and technical internally
- Avoid logs + returns duplication; return errors and let the caller decide logging

```go
if err := svc.Do(ctx, req); err != nil {
  return fmt.Errorf("svc.Do user=%s: %w", req.UserID, err)
}
```

---

## 6) Common Mistakes and Gotchas

1) Comparing wrapped errors with ==
```go
// ❌ err == os.ErrNotExist
// ✅ errors.Is(err, os.ErrNotExist)
```

2) Swallowing errors
```go
// ❌ Ignoring return error
if _, err := f(); err != nil { /* TODO */ }
// ✅ Handle or propagate with context
```

3) Returning typed-nil interfaces
```go
// ✅ Return (nil, err) or nil interface when absent
```

4) Excessive error wrapping
```go
// Keep wrap depth reasonable; add value at each layer
```

---

## 7) Best Practices

- Define errors close to where they’re produced; export only when necessary
- Use %w for wrapping; Is/As for checking
- Keep messages consistent and actionable; redact sensitive data
- Prefer typed errors for programmatic handling; provide helpers
- Document error contracts on public APIs

---

## 8) Performance Considerations

- Wrapping allocates; fine for most code but avoid in tight loops
- Avoid building error strings with fmt in hot paths; precompute or reuse
- Don’t log on every layer; avoid duplicate I/O

---

## 9) Advanced Challenge Questions

1) Why can an interface be non-nil while holding a nil pointer?
- Because interface equality uses (dynamic type, dynamic value); (T, nil) != (nil, nil).

2) How do errors.Is and errors.As differ?
- Is matches sentinels anywhere in chain; As finds the first error assignable to the target type.

3) When would you choose sentinels over typed errors?
- For ubiquitous conditions (io.EOF), or where simple checks suffice and coupling is acceptable.

