# Packages and Modules in Go: Advanced Developer Guide

## Table of Contents
1. [Packages, Modules, and go.mod](#toc-1-packages)
2. [Import Paths and Semantic Import Versioning (v2+)](#toc-2-import-paths)
3. [Internal Packages and Visibility](#toc-3-internal)
4. [init Functions and Initialization Order](#toc-4-init)
5. [Build Tags and Files](#toc-5-build-tags)
6. [Replace, Exclude, and Vendoring](#toc-6-replace-exclude)
7. [Workspaces (go work)](#toc-7-workspaces)
8. [Common Mistakes and Gotchas](#toc-8-mistakes)
9. [Best Practices](#toc-9-best)
10. [Advanced Challenge Questions](#toc-10-advanced)

---

## 1) Packages, Modules, and go.mod

- Package: unit of compilation; one directory = one package
- Module: versioned collection of packages; root has a go.mod file

```bash
# Create a new module in the current directory
go mod init github.com/you/yourmod

# Add or update dependencies
go get example.com/lib@v1.2.3

# Tidy unused/needed deps
go mod tidy
```

Notes:
- Keep module paths canonical; changing them is breaking for importers
- Commit go.mod and go.sum; they define your module’s build

---

## 2) Import Paths and Semantic Import Versioning (v2+)

- For major versions v2 and above, the module path must include the version suffix
  - example.com/lib/v2
- Importers must import the suffixed path

```go
// go.mod
module example.com/lib/v2

// import site
import "example.com/lib/v2/mypkg"
```

---

## 3) Internal Packages and Visibility

- Any package under internal/ is importable only by siblings/descendants of the parent
- Enforce encapsulation boundaries within a module

```text
mymod/
  go.mod
  internal/utils/...
  cmd/app/main.go  // can import internal/utils
  other/xyz.go     // can import internal/utils
external/
  othermod/...     // cannot import mymod/internal/utils
```

---

## 4) init Functions and Initialization Order

- init runs after package-level var initialization and before main()
- Multiple init functions allowed; order within a single file is top-to-bottom
- Avoid complex init logic; prefer explicit constructors

```go
var Default *Client

func init() {
  Default = &Client{Timeout: time.Second}
}
```

Pitfalls:
- init across packages can create import-order dependence; keep simple and deterministic

---

## 5) Build Tags and Files

- Use //go:build lines (and // +build for legacy) to include/exclude files

```go
//go:build linux && amd64
package sys

// file only builds on linux/amd64
```

- Suffixes also affect builds: file_linux.go, file_windows.go

---

## 6) Replace, Exclude, and Vendoring

- replace in go.mod swaps a dependency source (local path or different version)
- exclude prevents a version from being used

```go
replace example.com/lib v1.2.3 => ../lib
exclude example.com/lib v1.2.2
```

Vendoring:
```bash
go mod vendor   # creates vendor directory
go build -mod=vendor
```

Use vendoring when policy requires pinning dependencies in repo.

---

## 7) Workspaces (go work)

- Workspaces allow multiple modules to be built together without replace

```bash
go work init ./modA ./modB
# adds go.work and go.work.sum
```

---

## 8) Common Mistakes and Gotchas

1) Import cycles
```go
// ❌ a imports b, b imports a -> cycle
// ✅ Extract shared code to a third package
```

2) Overusing init
```go
// ❌ Hidden side effects, order surprises
// ✅ Prefer explicit New/Configure functions
```

3) Breaking import paths on major version bumps
```go
// ❌ module path missing /v2 in go.mod for v2+
```

4) Using dot imports
```go
// ❌ import . "pkg" pollutes namespace
// ✅ Use explicit identifiers for clarity
```

5) Ignoring go mod tidy
```bash
// ❌ go.mod/go.sum drift
// ✅ run: go mod tidy
```

---

## 9) Best Practices

- Keep packages cohesive; avoid giant “util” packages
- Limit exported surface; prefer small, focused APIs
- Use internal/ to enforce boundaries
- Keep init simple; wire dependencies explicitly
- Adopt semantic import versioning correctly for v2+
- Use workspaces for multi-repo or multi-module dev flows

---

## 10) Advanced Challenge Questions

1) Why is semantic import versioning required for v2+ modules?
- To allow co-existence of multiple major versions and avoid breaking importers.

2) When to use workspaces vs replace in go.mod?
- Workspaces for local multi-module development; replace is a per-module override.

3) How do build tags and file suffixes interplay?
- Both filter files for a target platform; suffixes are simpler for OS/arch; tags for arbitrary predicates.

4) How can internal/ help in large codebases?
- It prevents external packages from importing implementation details, enforcing modular boundaries.

