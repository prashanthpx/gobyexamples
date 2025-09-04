
# Go Interview Fastpack — Extended (With Dual Examples)

> A compact but deeper guide. Each topic has **two examples**:  
> 1. **Basic** → for quick understanding.  
> 2. **Complex** → demonstrates a core/real-world use case.  

---

## Page 1 — Language Basics

### Types, Vars, Consts
Go is statically typed; variables default to zero values.  
- **Basic:**
```go
var a int = 10
b := "hello"
const Pi = 3.14
```
- **Complex (iota constants):**
```go
const (
    Read = 1 << iota
    Write
    Execute
)
fmt.Println(Read, Write, Execute) // 1 2 4
```

### Operators & Control Flow
- **Basic:**
```go
for i := 0; i < 3; i++ {
    fmt.Println(i)
}
```
- **Complex (switch & fallthrough):**
```go
n := 2
switch n {
case 1: fmt.Println("One")
case 2: fmt.Println("Two"); fallthrough
case 3: fmt.Println("Three")
}
```

### Functions & Multiple Returns
- **Basic:**
```go
func add(a, b int) int { return a+b }
```
- **Complex (error handling):**
```go
func div(a, b int) (int, error) {
    if b==0 { return 0, fmt.Errorf("divide by zero") }
    return a/b, nil
}
```

### Pointers
- **Basic:**
```go
x := 5
p := &x
*p = 10
```
- **Complex (mutating struct via pointer):**
```go
type Counter struct{ N int }
func (c *Counter) Inc() { c.N++ }
c := &Counter{}; c.Inc()
```

### Structs & Methods
- **Basic:**
```go
type Point struct{ X, Y int }
p := Point{X:1, Y:2}
```
- **Complex (methods with pointer receiver):**
```go
func (p *Point) Move(dx,dy int){ p.X+=dx; p.Y+=dy }
p := &Point{1,2}; p.Move(3,4)
```

### Interfaces
- **Basic:**
```go
type Shape interface{ Area() int }
```
- **Complex (implicit implementation):**
```go
type Square struct{ S int }
func (sq Square) Area() int { return sq.S*sq.S }
var sh Shape = Square{2}
```

---

## Page 2 — Collections & Text

### Arrays vs Slices
- **Basic:**
```go
arr := [3]int{1,2,3}
```
- **Complex (append growth):**
```go
s := []int{1,2}
s = append(s, 3,4,5)
```

### Maps
- **Basic:**
```go
m := map[string]int{"a":1}
```
- **Complex (comma-ok idiom):**
```go
if v, ok := m["missing"]; ok {
    fmt.Println(v)
} else {
    fmt.Println("not found")
}
```

### Strings & Runes
- **Basic:**
```go
s := "Go"
fmt.Println(len(s)) // 2 bytes
```
- **Complex (iterate runes):**
```go
s := "Gophér"
for i,r := range s {
    fmt.Printf("%d %c\n", i, r)
}
```

### Formatting
- **Basic:**
```go
fmt.Printf("%d", 42)
```
- **Complex (struct formatting):**
```go
type User struct{ Name string; Age int }
u := User{"Tom",30}
fmt.Printf("%+v\n", u)
```

---

## Page 3 — Concurrency Primitives

### Goroutines
- **Basic:**
```go
go fmt.Println("hi")
```
- **Complex (launch many):**
```go
for i:=0;i<5;i++{
    go func(i int){ fmt.Println(i) }(i)
}
```

### Channels & Select
- **Basic:**
```go
ch := make(chan int)
go func(){ ch<-1 }()
fmt.Println(<-ch)
```
- **Complex (select with timeout):**
```go
select {
case v:=<-ch: fmt.Println(v)
case <-time.After(1*time.Second): fmt.Println("timeout")
}
```

### WaitGroup
- **Basic:**
```go
var wg sync.WaitGroup
wg.Add(1)
go func(){ defer wg.Done() }()
wg.Wait()
```
- **Complex (multiple tasks):**
```go
for i:=0;i<3;i++{
    wg.Add(1)
    go func(){ defer wg.Done(); work() }()
}
wg.Wait()
```

### Mutex & Atomic
- **Basic:**
```go
var mu sync.Mutex; var c int
mu.Lock(); c++; mu.Unlock()
```
- **Complex (atomic counter):**
```go
var n int64
atomic.AddInt64(&n,1)
```

### Worker / Job Queues
- **Basic:**
```go
jobs := make(chan int)
```
- **Complex (fan-out workers):**
```go
for w:=0;w<3;w++{
    go func(id int){
        for j:=range jobs{ fmt.Println("worker",id,"got",j) }
    }(w)
}
```

---

## Page 4 — I/O, Files, Time, Profiles

### io.Reader / io.Writer
- **Basic:**
```go
r := strings.NewReader("hi")
buf := make([]byte,2)
r.Read(buf)
```
- **Complex (io.Copy):**
```go
io.Copy(os.Stdout, strings.NewReader("hello world"))
```

### File I/O
- **Basic:**
```go
f,_ := os.Create("a.txt"); defer f.Close()
```
- **Complex (scan file):**
```go
f,_ := os.Open("a.txt")
defer f.Close()
scanner:=bufio.NewScanner(f)
for scanner.Scan(){ fmt.Println(scanner.Text()) }
```

### Time
- **Basic:**
```go
t:=time.Now()
```
- **Complex (ticker):**
```go
ticker:=time.NewTicker(time.Second)
for i:=0;i<3;i++{ <-ticker.C; fmt.Println("tick") }
```

### pprof
- **Basic:**
```go
import _ "net/http/pprof"
```
- **Complex (run server):**
```go
go http.ListenAndServe(":6060", nil)
```

### YAML / Validation
- **Basic (yaml):**
```go
data := []byte("port: 8080")
var cfg map[string]int
yaml.Unmarshal(data,&cfg)
```
- **Complex (validator):**
```go
type Cfg struct{ Port int `validate:"gte=1,lte=65535"` }
v := validator.New(); v.Struct(Cfg{Port:8080})
```

---

## Page 5 — Errors & Testing

### Errors
- **Basic:**
```go
err := errors.New("fail")
```
- **Complex (wrap & check):**
```go
if err := do(); err!=nil {
    return fmt.Errorf("do: %w", err)
}
if errors.Is(err,targetErr){ ... }
```

### Testing
- **Basic:**
```go
func TestX(t *testing.T){ if add(1,2)!=3 { t.Fail() } }
```
- **Complex (table-driven):**
```go
cases := []struct{a,b,w int}{{1,2,3},{-1,1,0}}
for _,c:=range cases{
    if got:=add(c.a,c.b); got!=c.w{
        t.Fatalf("got %d want %d",got,c.w)
    }
}
```

---

## Page 6 — Gotchas & Best Practices

- **Range var capture**
- **Nil maps**
- **Slice sharing**
- **Defer in loops**
- **Context for cancellation**
- **Small interfaces, concrete returns**

- **Basic Gotcha (nil map):**
```go
var m map[string]int
// m["x"]=1 // panic!
```
- **Complex Gotcha (range closure):**
```go
for i:=0;i<3;i++{
    go func(){ fmt.Println(i) }() // captures same i!
}
```

---
