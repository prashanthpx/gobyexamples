package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	defer func() {
		<-ch
		fmt.Printf("Exiting defer")
	}()

	val := 5
	if val < 10 {
		ch<-1
		fmt.Printf("line 16")
		return 
	}
	//ch <-1
	fmt.Printf(" line 14")
}