package main

import (
	"fmt"
	"time"
)

func main() {
	a := make(chan int)
	go run(a)
	// In this select, default gets selected always due to delay in
	// run(). If the delay is removed, then anyone can be selected
	// and depends on compiler
	select {
	case <-a:
		fmt.Println("got value")
	default:
		fmt.Println(" default case")
	}
	//<-a
	fmt.Println(" exit main")
}

func run(ch chan int) {
	time.Sleep(2 * time.Second)
	close(ch)
	fmt.Println(" channel closed")
}