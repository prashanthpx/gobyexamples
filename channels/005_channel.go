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
		ch <- 1
		fmt.Printf("line 16")
		return
	}
	//ch <-1
	fmt.Printf(" line 14")
}

/*
Output (panic: deadlock)
fatal error: all goroutines are asleep - deadlock!

... stack trace omitted ...
*/

/*
Code Explanation:
- Purpose: Show deadlock when sending on an unbuffered channel without a concurrent receiver
- ch<-1 blocks; the deferred receive runs only on return, so main blocks before defer executes
- Fix by receiving concurrently (e.g., a goroutine) or making the channel buffered
*/
