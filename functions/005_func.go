package main

import "fmt"

func main() {
	slice := []int{10, 20, 30}

	// when we want to pass a slice to function, we inclide
	// ... after the slice
	print(slice...)
	// printing the modified slice
	fmt.Println(slice)

}

func print(para ...int) {
	fmt.Println(" Inside func print")
	fmt.Println(para)
	// modiyfing the slice
	para[0] = 100
}

/*
Output
 Inside func print
[10 20 30]
[100 20 30]
*/

/*
Code Explanation:
- Purpose: Passing a slice into a variadic function uses the same backing array
- para shares the backing array with slice, so modifying para[0] updates caller’s slice
- Demonstrates that variadic parameter behaves like a slice
*/
