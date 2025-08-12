package main

// Run with: go run maps/mistakes/nil_write.go
// Demonstrates nil map write panic (uncomment to see panic)

func main() {
	var m map[string]int
	// m["x"] = 1 // panic: assignment to entry in nil map
	_ = m
}

