package main

import (
	"fmt"
	"sync"
)

// Run with: go run mutex/004_copy_mutex.go

type Box struct {
	mu sync.Mutex
	v  int
}

func (b *Box) Set(x int) { b.mu.Lock(); b.v = x; b.mu.Unlock() }

func main() {
	b := &Box{}
	b.Set(1)

	// ❌ Copying a struct with a mutex after use — don't do this
	copy := *b // copies the mutex internal state!

	// Using either b or copy now is unsafe; this is just to illustrate
	fmt.Println("copied value:", copy.v)

	// ✅ Prefer passing pointers; avoid copying
}

