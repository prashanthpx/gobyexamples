package main

import "fmt"

func main() {
	elements := map[string]string{
		"drive": "ssd",
		"name":  "pk",
	}
	for i, j := range elements {
		fmt.Println(i, j)
	}

}

/*
Output
drive ssd
name pk
*/

/*
Code Explanation:
- Purpose: Iterate over a map[string]string and print key/value pairs
- range returns key and value; order is unspecified
- Prints each key followed by its value
*/
