package main

import (
	"fmt"
)

func number() int {
	num := 15 * 5
	return num
}

func main() {

	switch num := number(); { //num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}

}

/*
Output
75 is lesser than 100
75 is lesser than 200
*/

/*
Code Explanation:
- Purpose: Show switch with an initializer and chained fallthrough behavior
- number(): returns 15 * 5 = 75
- Switch with initializer: num := number(); switch { ... }
  - case num < 50: false for 75 (skipped), but if true would fallthrough
  - case num < 100: true for 75 -> prints "75 is lesser than 100"; fallthrough forces next case
  - case num < 200: then prints "75 is lesser than 200"
*/
