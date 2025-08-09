// Demonstrating map delete
package main
import "fmt"

func main() {
	numbers := map[string]int {
		"one": 1,
		"two": 2,
		"three": 3,
	}

	// deleting a entry
	delete(numbers, "two")
	fmt.Println(numbers)

}