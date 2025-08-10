/*
Program to demo annonymous structure
*/

package main

import "fmt"

func main() {
	//Here we declate annonymous stcuture and initialize the same
	building := struct {
		height, length int
		colour         string
	}{
		colour: "white",
		height: 100,
		length: 500,
	}

	fmt.Println(" building : ", building)
}

/*
Output
 building :  {100 500 white}
*/

/*
Code Explanation:
- Purpose: Use an anonymous struct literal and initialize fields
- Anonymous struct has fields height, length, colour; initialized with keyed fields
- Printing shows the composite literal values in order of fields
*/
