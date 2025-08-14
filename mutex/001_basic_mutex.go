package main

import (
	"fmt"
	"sync"
)

// Run with: go run mutex/001_basic_mutex.go

type Counter struct {
	mu sync.Mutex
	n  int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.n++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

func main() {
	var c Counter // zero value mutex is ready to use
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() { c.Inc(); wg.Done() }()
	}
	wg.Wait()
	fmt.Println("count:", c.Value())
}

