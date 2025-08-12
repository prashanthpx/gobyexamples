# Go Slices: Complete Guide

Run these examples
- Hidden retention fix: go run slices/mistakes/retention.go

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

## **Summary**

| Scenario | Capacity Behavior | Memory Efficient? |
|----------|-------------------|-------------------|
| `make([]T, N)` | `cap = N` | ❌ Not efficient for append-heavy code |
| `make([]T, 0, C)` | `cap = C` | ✅ Efficient if you expect appends |
| `append()` | grows cap automatically | ⚠️ Yes, but may cause realloc/copy |
| **Best practice** | preallocate with make and cap | ✅ **Recommended for performance** |

---

> **💡 Pro Tip:** Always consider preallocating slice capacity when you know the approximate final size to avoid unnecessary memory allocations and improve performance.
