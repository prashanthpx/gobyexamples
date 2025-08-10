package main

import (
	"fmt"
)

func main() {
	a := 100
	var b *int

	if b == nil {
		fmt.Println(" b is ", b)
		b = &a
		fmt.Println("b is now initialized to ", b)
	}
}

/*
Output
 b is  <nil>
b is now initialized to  0x<addr>
*/

/*
Code Explanation:
- Purpose: Show zero value of a pointer and initialization by taking an address
- b is nil until assigned &a; printing b shows an address
*/
