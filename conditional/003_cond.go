package main

import (
	"fmt"
)

func main() {
	switch finger := 8; finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default: //default case
		fmt.Println("incorrect finger number")
	}
}

/*
Output
incorrect finger number
*/

/*
Code Explanation:
- Purpose: Print the finger name corresponding to a given number
- Starting value: finger := 8 (set in the switch initializer)
- Switch on value: Compares finger against cases 1..5
  - 1 -> "Thumb"
  - 2 -> "Index"
  - 3 -> "Middle"
  - 4 -> "Ring"
  - 5 -> "Pinky"
- Default: Executed for any other value (prints "incorrect finger number")
*/
