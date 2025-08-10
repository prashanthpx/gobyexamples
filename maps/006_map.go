package main

import "fmt"

func main() {
	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for key, val := range numbers {
		fmt.Println(" key =", key, "val =", val)
	}
}

/*
Output
 key = two val = 2
 key = three val = 3
 key = one val = 1
*/

/*
Code Explanation:
- Purpose: Iterate over a map with range
- for key, val := range numbers { ... } yields each key-value pair in unspecified order
- Prints each key and its value
*/
