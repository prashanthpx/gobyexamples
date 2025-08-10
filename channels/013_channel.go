package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTimeRequest() <-chan int32 {
	r := make(chan int32)
	fmt.Printf(" line 11")
	go func() {
		// Simulate a workload.
		time.Sleep(time.Second * 3)
		r <- rand.Int31n(100)
	}()

	return r
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano())

	a, b := longTimeRequest(), longTimeRequest()
	fmt.Println(sumSquares(<-a, <-b))
}

/*
Output (random-dependent; ~3s delay)
 line 11 line 11<some_number>
*/

/*
Code Explanation:
- Purpose: Run two concurrent requests returning random ints and combine results
- longTimeRequest sleeps 3s then sends a random int; main reads both and sums squares
- print shows the indicator from each goroutine then the final value
*/
