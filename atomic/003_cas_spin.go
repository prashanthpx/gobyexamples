package main

import (
	"fmt"
	"sync/atomic"
)

// Run with: go run atomic/003_cas_spin.go
// CAS loop to compute min value across goroutines
func main() {
	var min atomic.Int64
	min.Store(1<<62)

	update := func(v int64) {
		for {
			old := min.Load()
			if v >= old { return }
			if min.CompareAndSwap(old, v) { return }
		}
	}

	update(100)
	update(50)
	update(75)
	fmt.Println("min:", min.Load())
}

