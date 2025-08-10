package main

import "fmt"

func main() {
	bool1 := true
	bool2 := false

	fmt.Println(bool1 && bool2)
	fmt.Println(bool1 || bool2)
	fmt.Println(!bool1)
}

/*
Output
false
true
false
*/

/*
Code Explanation:
- Purpose: Demonstrate logical operators AND, OR, NOT
- Variables: bool1 := true, bool2 := false
- Expressions:
  - bool1 && bool2 -> false
  - bool1 || bool2 -> true
  - !bool1 -> false
*/
