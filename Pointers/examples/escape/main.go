package main

import "fmt"

// Run with: go build -gcflags=-m ./Pointers/examples/escape && ./escape
// Also inspect compile output to see which variables escape.

func stackAlloc() {
	x := 42 // stays on stack
	_ = x
}

func heapAlloc() *int {
	x := 42 // escapes to heap because address returned
	return &x
}

func closureEscape() func() int {
	x := 42 // escapes to heap because captured by closure
	return func() int { return x }
}

func interfaceEscape() interface{} {
	x := 42
	var i interface{} = x // x may escape due to boxing
	return i
}

func main() {
	stackAlloc()
	p := heapAlloc()
	fmt.Println(*p)
	f := closureEscape()
	fmt.Println(f())
	fmt.Println(interfaceEscape())
}

