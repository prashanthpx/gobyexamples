// Initialzing map during declaration

package main

import "fmt"

func main() {
	numbers := map[string]int{
		"one": 1,
		"two": 2,
	}
	numbers["three"] = 3
	// pritning the map
	fmt.Println(" number map : ", numbers)
}

/*
Output
 number map :  map[one:1 three:3 two:2]
*/

/*
Code Explanation:
- Purpose: Initialize a map literal and add a key
- numbers := map[string]int{...} creates a map with two keys
- numbers["three"] = 3 inserts another key-value pair
- Printing shows the map (unordered key order)
*/
