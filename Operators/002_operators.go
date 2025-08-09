package main

import "fmt"

func main() {
	a := 10
	b := 20
	c := 30

	if a == 1 &&
		b == 20 || c == 0 {
			fmt.Printf("inside if")
		} else {
			fmt.Printf("\ninside else")
		}
}