package main

import "fmt"

// these are called name structures
type car struct {
	colour string
	length int
	price  int
}

func main() {
	// creating a structure
	volvo := car{
		colour: "blue",
		length: 5,
		price:  10000,
	}

	// another way. Here we an see, we don't mention struct member name
	// Here the order of assignment of struct memebers has to be maintained
	benz := car{"white", 10, 20000}

	fmt.Println("volvo ", volvo)
	fmt.Println(" benz ", benz)
}

/*
Output
volvo  {blue 5 10000}
 benz  {white 10 20000}
*/

/*
Code Explanation:
- Purpose: Define a named struct and initialize via keyed and ordered literals
- car has fields colour, length, price
- volvo uses keyed fields; benz uses positional fields
*/
