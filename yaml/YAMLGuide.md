# YAML in Go: Advanced Developer Guide

## Table of Contents
1. YAML Libraries and Setup
2. Struct Tags and Mapping Rules
3. Strict Decoding and Unknown Fields
4. Time, Numbers, and String Pitfalls
5. Custom Marshaling/Unmarshaling
6. Anchors, Aliases, and Merge Keys
7. Common Mistakes and Gotchas
8. Best Practices
9. Advanced Challenge Questions


Run these examples
- Strict decoding: go run yaml/examples/strict_decode.go

---

## 1) YAML Libraries and Setup

Popular choices:
- gopkg.in/yaml.v3 (canonical, feature-rich)
- sigs.k8s.io/yaml (wraps yaml.v3 with JSON-compatible behavior)

```go
// go get gopkg.in/yaml.v3
import "gopkg.in/yaml.v3"
```

---

## 2) Struct Tags and Mapping Rules

```go
type Item struct {
  Name  string   `yaml:"name"`
  Count int      `yaml:"count,omitempty"`
  Tags  []string `yaml:"tags,omitempty"`
}
```

Notes:
- Unexported fields are ignored
- omitempty skips zero values; define what zero means in your app
- Inline structs use `,inline` to merge fields (watch for conflicts)

---

## 3) Strict Decoding and Unknown Fields

By default, unknown fields are ignored. Enable strict decoding to catch typos.

```go
var it Item
dec := yaml.NewDecoder(r)
dec.KnownFields(true)
if err := dec.Decode(&it); err != nil { return err }
```

With sigs.k8s.io/yaml, you can decode via json first to leverage json.Unmarshal behavior, but it’s not strict by default either.

---

## 4) Time, Numbers, and String Pitfalls

- YAML has implicit typing; strings like "on", "yes" can be parsed as booleans in some modes
- Prefer quoting ambiguous scalars: "on", "01"
- time.Time formatting: use RFC3339 strings; consider custom (un)marshaling for other formats

```yaml
count: "01"    # keep as string
flag: "on"     # avoid implicit bool
when: "2024-06-01T10:00:00Z"
```

---

## 5) Custom Marshaling/Unmarshaling

```go
type Duration struct{ time.Duration }

func (d *Duration) UnmarshalYAML(n *yaml.Node) error {
  var s string
  if err := n.Decode(&s); err != nil { return err }
  dur, err := time.ParseDuration(s)
  if err != nil { return err }
  d.Duration = dur
  return nil
}

func (d Duration) MarshalYAML() (interface{}, error) {
  return d.Duration.String(), nil
}
```

---

## 6) Anchors, Aliases, and Merge Keys

yaml.v3 supports anchors/aliases and merge keys.

```yaml
base: &base
  name: common
  count: 1

item1:
  <<: *base
  name: specific
```

Be cautious; excessive indirection can harm readability.

---

## 7) Common Mistakes and Gotchas

1) Relying on default decoding and missing typos
```go
// Enable KnownFields(true) for strict decoding
```

2) Unexported fields not set
```go
// Export fields to be decoded; yaml/json ignore unexported ones
```

3) Ambiguous scalars parsed unexpectedly
```go
// Quote strings like "on", "yes", "01" to preserve types
```

4) Ignoring time zones and formats
```go
// Parse times explicitly; document expected layouts
```

---

## 8) Best Practices

- Keep schemas documented; validate after decoding
- Use KnownFields(true) in config-heavy apps
- Quote ambiguous strings and leading-zero numbers
- Prefer simple structures over anchors/aliases for maintainability

---

## 9) Advanced Challenge Questions

1) How do you enforce strict YAML decoding?
- Use Decoder.KnownFields(true) with yaml.v3 to reject unknown fields.

2) When should you implement custom (Un)Marshal for YAML?
- For types like time.Duration or custom enums where string forms are preferable.

3) How do you avoid implicit typing surprises in YAML?
- Quote ambiguous scalars and validate types post-decode.

