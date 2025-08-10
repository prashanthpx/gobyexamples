package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("Original slice 's' : ", s)
	// Passing slice to function and operating on it
	AddOneToEachElement(s)

	// Pritning slice value
	fmt.Println("Value of slice 's' after function call : ", s)

}

func AddOneToEachElement(slice []int) {
	for i := range slice {
		slice[i]++
	}
}

/*
Output
Original slice 's' :  [10 20 30 40 50 60 70 80 90 100]
Value of slice 's' after function call :  [11 21 31 41 51 61 71 81 91 101]
*/

/*
Code Explanation:
- Purpose: Show that slices are references to underlying arrays when passed to functions
- AddOneToEachElement increments each element via the slice reference
- The original slice reflects the modifications
*/
