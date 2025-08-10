package main

import "fmt"

type rect struct {
	length  int
	breadth int
	color   string
}

func main() {
	rect1 := new(rect) // rect1 is a pointer to an instance of rectangle
	rect1.length = 10
	rect1.breadth = 20
	rect1.color = "Green"
	fmt.Println(rect1)

	var rect2 = new(rect) // rect2 is an instance of rectangle
	rect2.length = 10
	rect2.color = "Red"
	fmt.Println(rect2)
}

/*
Output
&{10 20 Green}
&{10 0 Red}
*/

/*
Code Explanation:
- Purpose: Use new to allocate struct values and set fields through pointers
- rect1 := new(rect) yields *rect; fields set via pointer
- rect2 := new(rect) shows zero-value for uninitialized fields
*/
