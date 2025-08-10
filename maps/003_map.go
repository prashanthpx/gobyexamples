package main

import "fmt"

func main() {
	numbers := map[int]string{
		1: "one",
		2: "two",
	}

	fmt.Println("numbers :", numbers)
}

/*
Output
numbers : map[1:one 2:two]
*/

/*
Code Explanation:
- Purpose: Demonstrate a map with non-string keys
- numbers := map[int]string{1:"one", 2:"two"} uses int keys
- Printing shows the mapping (unordered)
*/
