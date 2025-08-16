package main

import "fmt"

// Run with: go run types/003_conversions.go

type A int

type B int

func main() {
	var a A = 5
	// var b B = a // compile error
	var b B = B(a) // explicit conversion
	fmt.Println(a, b)

	// Slices of different element types are not interchangeable
	t1 := []A{1,2,3}
	// t2 := []B(t1) // compile error
	_ = t1
}

