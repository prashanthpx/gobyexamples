# Go Maps: Advanced Developer Guide

## **Table of Contents**
1. [Map Fundamentals](#map-fundamentals)
2. [Hash Map Implementation](#hash-map-implementation)
3. [Key Requirements and Constraints](#key-requirements-and-constraints)
4. [Map Operations](#map-operations)
5. [Zero Values and Nil Maps](#zero-values-and-nil-maps)
6. [Concurrent Access Issues](#concurrent-access-issues)
7. [Iteration and Ordering](#iteration-and-ordering)
8. [Common Mistakes and Gotchas](#common-mistakes-and-gotchas)
9. [Best Practices](#best-practices)
10. [Performance Characteristics](#performance-characteristics)
11. [Advanced Challenge Questions](#advanced-challenge-questions)


Run these examples
- Nil map write demo: go run maps/mistakes/nil_write.go
- Deterministic iteration: go run maps/examples/iterate_order.go

---

## **Map Fundamentals**

### **What is a Map in Go?**
A map is an unordered collection of key-value pairs, implemented as a hash table. Maps are reference types.

```go
// Declaration
var m map[string]int              // nil map
m = make(map[string]int)          // initialized empty map

// Map literal
scores := map[string]int{
    "Alice": 95,
    "Bob":   87,
    "Carol": 92,
}

// Short declaration
ages := map[string]int{}          // empty map
```

### **Key Characteristics**
- **Reference type**: Maps are pointers to hash table structures
- **Unordered**: No guaranteed iteration order
- **Dynamic**: Can grow and shrink at runtime
- **Zero value**: `nil` (cannot be used until initialized)
- **Comparable keys**: Keys must be comparable types

---

## **Hash Map Implementation**

### **Internal Structure**
Go's map implementation uses:
- **Buckets**: Array of buckets, each holding ~8 key-value pairs
- **Hash function**: Converts keys to bucket indices
- **Overflow buckets**: Handle hash collisions
- **Load factor**: Triggers rehashing when buckets get too full

```go
// Simplified internal structure
type hmap struct {
    count     int    // number of key-value pairs
    flags     uint8  // iterator flags, etc.
    B         uint8  // log_2 of number of buckets
    noverflow uint16 // number of overflow buckets
    hash0     uint32 // hash seed
    buckets   unsafe.Pointer // array of 2^B buckets
    oldbuckets unsafe.Pointer // previous bucket array (during growth)
    // ... more fields
}
```

### **Hash Function**
```go
func hashExample() {
    m := make(map[string]int)

    // Go uses different hash functions for different key types
    m["hello"] = 1    // String hash

    m2 := make(map[int]string)
    m2[42] = "answer" // Integer hash

    // Hash function is deterministic within a program run
    // but may change between runs (hash randomization)
}
```

### **Bucket Structure**
```go
// Each bucket contains:
// - 8 key-value pairs
// - Overflow pointer to next bucket
// - Top hash bits for quick comparison

type bmap struct {
    tophash [8]uint8  // top 8 bits of hash for each key
    // keys and values are stored separately for better alignment
    // keys: [8]keytype
    // values: [8]valuetype
    // overflow: *bmap
}
```

---

## **Key Requirements and Constraints**

### **Comparable Key Types**
```go
// ✅ Valid key types (comparable)
var m1 map[string]int           // strings
var m2 map[int]string           // integers
var m3 map[bool]int             // booleans
var m4 map[float64]string       // floats (careful with NaN!)
var m5 map[[3]int]string        // arrays
var m6 map[*int]string          // pointers

// Structs with comparable fields
type Person struct {
    Name string
    Age  int
}
var m7 map[Person]string        // comparable struct

// ❌ Invalid key types (not comparable)
// var m8 map[[]int]string      // slices
// var m9 map[map[string]int]string // maps
// var m10 map[func()]string    // functions

// Structs with non-comparable fields
type BadKey struct {
    Name string
    Tags []string  // slice makes struct non-comparable
}
// var m11 map[BadKey]string    // Compile error
```

### **Float Keys Gotcha**
```go
func floatKeyGotcha() {
    m := make(map[float64]string)

    // NaN keys are problematic
    nan1 := math.NaN()
    nan2 := math.NaN()

    m[nan1] = "first"
    m[nan2] = "second"

    fmt.Println(len(m))           // 2 (NaN != NaN)
    fmt.Println(m[nan1])          // "" (can't retrieve!)
    fmt.Println(m[math.NaN()])    // "" (can't retrieve!)

    // Each NaN is unique, but you can't look them up
}
```

### **Pointer Keys**
```go
func pointerKeys() {
    m := make(map[*int]string)

    x, y := 42, 42
    m[&x] = "x pointer"
    m[&y] = "y pointer"

    fmt.Println(len(m))     // 2 (different pointers)
    fmt.Println(m[&x])      // "x pointer"

    // Pointer values matter, not what they point to
    z := &x
    fmt.Println(m[z])       // "x pointer" (same pointer)
}
```

---

## **Map Operations**

### **Basic Operations**
```go
func basicOperations() {
    m := make(map[string]int)

    // Insert/Update
    m["key1"] = 10
    m["key2"] = 20

    // Read with zero value
    value := m["key3"]        // 0 (zero value for int)

    // Read with existence check
    value, exists := m["key1"]
    if exists {
        fmt.Println("Found:", value)
    }

    // Delete
    delete(m, "key1")

    // Length
    fmt.Println(len(m))       // 1
}
```

### **The Comma Ok Idiom**
```go
func commaOkIdiom() {
    m := map[string]int{"exists": 0}

    // Without comma ok - ambiguous
    value := m["exists"]      // 0 - exists with value 0
    value = m["missing"]      // 0 - missing, zero value

    // With comma ok - clear
    value, ok := m["exists"]
    if ok {
        fmt.Println("Key exists with value:", value)
    }

    value, ok = m["missing"]
    if !ok {
        fmt.Println("Key does not exist")
    }
}
```

### **Map Initialization Patterns**
```go
func initializationPatterns() {
    // Empty map
    m1 := make(map[string]int)
    m2 := map[string]int{}

    // With initial capacity hint
    m3 := make(map[string]int, 100)

    // Map literal
    m4 := map[string]int{
        "one":   1,
        "two":   2,
        "three": 3,
    }

    // Building from slice
    keys := []string{"a", "b", "c"}
    m5 := make(map[string]int)
    for i, key := range keys {
        m5[key] = i
    }
}
```

---

## **Zero Values and Nil Maps**

### **Nil Map Behavior**
```go
func nilMapBehavior() {
    var m map[string]int  // nil map

    // ✅ Safe operations on nil map
    fmt.Println(len(m))           // 0
    fmt.Println(m["key"])         // 0 (zero value)
    value, ok := m["key"]         // 0, false

    for k, v := range m {         // No iterations
        fmt.Println(k, v)
    }

    // ❌ Panic operations on nil map
    // m["key"] = 1               // Runtime panic!
    // delete(m, "key")           // Runtime panic!
}
```

### **Nil vs Empty Map**
```go
func nilVsEmpty() {
    var nilMap map[string]int           // nil
    emptyMap := make(map[string]int)    // empty but not nil

    fmt.Println(nilMap == nil)          // true
    fmt.Println(emptyMap == nil)        // false
    fmt.Println(len(nilMap))            // 0
    fmt.Println(len(emptyMap))          // 0

    // Both behave the same for reads
    fmt.Println(nilMap["key"])          // 0
    fmt.Println(emptyMap["key"])        // 0

    // Only empty map allows writes
    emptyMap["key"] = 1                 // ✅ Works
    // nilMap["key"] = 1                // ❌ Panic
}
```

---

## **Concurrent Access Issues**

### **Race Conditions**
```go
func raceCondition() {
    m := make(map[string]int)

    // ❌ This will cause race condition
    go func() {
        for i := 0; i < 1000; i++ {
            m[fmt.Sprintf("key%d", i)] = i
        }
    }()

    go func() {
        for i := 0; i < 1000; i++ {
            _ = m[fmt.Sprintf("key%d", i)]
        }
    }()

    // Runtime error: concurrent map read and map write
}
```

### **Safe Concurrent Access**
```go
import "sync"

// Option 1: Mutex protection
type SafeMap struct {
    mu sync.RWMutex
    m  map[string]int
}

func (sm *SafeMap) Set(key string, value int) {
    sm.mu.Lock()
    defer sm.mu.Unlock()
    sm.m[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
    sm.mu.RLock()
    defer sm.mu.RUnlock()
    value, ok := sm.m[key]
    return value, ok
}

// Option 2: sync.Map (for specific use cases)
func syncMapExample() {
    var m sync.Map

    // Store
    m.Store("key1", 42)

    // Load
    value, ok := m.Load("key1")
    if ok {
        fmt.Println(value.(int))
    }

    // LoadOrStore
    actual, loaded := m.LoadOrStore("key2", 100)

    // Delete
    m.Delete("key1")

    // Range
    m.Range(func(key, value interface{}) bool {
        fmt.Printf("%s: %v\n", key, value)
        return true // continue iteration
    })
}
```

### **When to Use sync.Map**
```go
// ✅ Good for sync.Map:
// - Read-heavy workloads
// - Keys are write-once, read-many
// - Disjoint sets of keys accessed by different goroutines

// ❌ Avoid sync.Map for:
// - Write-heavy workloads
// - Frequent updates to same keys
// - Type safety requirements (interface{} values)
```

---

## **Iteration and Ordering**

### **Random Iteration Order**
```go
func randomIteration() {
    m := map[string]int{
        "a": 1, "b": 2, "c": 3, "d": 4, "e": 5,
    }

    // Order is random and may change between runs
    for key, value := range m {
        fmt.Printf("%s: %d\n", key, value)
    }

    // Go intentionally randomizes iteration order
    // to prevent code from depending on it
}
```

### **Deterministic Iteration**
```go
func deterministicIteration() {
    m := map[string]int{
        "charlie": 3, "alice": 1, "bob": 2,
    }

    // Collect and sort keys
    keys := make([]string, 0, len(m))
    for key := range m {
        keys = append(keys, key)
    }
    sort.Strings(keys)

    // Iterate in sorted order
    for _, key := range keys {
        fmt.Printf("%s: %d\n", key, m[key])
    }
}
```

### **Iteration During Modification**
```go
func iterationDuringModification() {
    m := map[string]int{"a": 1, "b": 2, "c": 3}

    // ✅ Safe: delete current key
    for key := range m {
        if key == "b" {
            delete(m, key)
        }
    }

    // ✅ Safe: delete other keys
    for key := range m {
        if key == "a" {
            delete(m, "c")
        }
    }

    // ❌ Undefined: add keys during iteration
    for key := range m {
        if key == "a" {
            m["new"] = 100 // May or may not be visited
        }
    }
}
```

---

## **Common Mistakes and Gotchas**

### **1. Nil Map Assignment**
```go
func nilMapAssignment() {
    var m map[string]int

    // ❌ This panics
    // m["key"] = 1

    // ✅ Initialize first
    m = make(map[string]int)
    m["key"] = 1
}
```

### **2. Map Comparison**
```go
func mapComparison() {
    m1 := map[string]int{"a": 1}
    m2 := map[string]int{"a": 1}

    // ❌ Maps are not comparable
    // fmt.Println(m1 == m2) // Compile error

    // ✅ Only comparable to nil
    fmt.Println(m1 == nil) // false

    // ✅ Manual comparison
    func mapsEqual(m1, m2 map[string]int) bool {
        if len(m1) != len(m2) {
            return false
        }
        for k, v1 := range m1 {
            if v2, ok := m2[k]; !ok || v1 != v2 {
                return false
            }
        }
        return true
    }
}
```

### **3. Map Value Modification**
```go
type Person struct {
    Name string
    Age  int
}

func mapValueModification() {
    m := map[string]Person{
        "alice": {"Alice", 30},
    }

    // ❌ Cannot modify map value directly
    // m["alice"].Age = 31 // Compile error

    // ✅ Must reassign entire value
    p := m["alice"]
    p.Age = 31
    m["alice"] = p

    // ✅ Or use pointer values
    m2 := map[string]*Person{
        "alice": &Person{"Alice", 30},
    }
    m2["alice"].Age = 31 // Works with pointer
}
```

### **4. Range Variable Reuse**
```go
func rangeVariableReuse() {
    m := map[string]int{"a": 1, "b": 2, "c": 3}
    var pointers []*string

    // ❌ Wrong - all pointers point to same variable
    for key := range m {
        pointers = append(pointers, &key)
    }

    for _, p := range pointers {
        fmt.Println(*p) // Prints last key 3 times
    }

    // ✅ Correct - create new variable
    for key := range m {
        key := key // Create new variable
        pointers = append(pointers, &key)
    }
}
```

### **5. Zero Value Confusion**
```go
func zeroValueConfusion() {
    m := map[string]int{}

    // These look the same but are different
    value1 := m["missing"]     // 0 (key doesn't exist)
    m["zero"] = 0
    value2 := m["zero"]        // 0 (key exists with value 0)

    // Use comma ok to distinguish
    _, exists1 := m["missing"] // false
    _, exists2 := m["zero"]    // true
}
```

---

## **Best Practices**

- Preallocate when size is known: `make(map[K]V, n)` reduces rehash cost
- Never rely on iteration order; sort keys for deterministic output
- Treat nil maps as read-only; initialize before writes
- Avoid float keys unless normalized; beware NaN lookups



### **1. Initialize Maps Properly**
```go
// ✅ Good initialization patterns
func goodInitialization() {
    // Empty map
    m1 := make(map[string]int)

    // With capacity hint for performance
    m2 := make(map[string]int, 100)

    // Map literal for known values
    m3 := map[string]int{
        "one": 1,
        "two": 2,
    }

    // Check for nil before use
    var m4 map[string]int
    if m4 == nil {
        m4 = make(map[string]int)
    }
}
```

### **2. Use Comma Ok Idiom**
```go
// ✅ Always use comma ok for existence checks
func checkExistence(m map[string]int, key string) {
    if value, ok := m[key]; ok {
        fmt.Printf("Found %s: %d\n", key, value)
    } else {
        fmt.Printf("Key %s not found\n", key)
    }
}
```

### **3. Protect Concurrent Access**
```go
// ✅ Use mutex for concurrent access
type ConcurrentMap struct {
    mu sync.RWMutex
    m  map[string]int
}

func (cm *ConcurrentMap) Get(key string) (int, bool) {
    cm.mu.RLock()
    defer cm.mu.RUnlock()
    value, ok := cm.m[key]
    return value, ok
}

func (cm *ConcurrentMap) Set(key string, value int) {
    cm.mu.Lock()
    defer cm.mu.Unlock()
    cm.m[key] = value
}
```

### **4. Use Appropriate Key Types**
```go
// ✅ Good key types
type UserID int
type SessionToken string

func goodKeyTypes() {
    users := make(map[UserID]*User)
    sessions := make(map[SessionToken]*Session)

    // Type safety and clarity
    var userID UserID = 123
    user := users[userID]
}

// ❌ Avoid problematic key types
func avoidProblematicKeys() {
    // Avoid float keys (NaN issues)
    // m1 := make(map[float64]string)

    // Avoid large struct keys (copying overhead)
    // type LargeKey struct { data [1000]byte }
    // m2 := make(map[LargeKey]string)
}
```

---

## **Performance Characteristics**

### **Time Complexity**
- **Average case**: O(1) for get, set, delete
- **Worst case**: O(n) when all keys hash to same bucket
- **Space**: O(n) where n is number of key-value pairs

### **Memory Usage**
```go
func memoryUsage() {
    // Map overhead: ~48 bytes + bucket array
    m := make(map[string]int)

    // Each bucket: ~200 bytes (holds 8 key-value pairs)
    // Load factor: ~6.5 (rehash when bucket gets too full)

    // Memory grows in powers of 2
    for i := 0; i < 100; i++ {
        m[fmt.Sprintf("key%d", i)] = i
    }
}
```

### **Capacity Hints**
```go
func capacityHints() {
    // ✅ Provide capacity hint if known
    m1 := make(map[string]int, 1000) // Allocates appropriate buckets

    // ❌ Without hint - multiple rehashes
    m2 := make(map[string]int)
    for i := 0; i < 1000; i++ {
        m2[fmt.Sprintf("key%d", i)] = i // Multiple rehashes
    }
}
```

### **Benchmark Comparison**
```go
func BenchmarkMapAccess(b *testing.B) {
    m := make(map[string]int, 1000)
    for i := 0; i < 1000; i++ {
        m[fmt.Sprintf("key%d", i)] = i
    }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = m["key500"] // ~1-2 ns on modern hardware
    }
}
```

---

## **Advanced Challenge Questions**

### **Q1: What's the output?**
```go
func question1() {
    m := map[string]int{}

    fmt.Println(m["missing"])     // ?

    m["zero"] = 0
    fmt.Println(m["zero"])        // ?

    _, ok1 := m["missing"]
    _, ok2 := m["zero"]
    fmt.Println(ok1, ok2)         // ?
}
// Answer: 0, 0, false true
```

### **Q2: Race condition?**
```go
func question2() {
    m := make(map[int]int)

    go func() {
        for i := 0; i < 1000; i++ {
            m[i] = i
        }
    }()

    go func() {
        for i := 0; i < 1000; i++ {
            _ = m[i]
        }
    }()
}
// Answer: Yes, concurrent read/write causes race condition
```

### **Q3: Memory leak?**
```go
func question3() {
    m := make(map[string][]byte)

    for i := 0; i < 1000000; i++ {
        key := fmt.Sprintf("key%d", i)
        m[key] = make([]byte, 1024) // 1KB per entry
    }

    // Delete all entries
    for key := range m {
        delete(m, key)
    }

    // Is memory freed?
}
// Answer: Buckets remain allocated; map doesn't shrink automatically
```

### **Q4: Iteration order?**
```go
func question4() {
    m := map[string]int{"a": 1, "b": 2, "c": 3}

    for i := 0; i < 3; i++ {
        for k := range m {
            fmt.Print(k)
        }
        fmt.Println()
    }
}
// Answer: Order is random and may differ between iterations
```

### **Q5: What happens?**
```go

## Passing Maps to Functions: What Reflects Back?

Go passes map values by value, but a map is a small header pointing to shared backing storage. When you pass a map, the header is copied, but both headers point to the same data.

What reflects back to the caller
```go
func touch(m map[string]int) {
    m["x"] = 42        // insert/update: visible to caller
    delete(m, "y")     // delete: visible to caller
}

func main() {
    m := map[string]int{"y": 1}
    touch(m)
    // m == map[string]int{"x":42}
}
```

What does not reflect
Reassigning the local map variable only changes the callee’s copy of the header:
```go
func replace(m map[string]int) {
    m = map[string]int{"new": 1} // caller NOT affected
}

func main() {
    m := map[string]int{"old": 1}
    replace(m)
    // m is still {"old":1}
}
```

If you need to change which map the caller sees, return the new map or pass a pointer to the map:
```go
func replaceAndReturn(m map[string]int) map[string]int {
    return map[string]int{"new": 1}
}

func replaceViaPtr(pm *map[string]int) {
    *pm = map[string]int{"new": 1}
}
```

Gotchas with values stored in maps
- Struct values: retrieving v := m[key] gives a copy. Mutating fields on v doesn’t update the map entry unless you assign it back.
```go
type T struct{ N int }
m := map[string]T{"a": {1}}
v := m["a"]
v.N = 2          // modifies the copy
m["a"] = v       // write it back; now reflected
// or store pointers: map[string]*T
```
- Slice values: s := m[key] copies the slice header; s[i] = ... affects the same backing array. If you re-slice/append causing reallocation, assign back: m[key] = s.

Concurrency note
- Maps are not safe for concurrent writes. Guard with sync.Mutex/sync.RWMutex or use sync.Map.

TL;DR
- Mutate entries and you’ll see changes in the caller; reassign the map and you won’t — return it or pass a pointer if you need that.

func question5() {
    var m map[string]int

    fmt.Println(len(m))    // ?
    fmt.Println(m["key"])  // ?

    m["key"] = 1           // ?
}
// Answer: 0, 0, panic (assignment to nil map)
```

---

**🎯 Key Takeaways for Practitioners:**
1. **Maps are reference types** with nil zero value
2. **Keys must be comparable** - no slices, maps, or functions
3. **Iteration order is random** - Go intentionally randomizes it
4. **Not safe for concurrent access** - use mutex or sync.Map
5. **Hash collisions and load factor** affect performance

This guide covers essential map concepts for advanced Go work and assessments. Maps are fundamental to Go programming and understanding their internals shows deep language knowledge!
