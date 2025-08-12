# Go Functions: Advanced Developer Guide

## **Table of Contents**
1. [Function Fundamentals](#function-fundamentals)
2. [Advanced Function Features](#advanced-function-features)
3. [Closures and Lexical Scoping](#closures-and-lexical-scoping)
4. [Higher-Order Functions](#higher-order-functions)
5. [Variadic Functions](#variadic-functions)
6. [Named Returns and Naked Returns](#named-returns-and-naked-returns)
7. [Defer, Panic, and Recover](#defer-panic-and-recover)
8. [Common Mistakes and Gotchas](#common-mistakes-and-gotchas)
9. [Best Practices](#best-practices)
10. [Performance Considerations](#performance-considerations)
11. [Advanced Challenge Questions](#advanced-challenge-questions)


Run these examples
- Loop var capture: go run functions/mistakes/loop_var_capture.go
- Defer-in-loop HTTP: go run functions/mistakes/defer_loop_http.go

---

## **Function Fundamentals**

### **What are Functions in Go?**
Functions in Go are first-class citizens - they can be:
- Assigned to variables
- Passed as arguments
- Returned from other functions
- Stored in data structures

### **Basic Syntax**
```go
func functionName(param1 type1, param2 type2) (returnType1, returnType2) {
    // function body
    return value1, value2
}
```

### **Key Characteristics**
- **Statically typed**: Parameter and return types must be declared
- **Multiple return values**: Functions can return multiple values
- **Pass by value**: Arguments are copied (except slices, maps, channels, pointers)
- **No function overloading**: Function names must be unique within a package

---

## **Advanced Function Features**

### **1. Multiple Return Values**
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Usage
result, err := divide(10, 2)
if err != nil {
    log.Fatal(err)
}
```

### **2. Function Types**
```go
// Define a function type
type Calculator func(int, int) int

// Functions that match this signature
func add(a, b int) int { return a + b }
func multiply(a, b int) int { return a * b }

// Use function type
var calc Calculator = add
result := calc(5, 3) // 8
```

### **3. Anonymous Functions**
```go
// Immediate execution
result := func(x, y int) int {
    return x + y
}(5, 3)

// Assigned to variable
add := func(x, y int) int {
    return x + y
}
```

---

## **Closures and Lexical Scoping**

### **What is a Closure?**
A closure is a function that captures variables from its surrounding lexical scope.

```go
func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

counter1 := makeCounter()
counter2 := makeCounter()

fmt.Println(counter1()) // 1
fmt.Println(counter1()) // 2
fmt.Println(counter2()) // 1 (independent closure)
```

### **Memory Implications**
```go
func createClosures() []func() int {
    var funcs []func() int
    
    // ❌ Common mistake - all closures capture same variable
    for i := 0; i < 3; i++ {
        funcs = append(funcs, func() int {
            return i // All will return 3!
        })
    }
    
    // ✅ Correct approach
    for i := 0; i < 3; i++ {
        i := i // Create new variable in each iteration
        funcs = append(funcs, func() int {
            return i
        })
    }
    
    return funcs
}
```

### **Escape Analysis**
Variables captured by closures escape to the heap:

```go
func createClosure() func() int {
    x := 42 // This will be allocated on heap due to closure
    return func() int {
        return x
    }
}
```

---

## **Higher-Order Functions**

### **Functions as Parameters**
```go
func applyOperation(a, b int, op func(int, int) int) int {
    return op(a, b)
}

// Usage
result := applyOperation(5, 3, func(x, y int) int {
    return x * y
})
```

### **Functions as Return Values**
```go
func getValidator(minLength int) func(string) bool {
    return func(s string) bool {
        return len(s) >= minLength
    }
}

isValidPassword := getValidator(8)
fmt.Println(isValidPassword("secret")) // false
```

### **Function Composition**
```go
func compose(f, g func(int) int) func(int) int {
    return func(x int) int {
        return f(g(x))
    }
}

double := func(x int) int { return x * 2 }
square := func(x int) int { return x * x }

doubleSquare := compose(double, square)
result := doubleSquare(3) // double(square(3)) = double(9) = 18
```

---

## **Variadic Functions**

### **Basic Variadic Functions**
```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Usage
fmt.Println(sum(1, 2, 3, 4)) // 10
fmt.Println(sum())           // 0
```

### **Passing Slices to Variadic Functions**
```go
numbers := []int{1, 2, 3, 4}
result := sum(numbers...) // Spread operator
```

### **Mixed Parameters**
```go
func printf(format string, args ...interface{}) {
    // format is required, args is variadic
    fmt.Printf(format, args...)
}
```

### **Variadic Function Gotchas**
```go
func modifyVariadic(nums ...int) {
    if len(nums) > 0 {
        nums[0] = 999 // This modifies the original slice!
    }
}

slice := []int{1, 2, 3}
modifyVariadic(slice...)
fmt.Println(slice) // [999, 2, 3] - original slice modified!
```

---

## **Named Returns and Naked Returns**

### **Named Return Values**
```go
func divide(a, b float64) (result float64, err error) {
    if b == 0 {
        err = errors.New("division by zero")
        return // naked return
    }
    result = a / b
    return // naked return
}
```

### **Benefits and Drawbacks**

**Benefits:**
- Self-documenting code
- Can be modified in defer functions
- Cleaner error handling patterns

**Drawbacks:**
- Can be confusing with naked returns
- Shadowing issues

### **Shadowing Gotcha**
```go
func problematic() (result int) {
    result = 10
    if true {
        result := 20 // This shadows the named return!
        _ = result
    }
    return // Returns 10, not 20
}
```

---

## **Defer, Panic, and Recover**

### **Defer Mechanics**
```go
func deferExample() {
    defer fmt.Println("First defer")
    defer fmt.Println("Second defer")
    fmt.Println("Function body")
}
// Output:
// Function body
// Second defer
// First defer
```

### **Defer with Closures**
```go
func deferLoop() {
    for i := 0; i < 3; i++ {
        defer func() {
            fmt.Println(i) // Prints 3, 3, 3
        }()
    }
    
    for i := 0; i < 3; i++ {
        defer func(val int) {
            fmt.Println(val) // Prints 2, 1, 0
        }(i)
    }
}
```

### **Panic and Recover**
```go
func safeDivide(a, b float64) (result float64, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic recovered: %v", r)
        }
    }()
    
    if b == 0 {
        panic("division by zero")
    }
    
    result = a / b
    return
}
```

### **Defer Performance**
```go
// ❌ Expensive - defer has overhead
func inefficient() {
    for i := 0; i < 1000000; i++ {
        defer func() {}()
    }
}

// ✅ Better - single defer
func efficient() {
    defer func() {
        for i := 0; i < 1000000; i++ {
            // cleanup work
        }
    }()
}
```

---

## **Common Mistakes and Gotchas**

### **1. Loop Variable Capture**
```go
// ❌ Wrong
var funcs []func()
for i := 0; i < 3; i++ {
    funcs = append(funcs, func() {
        fmt.Println(i) // All print 3
    })
}

// ✅ Correct
for i := 0; i < 3; i++ {
    i := i // Create new variable
    funcs = append(funcs, func() {
        fmt.Println(i)
    })
}
```

### **2. Defer in Loops**
```go
// ❌ Wrong - defers accumulate
func processFiles(filenames []string) {
    for _, filename := range filenames {
        file, err := os.Open(filename)
        if err != nil {
            continue
        }
        defer file.Close() // All files stay open until function returns!
        // process file
    }
}

// ✅ Correct - use anonymous function
func processFiles(filenames []string) {
    for _, filename := range filenames {
        func() {
            file, err := os.Open(filename)
            if err != nil {
                return
            }
            defer file.Close() // Closes when anonymous function returns
            // process file
        }()
    }
}
```

### **3. Nil Function Calls**
```go
var fn func()
fn() // Runtime panic: nil pointer dereference

// Safe approach
if fn != nil {
    fn()
}
```

### **4. Method Values vs Method Expressions**
```go
type Counter struct {
    count int
}

func (c *Counter) Increment() {
    c.count++
}

c := &Counter{}

// Method value - receiver is bound
increment := c.Increment
increment() // Works

// Method expression - receiver must be passed
increment2 := (*Counter).Increment
increment2(c) // Must pass receiver
```

### **5. Not closing HTTP response bodies (leaks connections)**
```go
// ❌ Leaks TCP connections if Body isn't closed
resp, err := http.Get(url)
if err != nil { return err }
// Missing: defer resp.Body.Close()

// ✅ Always close, and drain if you don't read
resp, err = http.Get(url)
if err != nil { return err }
defer resp.Body.Close()
// If you won't read the body, drain to allow connection reuse
io.Copy(io.Discard, resp.Body)
```

---

## **Best Practices**

### **1. Function Naming**
```go
// ✅ Good - verb-based, descriptive
func calculateTotalPrice(items []Item) float64
func validateEmail(email string) error
func parseJSON(data []byte) (*Config, error)

// ❌ Bad - unclear purpose
func process(data interface{}) interface{}
func handle(x, y int) int
```

### **2. Error Handling**
```go
// ✅ Good - explicit error handling
func readConfig(filename string) (*Config, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("failed to read config file %s: %w", filename, err)
    }
    
    var config Config
    if err := json.Unmarshal(data, &config); err != nil {
        return nil, fmt.Errorf("failed to parse config: %w", err)
    }
    
    return &config, nil
}
```

### **3. Function Size**
- Keep functions small and focused (single responsibility)
- Aim for functions that fit on one screen
- Extract complex logic into separate functions

### **4. Parameter Validation**
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    if math.IsNaN(a) || math.IsNaN(b) {
        return 0, errors.New("invalid input: NaN")
    }
    return a / b, nil
}
```

---

## **Performance Considerations**

### **1. Function Call Overhead**
- Function calls have overhead (stack frame creation)
- Inlining can eliminate this for small functions
- Use `go build -gcflags="-m"` to see inlining decisions

### **2. Closure Performance**
```go
// Heap allocation for captured variables
func createClosure() func() int {
    x := 42 // Escapes to heap
    return func() int {
        return x
    }
}

// Stack allocation
func regularFunction() int {
    x := 42 // Stays on stack
    return x
}
```

### **3. Defer Performance**
- Defer has ~50ns overhead per call
- Avoid defer in tight loops
- Consider manual cleanup for performance-critical code

---

## **Advanced Challenge Questions**

### **Q1: What happens with this defer?**
```go
func question1() (result int) {
    defer func() {
        result++
    }()
    return 5
}
// Answer: Returns 6 (defer modifies named return)
```

### **Q2: Explain the output**
```go
func question2() {
    for i := 0; i < 3; i++ {
        defer func() {
            fmt.Print(i)
        }()
    }
}
// Answer: Prints "333" (closure captures variable, not value)
```

### **Q3: Memory leak potential?**
```go
func question3() []func() int {
    var funcs []func() int
    data := make([]int, 1000000) // Large slice
    
    for i := 0; i < 10; i++ {
        funcs = append(funcs, func() int {
            return data[0] // Entire slice stays in memory!
        })
    }
    return funcs
}
// Answer: Yes, entire data slice is kept alive by closures
```

### **Q4: What's the difference?**
```go
// Version A
func versionA() func() {
    return func() {
        fmt.Println("A")
    }
}

// Version B  
func versionB() func() {
    f := func() {
        fmt.Println("B")
    }
    return f
}
// Answer: Functionally identical, but B might be slightly more readable
```

---

**🎯 Key Takeaways for Practitioners:**
1. **Understand closure mechanics** and variable capture
2. **Know defer execution order** and performance implications  
3. **Master error handling patterns** with multiple returns
4. **Recognize memory implications** of closures and escape analysis
5. **Practice function composition** and higher-order function patterns

This guide covers the essential function concepts needed for advanced Go work and assessments. Practice these patterns and understand the underlying mechanics!
