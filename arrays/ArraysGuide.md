# Go Arrays: Advanced Developer Guide

## **Table of Contents**
1. [Array Fundamentals](#array-fundamentals)
2. [Memory Layout and Performance](#memory-layout-and-performance)
3. [Array Operations](#array-operations)
4. [Arrays vs Slices](#arrays-vs-slices)
5. [Multi-dimensional Arrays](#multi-dimensional-arrays)
6. [Array Comparison and Equality](#array-comparison-and-equality)
7. [Common Mistakes and Gotchas](#common-mistakes-and-gotchas)
8. [Best Practices](#best-practices)
9. [Performance Considerations](#performance-considerations)
10. [Advanced Challenge Questions](#advanced-challenge-questions)

---

## **Array Fundamentals**

### **What is an Array in Go?**
An array is a **fixed-size** sequence of elements of the same type. The size is part of the array's type.

```go
var arr [5]int                    // Array of 5 integers, zero-initialized
var names [3]string               // Array of 3 strings
var matrix [2][3]int              // 2D array: 2 rows, 3 columns
```

### **Key Characteristics**
- **Fixed size**: Size is determined at compile time
- **Value type**: Arrays are values, not references
- **Zero-initialized**: Elements get zero values by default
- **Type includes size**: `[3]int` and `[4]int` are different types

### **Array Declaration and Initialization**
```go
// Declaration with zero values
var arr1 [3]int                   // [0, 0, 0]

// Array literal
arr2 := [3]int{1, 2, 3}          // [1, 2, 3]

// Partial initialization
arr3 := [5]int{1, 2}             // [1, 2, 0, 0, 0]

// Compiler-determined size
arr4 := [...]int{1, 2, 3, 4}     // [4]int{1, 2, 3, 4}

// Index-based initialization
arr5 := [5]int{0: 10, 2: 20, 4: 40} // [10, 0, 20, 0, 40]
```

### Package-scope vs function-scope initialization (:= vs var)

- Declarations and statements must live inside a function body; at package scope you can only have declarations (var, const, type, func).
- Short declaration := is a statement and is only allowed inside functions.

Wrong (package scope):
```go
package main

arr1 := [2]int{0, 1}   // ERROR: expected declaration
```

Correct (inside a function):
```go
package main
import "fmt"

func main() {
    arr1 := [2]int{0, 1}   // OK
    fmt.Println(arr1)
}
```

Alternative at package level (use var):
```go
package main
import "fmt"

var arr1 = [2]int{0, 1}   // valid package-level declaration

func main() {
    fmt.Println(arr1)
}
```

Summary
- Use := inside functions.
- Use var at package scope.
- “expected declaration” means the compiler found a statement where only declarations are allowed.

Bonus: := vs var differences
- := (short declaration):
  - Only inside functions; must initialize; type is inferred; at least one new name on LHS
- var:
  - Works at package or function scope; can specify type or infer; can declare without initialization (zero value)

Examples
```go
func main() {
    x := 10            // inferred int
    var y int          // zero value 0
    var z = 3.14       // inferred float64
    fmt.Println(x, y, z)
}
```


---

## **Memory Layout and Performance**

### **Contiguous Memory Layout**
```go
import "unsafe"

func memoryLayout() {
    arr := [3]int{10, 20, 30}

    fmt.Printf("Array address: %p\n", &arr)
    fmt.Printf("Element 0: %p\n", &arr[0])
    fmt.Printf("Element 1: %p\n", &arr[1])
    fmt.Printf("Element 2: %p\n", &arr[2])

    // Elements are contiguous in memory
    // Address difference = sizeof(int) = 8 bytes on 64-bit
}
```

### **Memory Size**
```go
func arraySize() {
    var arr1 [3]int
    var arr2 [1000]int

    fmt.Println(unsafe.Sizeof(arr1))  // 24 bytes (3 * 8)
    fmt.Println(unsafe.Sizeof(arr2))  // 8000 bytes (1000 * 8)

    // Array size = element_size * length
}
```

### **Stack vs Heap Allocation**
```go
func stackArray() {
    arr := [1000]int{} // Allocated on stack (if doesn't escape)
    fmt.Println(len(arr))
}

func heapArray() *[1000]int {
    arr := [1000]int{} // Escapes to heap due to return
    return &arr
}
```

---

## **Array Operations**

### **Accessing Elements**
```go
func arrayAccess() {
    arr := [5]int{10, 20, 30, 40, 50}

    // Read access
    fmt.Println(arr[0])    // 10
    fmt.Println(arr[4])    // 50

    // Write access
    arr[2] = 100
    fmt.Println(arr)       // [10, 20, 100, 40, 50]

    // Bounds checking at runtime
    // fmt.Println(arr[5])  // Runtime panic: index out of range
}
```

### **Array Length**
```go
func arrayLength() {
    arr1 := [3]int{1, 2, 3}
    arr2 := [...]string{"a", "b", "c", "d"}

    fmt.Println(len(arr1)) // 3
    fmt.Println(len(arr2)) // 4

    // Length is compile-time constant
    const size = len(arr1) // This works!
}
```

### **Iterating Over Arrays**
```go
func arrayIteration() {
    arr := [4]string{"Go", "is", "awesome", "!"}

    // Index-based iteration
    for i := 0; i < len(arr); i++ {
        fmt.Printf("%d: %s\n", i, arr[i])
    }

    // Range-based iteration
    for index, value := range arr {
        fmt.Printf("%d: %s\n", index, value)
    }

    // Value-only iteration
    for _, value := range arr {
        fmt.Println(value)
    }

    // Index-only iteration
    for index := range arr {
        fmt.Println(index)
    }
}
```

---

## **Arrays vs Slices**

### **Key Differences**

| Feature | Array | Slice |
|---------|-------|-------|
| **Size** | Fixed at compile time | Dynamic |
| **Type** | `[n]T` (size is part of type) | `[]T` |
| **Memory** | Value type (copied) | Reference type (header) |
| **Zero value** | Array with zero elements | `nil` |
| **Append** | Not possible | `append()` function |
| **Comparison** | Comparable with `==` | Not comparable |

### **Conversion Between Arrays and Slices**
```go
func arraySliceConversion() {
    // Array to slice
    arr := [4]int{1, 2, 3, 4}
    slice1 := arr[:]        // Full slice
    slice2 := arr[1:3]      // Partial slice [2, 3]

    // Slice to array (Go 1.17+)
    slice := []int{1, 2, 3}
    arr2 := [3]int(slice)   // Convert slice to array

    // Array pointer to slice
    arrPtr := &arr
    slice3 := arrPtr[:]     // Slice from array pointer
}
```

### **When to Use Arrays vs Slices**

**Use Arrays when:**
- Fixed, known size at compile time
- Need value semantics (copying)
- Performance-critical code with small, fixed collections
- Mathematical operations (vectors, matrices)

**Use Slices when:**
- Dynamic size needed
- Working with collections that grow/shrink
- Passing to functions (avoid large copies)
- Most general-purpose programming

```go
// ✅ Good use of arrays
type RGB [3]uint8           // Color always has 3 components
type Matrix4x4 [16]float64  // 4x4 matrix always has 16 elements

// ✅ Good use of slices
func processItems(items []Item) {
    // Dynamic collection processing
}
```

---

## **Multi-dimensional Arrays**

### **2D Arrays**
```go
func twoDimensionalArrays() {
    // Declaration
    var matrix [3][4]int

    // Initialization
    grid := [2][3]int{
        {1, 2, 3},
        {4, 5, 6},
    }

    // Access elements
    fmt.Println(grid[0][1]) // 2
    grid[1][2] = 10

    // Iteration
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            fmt.Printf("%d ", grid[i][j])
        }
        fmt.Println()
    }

    // Range iteration
    for i, row := range grid {
        for j, value := range row {
            fmt.Printf("grid[%d][%d] = %d\n", i, j, value)
        }
    }
}
```

### **3D Arrays**
```go
func threeDimensionalArrays() {
    // 3D array: 2 layers, 3 rows, 4 columns
    var cube [2][3][4]int

    // Initialize
    cube[0][1][2] = 42

    // Size calculation
    fmt.Println(unsafe.Sizeof(cube)) // 2 * 3 * 4 * 8 = 192 bytes
}
```

---

## **Array Comparison and Equality**

### **Array Equality**
```go
func arrayEquality() {
    arr1 := [3]int{1, 2, 3}
    arr2 := [3]int{1, 2, 3}
    arr3 := [3]int{1, 2, 4}

    fmt.Println(arr1 == arr2) // true - same values
    fmt.Println(arr1 == arr3) // false - different values

    // Arrays of different sizes are different types
    arr4 := [4]int{1, 2, 3, 4}
    // fmt.Println(arr1 == arr4) // Compile error: different types
}
```

### **Comparable Element Types**
```go
func comparableArrays() {
    // ✅ Comparable - basic types
    arr1 := [2]int{1, 2}
    arr2 := [2]int{1, 2}
    fmt.Println(arr1 == arr2) // true

    // ✅ Comparable - structs with comparable fields
    type Point struct{ X, Y int }
    points1 := [2]Point{{1, 2}, {3, 4}}
    points2 := [2]Point{{1, 2}, {3, 4}}
    fmt.Println(points1 == points2) // true

    // ❌ Not comparable - slices
    // arr3 := [2][]int{{1}, {2}}
    // arr4 := [2][]int{{1}, {2}}
    // fmt.Println(arr3 == arr4) // Compile error
}
```

---

## **Common Mistakes and Gotchas**

### **1. Array Assignment Copies**
```go
func arrayCopyGotcha() {
    arr1 := [3]int{1, 2, 3}
    arr2 := arr1        // Copies entire array!

    arr2[0] = 100
    fmt.Println(arr1)   // [1, 2, 3] - unchanged
    fmt.Println(arr2)   // [100, 2, 3] - modified

    // For large arrays, this is expensive
    largeArr1 := [1000000]int{}
    largeArr2 := largeArr1 // Copies 8MB of data!
}
```

### **2. Function Parameter Copying**
```go
// ❌ Inefficient - copies entire array
func processArray(arr [1000]int) {
    // Working with copy
    arr[0] = 100 // Doesn't affect original
}

// ✅ Efficient - pass pointer
func processArrayPtr(arr *[1000]int) {
    arr[0] = 100 // Modifies original
}

// ✅ Most common - use slice
func processSlice(slice []int) {
    slice[0] = 100 // Modifies original backing array
}

func functionParams() {
    arr := [1000]int{1, 2, 3}

    processArray(arr)    // Copies 8KB
    processArrayPtr(&arr) // Passes 8-byte pointer
    processSlice(arr[:]) // Passes 24-byte slice header
}
```

### **3. Array Size in Type**
```go
func arraySizeInType() {
    var arr3 [3]int
    var arr4 [4]int

    // These are different types!
    // arr3 = arr4 // Compile error: cannot assign [4]int to [3]int

    // Function that takes specific array size
    func process3(arr [3]int) {}
    func process4(arr [4]int) {}

    process3(arr3) // ✅ Works
    // process3(arr4) // ❌ Compile error
}
```

### **4. Range Loop Variable**
```go
func rangeLoopGotcha() {
    arr := [3]int{1, 2, 3}
    var pointers []*int

    // ❌ Wrong - all pointers point to same variable
    for _, v := range arr {
        pointers = append(pointers, &v)
    }

    for _, p := range pointers {
        fmt.Println(*p) // Prints 3, 3, 3
    }

    // ✅ Correct - create new variable
    for _, v := range arr {
        v := v // Create new variable
        pointers = append(pointers, &v)
    }
}
```

### **5. Zero Value Confusion**
```go
func zeroValueConfusion() {
    var arr [3]int
    fmt.Println(arr == [3]int{}) // true - zero value is zero elements

    // But this might be surprising
    var slice []int
    fmt.Println(slice == nil) // true - zero value is nil

    // Array is never nil
    fmt.Println(&arr == nil) // false - array address is never nil
}
```

---

## **Best Practices**

### **1. Prefer Slices for Most Use Cases**
```go
// ❌ Usually avoid - inflexible
func processFixedData(data [100]Item) Result {
    // Can only handle exactly 100 items
}

// ✅ Prefer - flexible
func processData(data []Item) Result {
    // Can handle any number of items
}
```

### **2. Use Arrays for Fixed-Size Data**
```go
// ✅ Good use cases for arrays
type IPv4 [4]byte           // IP address always 4 bytes
type SHA256 [32]byte        // Hash always 32 bytes
type Point3D [3]float64     // 3D point always 3 coordinates

func calculateDistance(p1, p2 Point3D) float64 {
    // Mathematical operations benefit from fixed size
    var sum float64
    for i := 0; i < 3; i++ {
        diff := p1[i] - p2[i]
        sum += diff * diff
    }
    return math.Sqrt(sum)
}
```

### **3. Use Array Pointers for Large Arrays**
```go
// ❌ Inefficient for large arrays
func processLargeArray(arr [10000]int) {
    // Copies 80KB on each call
}

// ✅ Efficient
func processLargeArrayPtr(arr *[10000]int) {
    // Passes 8-byte pointer
    for i := range arr {
        // Process arr[i]
    }
}

// ✅ Most flexible
func processLargeSlice(slice []int) {
    // Works with any size, passes slice header
}
```

### **4. Initialize Arrays Explicitly**
```go
// ✅ Clear initialization
config := [3]string{"dev", "staging", "prod"}

// ✅ Sparse initialization
priorities := [10]int{0: 1, 5: 2, 9: 3} // Only set specific indices

// ✅ Compiler-determined size
weekdays := [...]string{
    "Monday", "Tuesday", "Wednesday",
    "Thursday", "Friday", "Saturday", "Sunday",
}
```

---

## **Performance Considerations**

### **1. Memory Locality**
```go
// ✅ Good cache locality - contiguous memory
func sumArray(arr [1000]int) int {
    sum := 0
    for _, v := range arr {
        sum += v // Sequential memory access
    }
    return sum
}

// ❌ Poor cache locality - scattered pointers
func sumPointers(ptrs [1000]*int) int {
    sum := 0
    for _, p := range ptrs {
        sum += *p // Random memory access
    }
    return sum
}
```

### **2. Copy Cost**
```go
// Benchmark: Array vs Slice vs Pointer
func BenchmarkArrayCopy(b *testing.B) {
    arr := [1000]int{}
    for i := 0; i < b.N; i++ {
        processArray(arr) // Copies 8KB each time
    }
}

func BenchmarkSlicePass(b *testing.B) {
    arr := [1000]int{}
    slice := arr[:]
    for i := 0; i < b.N; i++ {
        processSlice(slice) // Copies 24 bytes each time
    }
}
```

### **3. Bounds Check Elimination**
```go
func optimizedAccess(arr [100]int) int {
    sum := 0
    // Compiler can eliminate bounds checks in simple loops
    for i := 0; i < len(arr); i++ {
        sum += arr[i] // No bounds check needed
    }
    return sum
}
```

---

## **Advanced Challenge Questions**

### **Q1: What's the output?**
```go
func question1() {
    arr := [3]int{1, 2, 3}
    slice := arr[:]

    slice[0] = 100
    fmt.Println(arr[0]) // ?
}
// Answer: 100 (slice shares backing array)
```

### **Q2: Memory usage?**
```go
func question2() {
    arr1 := [1000000]int{}
    arr2 := arr1

    fmt.Println(unsafe.Sizeof(arr1)) // ?
    fmt.Println(unsafe.Sizeof(arr2)) // ?
}
// Answer: 8000000, 8000000 (both arrays are full copies)
```

### **Q3: Compilation error?**
```go
func question3() {
    var arr1 [3]int
    var arr2 [4]int

    arr1 = arr2 // Will this compile?
}
// Answer: No, different types [3]int vs [4]int
```

### **Q4: What happens here?**
```go
func question4() {
    arr := [2][2]int{{1, 2}, {3, 4}}

    for _, row := range arr {
        row[0] = 100 // Does this modify arr?
    }

    fmt.Println(arr) // ?
}
// Answer: [[1 2] [3 4]] - row is a copy, original unchanged
```

### **Q5: Performance comparison?**
```go
func question5() {
    // Which is faster for 1000 elements?

    // Option A: Array parameter
    func processA(arr [1000]int) int { /* ... */ }

    // Option B: Slice parameter
    func processB(slice []int) int { /* ... */ }

    // Option C: Array pointer
    func processC(arr *[1000]int) int { /* ... */ }
}
// Answer: C > B > A (pointer fastest, array copy slowest)
```

---

**🎯 Key Takeaways for Practitioners:**
1. **Arrays are value types** - assignment and parameter passing copies
2. **Size is part of the type** - `[3]int` ≠ `[4]int`
3. **Use slices for most cases** - arrays for fixed-size, mathematical data
4. **Understand memory layout** - contiguous, stack-allocated by default
5. **Know conversion patterns** - array to slice, slice to array (Go 1.17+)

This guide covers essential array concepts for real-world Go work and assessments. Arrays are fundamental but often misunderstood - mastering them shows deep Go knowledge!
