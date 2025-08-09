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
	shape := dimension{10,20 }
	shape.modify()
	// After modification
	fmt.Printf(" mod values: len %d width %d\n ", shape.len, shape.width)
}