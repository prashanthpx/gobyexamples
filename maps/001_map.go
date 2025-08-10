package main

import "fmt"

func main() {
	var age map[string]int
	if age == nil {
		fmt.Println("map is nil....")
		age = make(map[string]int)
		fmt.Println(age)
	}

	// adding items to map
	age["prashanth"] = 39
	age["kumar"] = 25
	age["krish"] = 30

	// printing map contents
	fmt.Println(age)
}

/*
Output
map is nil....
map[]
map[krish:30 kumar:25 prashanth:39]
*/

/*
Code Explanation:
- Purpose: Show zero value of a map, initialization with make, and insertion
- var age map[string]int declares a nil map; nil check succeeds
- age = make(map[string]int) allocates an empty map
- Inserting keys updates the map; printing shows key-value pairs (unordered)
*/
