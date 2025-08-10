/*
Basic pointer declaration program
*/

package main

import (
	"fmt"
)

func main() {
	a := 100
	var b *int
	b = &a // storing the address of a variable
	fmt.Printf("Type of b is %T\n", b)
	fmt.Printf("address of a \n", b)
}

/*
Output
Type of b is *int
address of a
%!(EXTRA *int=0x14000102020)
*/

/*
Code Explanation:
- Purpose: Basic pointer declaration and printing type/address
- b is a *int pointing to a; %T prints the pointer type
- Second Printf misses a %p verb, so fmt prints the EXTRA diagnostic showing the pointer argument
*/
