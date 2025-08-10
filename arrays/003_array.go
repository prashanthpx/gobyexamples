package main

import (
	"fmt"
)

func main() {
	a := [3]int{12} // assigned only firdt element, rest all
	// assgined default 0
	fmt.Println(a)
}

/*
Output
[12 0 0]
*/

/*
Code Explanation:
- Purpose: Demonstrate partial array initialization with defaults
- a := [3]int{12} sets the first element to 12; others default to 0
- Printing shows remaining elements as zero values
*/
