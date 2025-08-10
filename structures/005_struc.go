/*
Pointers to structure
*/

package main

import "fmt"

type person struct {
	name string
	age  int32
}

func main() {
	/// While initializing assigning
	p1 := &person{"jack", 100}
	fmt.Printf(" name %s, age %d\n", (*p1).name, p1.age)
}

/*
Output
 name jack, age 100
*/

/*
Code Explanation:
- Purpose: Initialize and use a pointer to a struct literal
- p1 := &person{"jack", 100} creates a pointer to a person value
- Accessing fields via (*p1).name or p1.age demonstrates pointer field access
*/
