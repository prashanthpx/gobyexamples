package main

import (
	"fmt"
)

func main() {
	a := [...]int{12, 78, 50} // ... makes the compiler determine the length
	fmt.Println(a)
}

/*
Output
[12 78 50]
*/

/*
Code Explanation:
- Purpose: Use ellipsis to infer array length from literals
- a := [...]int{12, 78, 50} lets the compiler determine length (3)
- Printing shows the elements in order
*/
