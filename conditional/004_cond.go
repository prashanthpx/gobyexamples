package main

import (
	"fmt"
)

func main() {
	num := 75
	switch { // expression is omitted
	case num >= 0 && num <= 50:
		fmt.Println("num is greater than 0 and less than 50")
	case num >= 51 && num <= 100:
		fmt.Println("num is greater than 51 and less than 100")
	case num >= 101:
		fmt.Println("num is greater than 100")
	}
}

/*
Output
num is greater than 51 and less than 100
*/

/*
Code Explanation:
- Purpose: Demonstrate a switch statement without an explicit tag expression
- Starting value: num := 75
- Switch with omitted expression evaluates boolean cases top-down
  - case num >= 0 && num <= 50: false for 75
  - case num >= 51 && num <= 100: true -> prints message and stops (no fallthrough)
  - case num >= 101: not evaluated due to earlier match
*/
