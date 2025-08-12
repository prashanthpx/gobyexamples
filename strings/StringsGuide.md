# Strings, Runes, and Bytes in Go: Advanced Developer Guide

## Table of Contents
1. Strings Are Immutable
2. Bytes vs Runes (UTF‑8)
3. Indexing, Iteration, and Unicode
4. Conversions and Allocations
5. Building Strings Efficiently
6. Common Mistakes and Gotchas
7. Best Practices
8. Performance Considerations
9. Advanced Challenge Questions

---

## 1) Strings Are Immutable

A string is a read-only slice of bytes. Any operation that “changes” a string allocates a new one.

```go
s := "go"
s2 := s + "lang" // allocates new string
_ = s2
```

---

## 2) Bytes vs Runes (UTF‑8)

- byte = alias for uint8 (raw bytes)
- rune = alias for int32 (Unicode code point)

```go
b := []byte("Hi")
r := []rune("Hi")
fmt.Println(len(b), len(r)) // 2 2

x := "é"              // U+00E9; 2 bytes in UTF‑8
fmt.Println(len(x))    // 2 bytes
fmt.Println(len([]rune(x))) // 1 rune
```

---

## 3) Indexing, Iteration, and Unicode

Combining characters and grapheme clusters can make “characters” span multiple runes.

```go
// "é" can be a single code point U+00E9,
// or composed: "e" + U+0301 (combining acute accent)
base := "e\u0301" // two runes, one grapheme cluster
a := "é"           // one rune
fmt.Println(len([]rune(base)), len([]rune(a))) // 2 1
```

If you need user-perceived characters, iterate by grapheme clusters via a library (e.g., golang.org/x/text/transform and x/text/unicode/norm) or a dedicated grapheme splitter.


- s[i] indexes bytes, not runes; may cut a multibyte rune
- Range over string yields runes (decoded)

```go
s := "éa"
fmt.Printf("% x\n", s)        // c3 a9 61
fmt.Printf("%c %c\n", s[0], s[1]) // Ã © (broken)
for i, r := range s { fmt.Println(i, string(r)) } // 0 "é" 2 "a"
```

Use utf8.DecodeRune to walk bytes safely.

---

## 4) Conversions and Allocations

bytes.Reader vs strings.Reader:
```go
rb := bytes.NewReader([]byte("data"))   // for bytes
rs := strings.NewReader("data")          // for strings
_ = rb; _ = rs
```

Converting between string and []byte allocates:


- string(b) copies the bytes; []byte(s) copies the data
- Avoid repeated conversions in hot paths; cache either form

```go
b := []byte("data")
s := string(b)   // allocation
bb := []byte(s)  // allocation
_ = bb
```

---

## 5) Building Strings Efficiently

Prefer strings.Builder or bytes.Buffer for incremental construction.

```go
var b strings.Builder
b.Grow(64)
for i := 0; i < 3; i++ { b.WriteString("x") }
s := b.String() // single allocation
```

For bytes: use bytes.Buffer and convert once at the end.

---

## 6) Common Mistakes and Gotchas

1) Indexing by bytes when you need characters
```go
// Use rune iteration or utf8.DecodeRune when handling Unicode
```

2) Assuming len(string) == number of characters
```go
// len is bytes; count runes with utf8.RuneCountInString
```

3) Modifying strings in place
```go
// Convert to []byte or []rune, modify, then convert back
```

4) Excessive conversions between string and []byte
```go
// Cache one representation; avoid repeated allocations
```

5) Using fmt.Sprintf to build strings in loops
```go
// Prefer strings.Builder or bytes.Buffer
```

---

## 7) Best Practices

- Choose []byte for I/O and mutation; string for stable text
- Iterate over runes when user-visible character boundaries matter
- Pre-size Builders/Buffers with Grow when possible
- Use %q when logging strings to reveal whitespace/escapes

---

## 8) Performance Considerations

- Conversions allocate; measure and minimize in hot paths
- strings.Builder avoids intermediate allocations; reuse Buffers for large data
- Avoid per-iteration Sprintf; write directly to a Buffer/Builder

---

## 9) Advanced Challenge Questions

1) Why does len("") differ from rune count for some strings?
- UTF‑8 encodes some runes as multiple bytes; len reports bytes.

2) How do you safely slice a string by characters?
- Convert to []rune first or walk with utf8.DecodeRune to avoid splitting multibyte runes.

3) When to use strings.Builder vs bytes.Buffer?
- Builder for text you will turn into a string; Buffer when you need []byte or mixed I/O.

