package main
import "fmt"

// Demonstrating mak() to build slice 
func main() {
	slice := make([]int, 10, 20)
	fmt.Println("len = ", len(slice), "cap = ", cap(slice))

	// modiying the capacity
	slice = make([]int, 10, 25)
	fmt.Println("modified slice, len = ", len(slice), "cap = ", cap(slice))
}