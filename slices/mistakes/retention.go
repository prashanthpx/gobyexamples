package main

import (
	"fmt"
)

// Run with: go run slices/mistakes/retention.go
// Demonstrates hidden retention and the 3-index slice fix.
func main() {
	big := make([]byte, 1<<20) // 1MB
	_ = big

	// small references entire 1MB backing array
	small := big[:10]
	fmt.Println("small len,cap:", len(small), cap(small))

	// Fix: limit capacity with 3-index slicing so big can be GC'd
	small2 := big[:10:10]
	fmt.Println("small2 len,cap:", len(small2), cap(small2))
}

