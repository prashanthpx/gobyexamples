package main

import "fmt"

func main() {
	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f \n", area, perimeter)
}

// Named return values. Returns area and perimeter
func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}

/*
Output
Area 60.480000 Perimeter 32.800000
*/

/*
Code Explanation:
- Purpose: Show named return values with an implicit return
- rectProps names area and perimeter in the signature and returns without args
- Go fills the named return variables when hitting bare return
*/
