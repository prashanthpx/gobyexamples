/*
Pointers to structure 
*/

package main

import "fmt"

type person struct {
	name string
	age int32
}

func main() {
	/// While initializing assigning 
	p1 := &person{"jack", 100}
	fmt.Printf(" name %s, age %d\n", (*p1).name, p1.age)
}