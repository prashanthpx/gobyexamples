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