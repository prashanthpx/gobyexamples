# Go Slices: Complete Guide

Run these examples
- Hidden retention fix: go run slices/mistakes/retention.go


## Slice syntax refresher

General form: `s[low : high : max]`

- low = starting index (inclusive). Defaults to 0
- high = ending index (exclusive). Defaults to len(s)
- max = capacity limit (optional). Defaults to cap(s)

So:
- s[low:high] → new slice from index low up to (but not including) high
- s[low:high:max] → also sets capacity = max - low

How to remember
- Think of [:] as “from : up to (but not including)”
- Left side (before colon) = where you start
- Right side (after colon) = where you stop (not included)
- If omitted, defaults are 0 on the left, len(s) on the right

Examples
```go
s := []int{10, 20, 30, 40, 50}
fmt.Println(s[1:3]) // [20 30]
fmt.Println(s[:3])  // [10 20 30]
fmt.Println(s[2:])  // [30 40 50]
fmt.Println(s[:])   // [10 20 30 40 50]
```

## **Slice Basics: Length vs Capacity**

A slice in Go is a lightweight data structure that wraps:
- **A pointer** to an array
- **A length** (the number of actual elements)
- **A capacity** (how many elements can be added before reallocating)

### **Basic Example**

```go
s := make([]int, 3, 5)
fmt.Println(len(s)) // 3 (number of elements currently in the slice)
fmt.Println(cap(s)) // 5 (maximum number of elements it can hold before growing)
```

```go
s := []int{1, 2, 3}
fmt.Println(len(s)) // 3
fmt.Println(cap(s)) // 3 (usually, unless explicitly made larger)
```

## **When You Append to a Slice**

```go
s := []int{1, 2, 3}
s = append(s, 4)
```

### **What Happens?**

- **If `len < cap`**: The element is added to the same underlying array
- **If `len == cap`**: Go allocates a new underlying array, copies the old elements, adds the new one, and adjusts capacity — often doubling the capacity

### **Parameter Changes**

| Operation | `len()` | `cap()` |
|-----------|---------|---------|
| `append()` (within cap) | ✅ Increases | ❌ Unchanged |
| `append()` (exceeds cap) | ✅ Increases | ✅ Increases (often doubles) |


### Q: In a slice, if `m2 := make([]int, 10)`, what happens if we add more than 10 elements?

A: Great one — this highlights a classic difference between maps and slices in Go.

- `m2 := make([]int, 10)` creates a slice with length 10 and (by default) capacity 10.
  - Initially: `len(m2) == 10`, `cap(m2) == 10`
  - Elements are zero-initialized (0 for int)

You cannot do `m2[10] = 42` (that panics: index out of range). Use `append` instead:

```go
m2 := make([]int, 10)          // len=10, cap=10
fmt.Println(len(m2), cap(m2))  // 10 10

m2 = append(m2, 100)           // add one more (11th element)
fmt.Println(len(m2), cap(m2))  // 11 20  // capacity grew
```

- On the first append beyond capacity, Go allocates a new backing array with a larger capacity, copies existing elements, then appends the new value
- Growth pattern: not spec’d but typically doubles for smaller capacities (10 → 20 → 40 → 80 …); for large slices it grows more conservatively (~25%)

Contrast with map
- `make(map[K]V, 10)` takes a hint; maps can grow to hold unlimited entries
- `make([]T, n)` sets an actual length; `append` grows by allocating new backing arrays as needed

✅ Summary
- `make([]int, 10)` gives 10 zeroed elements
- Appending beyond 10 triggers automatic capacity growth (new array + copy)
- Capacity usually doubles early, then grows slower for large sizes
- Direct indexing beyond the current length panics

## **Appending One Slice to Another**

```go
a := []int{1, 2}
b := []int{3, 4, 5}

a = append(a, b...)
```

### **Outcome:**
- `len(a)` increases by `len(b)`
- `cap(a)` may or may not increase depending on whether the existing capacity of `a` can accommodate `b`

## **Advantages of Capacity**

### **✅ 1. Efficient Memory Use**
By preallocating enough capacity, you avoid multiple re-allocations and copies.

```go
s := make([]int, 0, 1000) // Efficient if you know you'll append ~1000 items
```

### **✅ 2. Performance Boost**
Each time Go resizes a slice (when capacity is exceeded), it:
- Allocates a new array
- Copies old data
- Points the slice to the new array

Avoiding this with good capacity planning improves performance.

### **✅ 3. Use in Slice Tricks (e.g., reslicing)**

```go
s := make([]int, 0, 10)
s = s[:5] // increases length up to existing capacity
```

## **Creating Slices Without Specifying Capacity**

### **Case: Slice Literal or Basic `make()` Usage**

```go
s := []int{}                 // len = 0, cap = 0
s := make([]int, 0)          // len = 0, cap = 0
s := make([]int, 5)          // len = 5, cap = 5
```

> **Note:** If you don't specify capacity, Go defaults `cap = len` — so when you append, it needs to allocate a new underlying array.

### **When You Append Elements**

```go
s := []int{}       // len=0, cap=0
s = append(s, 1)   // allocates a new array (cap=1)
s = append(s, 2)   // allocates again (cap grows, often doubled)
```

**Result:** Go performs reallocation each time capacity is exceeded, incurring cost: memory allocation + copying existing elements to a new array.

## **How Much Capacity is Allocated?**

Go doesn't document the exact algorithm, but the general behavior is:
- When cap is exceeded, Go grows the capacity automatically
- It typically **doubles the capacity** each time (up to a point), then grows linearly

### **Example: Capacity Growth Pattern**

```go
var s []int
for i := 0; i < 10; i++ {
    s = append(s, i)
    fmt.Printf("Len: %d, Cap: %d\n", len(s), cap(s))
}
```

**Output:**
```
Len: 1, Cap: 1
Len: 2, Cap: 2
Len: 3, Cap: 4
Len: 4, Cap: 4
Len: 5, Cap: 8
Len: 6, Cap: 8
Len: 7, Cap: 8
Len: 8, Cap: 8
Len: 9, Cap: 16
Len: 10, Cap: 16
```

## **Memory Optimization**

### Prevent hidden retention with three-index slicing
```go
// ❌ Small slice retains the whole big array in memory
big := make([]byte, 1<<20) // 1MB
sub := big[:10]            // cap is large; big cannot be GC'd while sub is alive
_ = sub

// ✅ Cap the slice to the current length to drop the tail capacity
sub2 := big[:10:10] // len=10, cap=10
_ = sub2
```

### Copy vs share when appending
```go
// ❌ Sharing backing array can surprise callers
base := []int{1,2,3}
view := base[:2]
view = append(view, 9) // may overwrite base[2]

// ✅ Force copy when you need isolation
isolated := append([]int(nil), view...) // copy
```


### **❌ Not Setting Capacity = "Memory Not Optimized"**

In a tight loop or large-scale slice building, failing to specify capacity can lead to multiple reallocations, each involving:
- New memory allocation
- Copying existing data to the new array

### **✅ Best Practice**

If you know or can estimate the number of elements:

```go
s := make([]int, 0, 1000) // allocates space for 1000 ints upfront
```

This avoids reallocation and copying, making the operation more memory- and time-efficient.


## Passing Slices vs Arrays, and Why Use *[N]T

Slices ([]T)
- A slice is a small header (pointer, len, cap) pointing to an underlying array.
- Passing a slice copies the header, but both caller and callee share the same underlying array.
- Inside the callee:
  - s[i] = ... mutates the caller’s data.
  - s = append(s, ...) may allocate a new array when capacity is exceeded; unless you return the new slice (or pass *[]T), the caller won’t see the growth.

Arrays ([N]T)
- Passing an array by value copies all N elements. Mutations in the callee won’t affect the caller.
- To mutate the caller’s array, pass *[N]T (pointer to array).

Why pass a slice “as an array” to a function?
- Sometimes you want the callee to operate on exactly N elements with strong, compile-time shape guarantees.
- Go 1.17+ allows converting a slice (or subslice) to a pointer-to-array without copying, provided len(slice) >= N.
- This bakes the invariant (exact length) into the type and simplifies inner loops.

Use cases
- Fixed-size algorithms/invariants: crypto blocks, UUIDs, hashes, wire headers, etc. (*[16]byte, *[32]byte, ...)
- Zero-copy, index-safe loops: callee can index 0..N-1 without repeated checks.
- APIs that require [N]T for clarity/safety.

Example: operate on an exact 16-byte block (no copies)
```go
// Callee wants EXACTLY 16 bytes.
func xorBlock(dst, a, b *[16]byte) {
    for i := 0; i < 16; i++ {
        dst[i] = a[i] ^ b[i]
    }
}

func main() {
    a := make([]byte, 16)
    b := make([]byte, 16)
    dst := make([]byte, 16)

    // Convert slices to *[16]byte (Go 1.17+). Zero-copy view.
    xorBlock((*[16]byte)(dst), (*[16]byte)(a), (*[16]byte)(b))
}
```
Notes:
- If any slice’s len < 16, that conversion panics at runtime — enforcing the contract loudly.
- You can also convert a sub-slice if it has enough length, e.g., (*[16]byte)(buf[i:]).

Comparison with a plain slice API
```go
// Flexible but needs checks:
func xorBlockSlice(dst, a, b []byte) {
    if len(dst) < 16 || len(a) < 16 || len(b) < 16 {
        panic("need 16 bytes")
    }
    for i := 0; i < 16; i++ {
        dst[i] = a[i] ^ b[i]
    }
}
```
Both work, but *[16]byte bakes the invariant into the type and avoids repeated length checks at call sites.

When to use which?
- Most code: accept a slice ([]T). It’s idiomatic, flexible, and lets callers pass any length.
- Exactly-N contract, performance clarity, or API design: accept a pointer to array (*[N]T) and let callers convert with (*[N]T)(s). This gives:
  - Stronger guarantees (exact size)
  - Zero-copy conversion from slices
  - Simpler inner loops that assume fixed bounds

Reminders
- Arrays passed by value copy. If you need mutation, use *[N]T.
- Converting []T → *[N]T requires Go 1.17+ and len(slice) >= N, else it panics.

## **Summary**

| Scenario | Capacity Behavior | Memory Efficient? |
|----------|-------------------|-------------------|
| `make([]T, N)` | `cap = N` | ❌ Not efficient for append-heavy code |
| `make([]T, 0, C)` | `cap = C` | ✅ Efficient if you expect appends |
| `append()` | grows cap automatically | ⚠️ Yes, but may cause realloc/copy |
| **Best practice** | preallocate with make and cap | ✅ **Recommended for performance** |

---

> **💡 Pro Tip:** Always consider preallocating slice capacity when you know the approximate final size to avoid unnecessary memory allocations and improve performance.
