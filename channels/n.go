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

/*
Output (timing-dependent)
 default case
 exit main
*/

/*
Code Explanation:
- Purpose: Demonstrate non-blocking select with default when no channel case is ready
- run sleeps 2s then closes ch; the select runs first and chooses default
- If the sleep is removed, selection between cases becomes race/timing-dependent
*/
