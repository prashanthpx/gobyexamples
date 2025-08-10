/*
Program to demo initializtion of few fields and omit some
*/

package main

import "fmt"

type building struct {
	height, length int
	colour         string
}

func main() {
	//Here only height and colour is imitialized and length is not which by
	// default gets initialized to zero
	mall := building{
		colour: "white",
		height: 100,
	}
	var t *building
	t = &building{}
	fmt.Println("\n %v+t : ", t)

	fmt.Println(" building : ", mall)
}

/*
Output

 %v+t :  &{0 0 }
 building :  {100 0 white}
*/

/*
Code Explanation:
- Purpose: Show partial struct initialization and zero-values
- mall initializes only colour and height; omitted fields default to zero
- t := &building{} creates a zero-value struct pointer
*/
