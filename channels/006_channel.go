package main

import (
	"fmt"
)

func prod(ch chan int) {
	for i := 0; i < 10; i++ {
		ch<-i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go prod(ch)
	fmt.Printf("ch = %v", ch)
	for v := range ch {
		fmt.Println("v = ", v)
	}
}