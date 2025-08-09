package main
import "fmt"

func main() {
	numbers := map[string]int {
		"one": 1,
		"two": 2,
		"three": 3,
	}

	for key, val := range numbers {
		fmt.Println(" key =", key, "val =", val)
	}
}