package main

import (
	"fmt"
	"sort"
)

// Run with: go run maps/examples/iterate_order.go
func main() {
	m := map[int]string{2: "two", 1: "one", 3: "three"}
	fmt.Print("iteration: ")
	for k, v := range m { fmt.Printf("%d=%s ", k, v) }
	fmt.Println()

	// Deterministic order via sorted keys
	keys := make([]int, 0, len(m))
	for k := range m { keys = append(keys, k) }
	sort.Ints(keys)
	fmt.Print("sorted:    ")
	for _, k := range keys { fmt.Printf("%d=%s ", k, m[k]) }
	fmt.Println()
}

