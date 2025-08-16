package main

import "fmt"

// Run with: go run types/007_generics_constraints.go

type Number interface{ ~int | ~int64 | ~float64 }

func Sum[T Number](xs []T) T {
	var z T
	for _, v := range xs { z += v }
	return z
}

func main() {
	fmt.Println(Sum([]int{1,2,3}))
	fmt.Println(Sum([]int64{1,2,3}))
	fmt.Println(Sum([]float64{1,2,3}))
}

