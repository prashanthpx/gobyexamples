package main

import (
	"fmt"
	"sync"
)

// Run with: go run mutex/009_embedded_mutex.go

type Counter struct {
	sync.Mutex // embedded, unexported by virtue of the struct being unexported here
	n int
}

func (c *Counter) Inc() { c.Lock(); c.n++; c.Unlock() }
func (c *Counter) Get() int { c.Lock(); defer c.Unlock(); return c.n }

func main() {
	var c Counter
	c.Inc(); c.Inc()
	fmt.Println(c.Get())
}

