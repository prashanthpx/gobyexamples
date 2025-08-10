package main

import "fmt"

// Demonstrating mak() to build slice
func main() {
	slice := make([]int, 10, 20)
	fmt.Println("len = ", len(slice), "cap = ", cap(slice))

	// modiying the capacity
	slice = make([]int, 10, 25)
	fmt.Println("modified slice, len = ", len(slice), "cap = ", cap(slice))
}

/*
Output
len =  10 cap =  20
modified slice, len =  10 cap =  25
*/

/*
Code Explanation:
- Purpose: Demonstrate make for slice creation and adjusting capacity
- make([]int, 10, 20) creates a slice with len=10, cap=20
- Reassigning with make can change capacity; elements are zeroed in new allocation
*/
