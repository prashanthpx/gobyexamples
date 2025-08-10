package main

import "fmt"

func main() {
	a := 10
	b := 20
	c := 30

	if a == 1 &&
		b == 20 || c == 0 {
		fmt.Printf("inside if")
	} else {
		fmt.Printf("\ninside else")
	}
}

/*
Output
inside else
*/

/*
Code Explanation:
- Purpose: Show operator precedence and short-circuit evaluation in a compound condition
- Variables: a := 10, b := 20, c := 30
- Condition: if a == 1 && b == 20 || c == 0
  - == has higher precedence than &&, which has higher precedence than ||
  - a == 1 -> false
  - false && (b == 20) -> false (short-circuits &&)
  - false || (c == 0) -> false (since c == 30)
- Result: else branch executes, printing "inside else"
*/
