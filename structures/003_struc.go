/*
Program to demo initializtion of few fields and omit some
*/

package main

import "fmt"

type building struct {
	height, length int
	colour string
}
func main() {
	//Here only height and colour is imitialized and length is not which by
	// default gets initialized to zero
	mall := building{ 
		colour : "white",
		height: 100,	
	}
	var t *building
	t = &building{}
	fmt.Println("\n %v+t : ", t)

	fmt.Println(" building : ", mall)
}