package main

import "fmt"

type rect struct {
	length  int
	breadth int
	color   string
}

func main() {
	var rect1 = &rect{10, 20, "Green"} // Can't skip any value
	fmt.Println(rect1)

	var rect2 = &rect{}
	rect2.length = 10
	rect2.color = "Red"
	fmt.Println(rect2) // breadth skipped

	var rect3 = &rect{}
	(*rect3).breadth = 10
	(*rect3).color = "Blue"
	fmt.Println(rect3) // length skipped
}

/*
Output
&{10 20 Green}
&{10 0 Red}
&{0 10 Blue}
*/

/*
Code Explanation:
- Purpose: Different ways to initialize pointers to structs
- &rect{...} uses a composite literal; &rect{} then assign fields; (*rect3).field uses explicit deref
- Shows missing fields get zero values
*/
