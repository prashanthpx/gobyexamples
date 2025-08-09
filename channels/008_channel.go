package main

import (
	"fmt"
	"time"
)

func consume(ch chan int) {
	for i := 0; i < 5; i++ {
		fmt.Printf(" \n line 11 val written to chan: %v", i)
		ch<-i
	}
	close(ch)
}

func main() {
	ch := make(chan int, 2)
	go consume(ch)
	time.Sleep(2 * time.Second)
	
	for val := range ch {
		fmt.Printf(" \n line 23 ch val: %v", val)
		time.Sleep(2* time.Second)
	}
}