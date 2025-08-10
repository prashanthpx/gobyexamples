/*
Program to demonstrate modifying passed values using pointer receivers
*/
package main

import "fmt"

type dimension struct {
	len, width int
}

func (d *dimension) modify() {
	fmt.Printf(" Received dimension: len %d width %d\n ", d.len, d.width)

	// Now lets modify the value
	d.len = 100
	d.width = 200

}

func main() {
	shape := dimension{10, 20}
	shape.modify()
	// After modification
	fmt.Printf(" mod values: len %d width %d\n ", shape.len, shape.width)
}

/*
Output
 Received dimension: len 10 width 20
  mod values: len 100 width 200

*/

/*
Code Explanation:
- Purpose: Modify struct fields using a pointer receiver method
- dimension.modify updates len and width through the pointer
- After method call, the original struct reflects new values
*/
