package main

import "fmt"

// Run with: go run types/008_comparable_set.go

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(v T) { s[v] = struct{}{} }
func (s Set[T]) Has(v T) bool { _, ok := s[v]; return ok }

func main() {
	s := make(Set[string])
	s.Add("a"); s.Add("b")
	fmt.Println(s.Has("a"), s.Has("z"))
}

