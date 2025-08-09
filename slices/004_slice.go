package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("Original Slice")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	// Creating slice s1 from [1:5]
	s1 := s[1:5]
	fmt.Println("orig s1: ", s1)

	// Now chaning first element of slice to 101
	s1[0] = 101
	fmt.Println("mod s1: ", s1)

	// Creating another slice from slice array s1. Here slice is from [2:4] of s1 and not s
	s2 := s1[2:4]
	fmt.Println("orig s2: ", s2)
	s2[0] = 999
	fmt.Println("mod s2: ", s2)

	// Above slice array values were changed using slice array but as slice always points to base array, even original
	// array content is changed
	fmt.Println(" s no: ", s)
}
