package main

import "fmt"

func main() {
	numbers := map[string]int{
		"one": 1,
		"two": 2,
	}

	// to check if a key is present or not. If present, val is true
	keyPresent := "three"
	key, val := numbers[keyPresent]
	_ = key

	if val == true {
		fmt.Println(" key is present")
	} else {
		fmt.Println(" key is not present")
	}
}

/*
Output
 key is not present
*/

/*
Code Explanation:
- Purpose: Demonstrate the comma-ok idiom for map lookups
- key, val := numbers["three"]: val is false if key missing
- The value assigned to key is the zero value for int in this case (ignored)
- Branch prints that the key is not present
*/
