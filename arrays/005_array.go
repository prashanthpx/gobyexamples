package main

import "fmt"

func main() {
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a // A copy of "a" is assigned to "b"
	b[0] = "Singapore"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)
}

/*
Output
a is  [USA China India Germany France]
b is  [Singapore China India Germany France]
*/

/*
Code Explanation:
- Purpose: Show that assigning an array to a new variable copies the entire array
- a is copied to b; modifying b does not affect a
- Printing shows a unchanged and b with the updated first element
*/
