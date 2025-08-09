/*
Program to demonstrate pointers to structures
*/

package main

import "fmt"

type person struct {
	name string
	age int32
}
func main() {
	p1 := person{"Jack", 100}
	// assigning address of p1 to p2
	p2 := &p1
	// now accessing ps using p2
	fmt.Printf(" 1) name: %s\n", (*p2).name)
	// another way. Golang allows accessing without deferencing
	fmt.Printf(" 2) name: %s, age %d\n", p2.name, p2.age)
}