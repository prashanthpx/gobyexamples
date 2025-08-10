package main

import "fmt"

func main() {
	// Creating a slice using a slice literal
	var s = []int{3, 5, 7, 9, 11, 13, 17}

	// Short hand declaration
	t := []int{2, 4, 8, 16, 32, 64}

	fmt.Println("s = ", s)
	fmt.Println("t = ", t)
}

/*
Output
s =  [3 5 7 9 11 13 17]
t =  [2 4 8 16 32 64]
*/

/*
Code Explanation:
- Purpose: Show slice literals and printing slice contents
- s and t are slices created via literals; lengths/capacities inferred
- Printing displays their elements
*/
