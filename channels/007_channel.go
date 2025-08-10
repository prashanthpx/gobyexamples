package main

import (
	"fmt"
	"time"
)

func consume(ch chan int) {
	fmt.Printf(" \n line 8 ch val: %v", <-ch)
}

func main() {
	ch := make(chan int, 2)
	go consume(ch)
	ch <- 1
	time.Sleep(3 * time.Second)
	// As this is buffered channel, the current cont4xt of main
	// is not blocked as we have one more slot to write
	fmt.Printf("\n line 10")
}

/*
Output

 line 8 ch val: 1
 line 10
*/

/*
Code Explanation:
- Purpose: Buffered channel write doesn’t block when capacity not full
- ch is buffered (cap 2); consumer reads first value; main prints while buffer has spare capacity
*/
