package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("Original slice 's' : ", s)
	// Passing slice to function and operating on it
	AddOneToEachElement(s)

	// Pritning slice value
	fmt.Println("Value of slice 's' after function call : ", s)

}

func AddOneToEachElement(slice []int) {
    for i := range slice {
        slice[i]++
    }
}
