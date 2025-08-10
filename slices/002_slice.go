package main

import "fmt"

func main() {
	var a = [5]string{"Alpha", "Beta", "Gamma", "Delta", "Epsilon"}

	// Creating a slice from the array
	var s []string = a[1:4]

	fmt.Println("Array a = ", a)
	fmt.Println("Slice s = ", s)
}

/*
Output
Array a =  [Alpha Beta Gamma Delta Epsilon]
Slice s =  [Beta Gamma Delta]
*/

/*
Code Explanation:
- Purpose: Create a slice from an array using slicing
- a is an array of 5 strings; s := a[1:4] references elements 1..3
- Printing shows the original array and the slice view
*/
