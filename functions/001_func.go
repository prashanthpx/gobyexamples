package main

import "fmt"

func main() {
	retOne := calcResult(10, 20)
	fmt.Println(" Result = ", retOne)

	area, perimeter := calTwoResult(2.5, 6.5)
	fmt.Println(" area = ", area, "perimeter =", perimeter)
}

// Function taking two vals and retruning result
func calcResult(n1 int, n2 int) int {
	var result = n1 * n2
	return result
}

// function returning two values. Here as both the func
// parameters are of same type, we can mention type once
// atlast
// This function returns two value of type float
func calTwoResult(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

/*
Output
 Result =  200
 area =  16.25 perimeter = 18
*/

/*
Code Explanation:
- Purpose: Demonstrate single and multiple return values from functions
- calcResult multiplies two ints and returns the product
- calTwoResult returns area and perimeter for a rectangle given length and width
*/
