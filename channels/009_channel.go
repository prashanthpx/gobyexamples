package main

import (
	"fmt"
	//"time"
)

func main() {
	a := make(chan int, 2)
	a <-10
	a <-11
	close(a)
	fmt.Printf(" channel a: %v, %v", <-a, <-a) // we can read hte value after closed
}
