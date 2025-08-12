package main

import "fmt"

// Run with: go run Pointers/examples/nilmethod/main.go

type List struct{ next *List }

func (l *List) Len() int {
	if l == nil { return 0 }
	return 1 + l.next.Len()
}

func main() {
	var l *List
	fmt.Println(l.Len()) // 0, method on nil pointer is OK by design
}

