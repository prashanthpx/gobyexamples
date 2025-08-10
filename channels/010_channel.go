package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // an unbuffered channel
	go func(ch chan<- int, x int) {
		time.Sleep(time.Second)
		// <-ch    // fails to compile
		// Send the value and block until the result is received.
		ch <- x * x // 9 is sent
	}(c, 3)
	done := make(chan struct{})
	go func(ch <-chan int) {
		// Block until 9 is received.
		n := <-ch
		fmt.Println(n) // 9
		// ch <- 123   // fails to compile
		time.Sleep(time.Second)
		done <- struct{}{}
	}(c)
	// Block here until a value is received by
	// the channel "done".
	<-done
	fmt.Println("bye")
}

/*
Output
9
bye
*/

/*
Code Explanation:
- Purpose: Demonstrate unbuffered channel with directional parameters and synchronization
- Producer sends x*x after 1s; consumer receives, prints, then signals done
- Main waits on done channel and prints bye
*/
