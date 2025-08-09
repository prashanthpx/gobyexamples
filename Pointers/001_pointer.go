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