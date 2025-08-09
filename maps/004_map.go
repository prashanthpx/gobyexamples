package main
import "fmt"

func main() {
	numbers := map[string]int {
		"one": 1,
		"two": 2,
	}
	numbers["three"] = 3
	// pritning the map
	fmt.Println(" number map : ", numbers["one"])
	fmt.Println(" number map : ", numbers["three"])
}