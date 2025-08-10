/*
TO demonstrat poimter creation using new()
*/

package main

import (
	"fmt"
)

func main() {
	np := new(int) // creating pointer of type *int
	fmt.Printf("Type of np is %T, val is: %d, addr: %d\n", np, *np, np)
	// now assign new to pointer
	*np = 100
	fmt.Printf(" new value is %d\n", *np)
}

/*
Output
Type of np is *int, val is: 0, addr: <addr>
 new value is 100
*/

/*
Code Explanation:
- Purpose: Use new to allocate a zero-initialized int and work with its pointer
- *np dereferences to read/write; address print is platform dependent
*/
