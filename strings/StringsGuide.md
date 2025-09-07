# Strings, Runes, and Bytes in Go: Advanced Developer Guide

## Table of Contents
1. [Strings Are Immutable](#toc-1-immutable)
2. [Bytes vs Runes (UTF‑8)](#toc-2-bytes-runes)
3. [Indexing, Iteration, and Unicode](#toc-3-indexing-iteration)
4. [Conversions and Allocations](#toc-4-conversions)
5. [Building Strings Efficiently](#toc-5-building)
6. [Common Mistakes and Gotchas](#toc-6-mistakes)
7. [Best Practices](#toc-7-best)
8. [Performance Considerations](#toc-8-perf)
9. [Advanced Challenge Questions](#toc-9-advanced)
Run these examples
- café iteration: go run strings/examples/cafe_iter.go
- emoji iteration: go run strings/examples/emoji_iter.go
- grapheme clusters: go run strings/examples/grapheme_demo.go

- invalid UTF-8 handling: go run strings/examples/invalid_utf8.go

- clean invalid UTF-8 bytes: go run strings/examples/clean_invalid_utf8.go

---


---

<a id="toc-1-immutable"></a>

## 1) Strings Are Immutable

A string is a read-only slice of bytes. Any operation that “changes” a string allocates a new one.

```go
s := "go"
s2 := s + "lang" // allocates new string
_ = s2
```

---

<a id="toc-2-bytes-runes"></a>

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

<a id="toc-3-indexing-iteration"></a>

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


<a id="runes-range-semantics"></a>

### Runes and range over strings: index and value semantics

What is a rune?
- In Go, a rune is just an alias for int32.
- It represents a single Unicode code point (e.g., 'A' = U+0041, 'é' = U+00E9, '猫' = U+732B).
- A string in Go is a read-only sequence of bytes (UTF-8 encoded). One rune may take 1–4 bytes in UTF-8.

What does range over a string return?
- for i, r := range s { ... }
  - i = byte index (offset) of the start of this rune within the UTF-8 byte sequence
  - r = the rune (Unicode code point) decoded at that position
- “Byte index (start of rune)” means: if the current rune begins at byte 3 of the string’s underlying bytes, then i == 3.

Why not “character index”?
- Because Go strings are bytes. UTF-8 encodes some runes with multiple bytes, so the index that range gives you is a byte offset, not a “rune count” or “character position.”

Mini examples

ASCII vs non-ASCII
```go
s := "café" // bytes: [63 61 66 C3 A9]
fmt.Println(len(s))                    // 5 bytes
fmt.Println(utf8.RuneCountInString(s)) // 4 runes

for i, r := range s {
    fmt.Printf("i=%d  r=%U %q\n", i, r, r)
}
// i=0  r=U+0063 "c"
// i=1  r=U+0061 "a"
// i=2  r=U+0066 "f"
// i=3  r=U+00E9 "é"   <-- note i=3 (start byte of 'é')
```

Emoji / multiple runes
```go
s := "👋🏽" // waving hand + skin tone modifier (two runes)
for i, r := range s {
    fmt.Printf("i=%d  r=%U %q\n", i, r, r)
}
// i=0  r=U+1F44B "👋"
// i=4  r=U+1F3FD "🏽"
// (each rune uses multiple bytes; second starts at byte index 4)
```

Bytes vs runes vs strings (at a glance)
- byte = alias for uint8. One raw byte. []byte(s) gives the UTF-8 bytes of s.
- rune = int32 Unicode code point. []rune(s) converts to runes (decodes UTF-8).
- string = read-only UTF-8 byte sequence.

Common operations
- Count runes: utf8.RuneCountInString(s)
- Iterate runes: for i, r := range s { ... }
- Handling invalid UTF-8: range yields utf8.RuneError (U+FFFD) for invalid sequences; width still advances. See example: strings/examples/invalid_utf8.go

- Index by rune position (not byte): convert first → rs := []rune(s); fmt.Println(rs[3])

Cleaning invalid UTF-8
- When ingesting untrusted/external data, sanitize first so downstream code doesn’t see invalid sequences.
- Use bytes.ToValidUTF8 to replace invalid bytes with a replacement of your choice (e.g., "?").

Example
```go
cleaned := bytes.ToValidUTF8([]byte(s), []byte("?"))
safe := string(cleaned)
```
See runnable: strings/examples/clean_invalid_utf8.go

Warning
- bytes.ToValidUTF8 alters the data; if exact fidelity matters, prefer rejecting/quarantining the input or logging and failing fast instead of silently rewriting bytes.
- When sanitizing, consider retaining the original raw []byte alongside the cleaned string for audit/debug purposes, and document that replacement characters may appear.


- Get next rune from a byte offset: r, size := utf8.DecodeRuneInString(s[i:])

Gotcha: user-visible “characters”
- A single user-perceived character can be multiple runes (e.g., “🇮🇳” flag, emoji + skin tone, letters + combining accents). If you need to iterate what users see as characters (grapheme clusters), use a grapheme segmenter (e.g., golang.org/x/text/segment), not plain runes.

---

<a id="toc-4-conversions"></a>

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

<a id="toc-5-building"></a>

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

<a id="toc-6-mistakes"></a>

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

<a id="toc-7-best"></a>

## 7) Best Practices

- Choose []byte for I/O and mutation; string for stable text
- Iterate over runes when user-visible character boundaries matter
- Pre-size Builders/Buffers with Grow when possible
- Use %q when logging strings to reveal whitespace/escapes

---

<a id="toc-8-perf"></a>

## 8) Performance Considerations

- Conversions allocate; measure and minimize in hot paths
- strings.Builder avoids intermediate allocations; reuse Buffers for large data
- Avoid per-iteration Sprintf; write directly to a Buffer/Builder

---

<a id="toc-9-advanced"></a>

## 9) Advanced Challenge Questions

1) Why does len("") differ from rune count for some strings?
- UTF‑8 encodes some runes as multiple bytes; len reports bytes.

2) How do you safely slice a string by characters?
- Convert to []rune first or walk with utf8.DecodeRune to avoid splitting multibyte runes.

3) When to use strings.Builder vs bytes.Buffer?
- Builder for text you will turn into a string; Buffer when you need []byte or mixed I/O.

