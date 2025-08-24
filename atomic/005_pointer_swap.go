package main

import (
	"fmt"
	"sync/atomic"
)

// Run with: go run atomic/005_pointer_swap.go

type Node struct{ V int }

func main() {
	var p atomic.Pointer[Node]
	p.Store(&Node{V:1})
	old := p.Swap(&Node{V:2})
	fmt.Println("old:", old.V, "new:", p.Load().V)
}

