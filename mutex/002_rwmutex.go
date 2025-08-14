package main

import (
	"fmt"
	"sync"
)

// Run with: go run mutex/002_rwmutex.go

type Cache struct {
	mu sync.RWMutex
	m  map[string]int
}

func (c *Cache) Get(k string) (int, bool) {
	c.mu.RLock()
	v, ok := c.m[k]
	c.mu.RUnlock()
	return v, ok
}

func (c *Cache) Set(k string, v int) {
	c.mu.Lock();
	if c.m == nil { c.m = make(map[string]int) }
	c.m[k] = v
	c.mu.Unlock()
}

func main() {
	var c Cache
	c.Set("a", 1)
	if v, ok := c.Get("a"); ok { fmt.Println("a=", v) }
}

