# Go Types: From Basics to Advanced (with runnable examples)

Run these examples
- Basics (bool, numeric, byte/rune): go run types/001_basic_types.go
- Defined type vs alias: go run types/002_defined_vs_alias.go
- Conversions and underlying types: go run types/003_conversions.go
- Untyped consts and iota enums: go run types/004_untyped_iota.go
- Methods on defined types: go run types/005_methods_on_types.go
- Type assertions & switches: go run types/006_type_assert_switch.go
- Generics & constraints: go run types/007_generics_constraints.go
- comparable and sets/maps: go run types/008_comparable_set.go
- Custom JSON for newtypes: go run types/010_json_newtype.go

---
## Table of Contents
1. [Built-in types & zero values](#toc-1-builtins)
2. [Defined types vs type aliases](#toc-2-defined-vs-alias)
3. [Conversions, underlying types, and assignment rules](#toc-3-conversions)
4. [Untyped constants, iota, and enums](#toc-4-untyped)
5. [Methods on defined types; method sets and API design](#toc-5-methods)
6. [Interfaces: satisfaction, assertions, and switches](#toc-6-interfaces)
7. [Generics: type parameters, constraints, and ~underlying types](#toc-7-generics)
8. [comparable, maps/sets, and pitfalls](#toc-8-comparable)
9. [Custom serialization for newtypes](#toc-9-custom-serialization)



10. Common mistakes and gotchas
11. Best practices
12. FAQ

---

<a id="toc-1-builtins"></a>

## 1) Built-in types & zero values

- Booleans: bool (zero: false)
- Numeric: int/uint (implementation-sized), int8/16/32/64, uint8/16/32/64, uintptr
- Floating: float32/64; Complex: complex64/128
- Aliases: byte (uint8), rune (int32)
- Strings are immutable; zero value ""

See: types/001_basic_types.go
- See deeper discussion on runes and iteration: [Strings guide section](../strings/StringsGuide.md#runes-range-semantics)


---

<a id="toc-2-defined-vs-alias"></a>

## 2) Defined types vs type aliases

- Defined type creates a distinct type with the same underlying type
  - type UserID int // new, distinct type
- Type alias makes an exact alias of an existing type
  - type MyInt = int

Why it matters:
- Defined types enable methods, stronger APIs, and prevent accidental mixing
- Aliases ease refactors and interop without new methods or distinctness

See: types/002_defined_vs_alias.go

---

<a id="toc-3-conversions"></a>

## 3) Conversions, underlying types, and assignment rules

- Assignment between distinct defined types requires explicit conversion, even if underlying types match
- Slice[T] to Slice[U] does not convert elementwise implicitly
- Use underlying-type aware constraints in generics with ~

See: types/003_conversions.go

---

<a id="toc-4-untyped"></a>

## 4) Untyped constants, iota, and enums

- Untyped constants adopt a type at use site; useful for precise literals without overflow
- iota helps build enums; pair with String methods or use stringer

See: types/004_untyped_iota.go

---

<a id="toc-5-methods"></a>

## 5) Methods on defined types; method sets and API design

- You can attach methods to defined types (including on built-in underlying types)
- Decide pointer vs value receiver based on mutability and size

See: types/005_methods_on_types.go

---

<a id="toc-6-interfaces"></a>

## 6) Interfaces: satisfaction, assertions, and switches

- Satisfaction is implicit; method sets matter (T vs *T)
- Type assertions/switches extract dynamic values from interface values

See: types/006_type_assert_switch.go

---

<a id="toc-7-generics"></a>

## 7) Generics: type parameters, constraints, and ~underlying types

- Use constraints to accept families of types
- Use ~T to accept any defined type whose underlying type is T

See: types/007_generics_constraints.go

---

<a id="toc-8-comparable"></a>

## 8) comparable, maps/sets, and pitfalls

- Only comparable types are valid map keys
- floating NaN != NaN; be careful with floats as keys

See: types/008_comparable_set.go

---

<a id="toc-9-custom-serialization"></a>

## 9) Custom serialization for newtypes

- Defined types don’t automatically get special (un)marshal behavior
- Implement encoding/json Marshaler/Unmarshaler for custom wire formats

See: types/010_json_newtype.go

---

## 10) Common mistakes and gotchas

1) Confusing alias with defined type
- Aliases don’t create distinct types; you can’t attach new methods to an alias

2) Assuming implicit conversions
- Distinct types need explicit conversion; slices of different element types don’t convert

3) Typed-nil interface pitfalls
- interface value can be non-nil while holding a typed nil; check both

4) Using int assuming 64-bit everywhere
- int size varies by arch; use explicit sizes for I/O, network, and persistence

5) Using float as map keys or for money
- NaN breaks equality; prefer integer cents or decimal libraries

---

## 11) Best practices

- Use defined types to encode invariants (UserID, Email) and add methods
- Prefer explicit sizes for on-disk/on-wire formats
- Use iota + Stringer for readable enums
- Document pointer vs value receiver choices
- Prefer generics with clear constraints over interface{} (any)

---

## 12) FAQ

1) Explain defined type vs alias and their impact on method sets
2) Why do slices of different element types not convert automatically?
3) How does ~ in constraints affect generic functions? Show an example.
4) How do you avoid typed-nil interface bugs?
5) Discuss pitfalls of using float64 as map keys and alternatives.

