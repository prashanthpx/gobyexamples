package main

import (
	"fmt"
	"math"
)

// Run with: go run types/001_basic_types.go
func main() {
	var b bool
	var i int
	var u uint
	var f float64
	var c complex128
	var by byte
	var r rune
	var s string
	fmt.Printf("zeroes: %v %v %v %v %v %v %v %q\n", b, i, u, f, c, by, r, s)

	fmt.Println("NaN!=NaN:", math.NaN() == math.NaN())
}

