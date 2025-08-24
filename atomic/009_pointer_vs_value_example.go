package main

import (
	"fmt"
	"sync/atomic"
)

// Run with: go run atomic/009_pointer_vs_value_example.go
// Compare atomic.Pointer[T] vs atomic.Value use cases.

type cfg struct{ Name string; N int }

func main() {
	// Using atomic.Pointer[*cfg] for pointer-heavy cases
	var p atomic.Pointer[cfg]
	p.Store(&cfg{Name:"p", N:1})
	old := p.Swap(&cfg{Name:"p2", N:2})
	fmt.Println("ptr old:", old.Name, "new:", p.Load().Name)

	// Using atomic.Value for whole-value snapshots (non-pointer)
	var v atomic.Value
	v.Store(cfg{Name:"v", N:1})
	v.Store(cfg{Name:"v2", N:2})
	fmt.Println("val:", v.Load().(cfg).Name)
}

