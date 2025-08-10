// Demonstrating map delete
package main

import "fmt"

func main() {
	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	// deleting a entry
	delete(numbers, "two")
	fmt.Println(numbers)

}

/*
Output
map[one:1 three:3]
*/

/*
Code Explanation:
- Purpose: Delete a key from a map
- delete(numbers, "two") removes the entry with key "two"
- Printing shows the remaining entries
*/
