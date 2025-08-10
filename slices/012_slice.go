package main

import (
	"fmt"
)

func main() {
	vals := make([]int, 5)
	for i := 0; i < 5; i++ {
		vals = append(vals, i)
	}
	fmt.Println(vals)
}

/*
Output
[0 0 0 0 0 0 1 2 3 4]
*/

/*
Code Explanation:
- Purpose: Show append on a zero-initialized slice expands it
- vals starts with len=5 zeros; appending 0..4 extends slice to len=10
- The first 5 zeros are the initial elements; appended values follow
*/
