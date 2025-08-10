package main

import (
	"fmt"
)

func main() {
	// Here we can see before the condition is met,
	// statement is executed. In this case num := 10
	if num := 10; num%2 == 0 { //checks if number is even
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}

/*
Output
10 is even
*/

/*
Code Explanation:
- Purpose: Check if a number is even or odd
- Starting values: num := 10
- Condition: num % 2 == 0
- Body: print "num is even" if true, "num is odd" if false
*/
