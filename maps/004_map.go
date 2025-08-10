package main

import "fmt"

func main() {
	numbers := map[string]int{
		"one": 1,
		"two": 2,
	}
	numbers["three"] = 3
	// pritning the map
	fmt.Println(" number map : ", numbers["one"])
	fmt.Println(" number map : ", numbers["three"])
}

/*
Output
 number map :  1
 number map :  3
*/

/*
Code Explanation:
- Purpose: Read values from a map by key
- numbers["one"] returns 1; numbers["three"] returns 3
- Accessing by key returns the value type (int here)
*/
