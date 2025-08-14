package main

import (
	"fmt"
	"sync"
)

// Run with: go run mutex/007_map_with_mutex.go

type SafeMap struct {
	mu sync.Mutex
	m  map[string]int
}

func (s *SafeMap) Get(k string) (int, bool) {
	s.mu.Lock(); defer s.mu.Unlock()
	v, ok := s.m[k]
	return v, ok
}

func (s *SafeMap) Set(k string, v int) {
	s.mu.Lock()
	if s.m == nil { s.m = make(map[string]int) }
	s.m[k] = v
	s.mu.Unlock()
}

func main() {
	s := &SafeMap{}
	s.Set("x", 42)
	if v, ok := s.Get("x"); ok { fmt.Println("x=", v) }
}

