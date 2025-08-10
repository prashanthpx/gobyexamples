package main

import (
	"fmt"
)

// Generator returns a channel that produces the numbers 1, 2, 3,…
// To stop the underlying goroutine, close the channel.
func Generator() chan int {
	ch := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				fmt.Printf(" \n line 16 ")
				n++
			case <-ch:
				fmt.Println(" line 19")
				return
			}
			fmt.Println(" line 22")
		}
	}()
	return ch
}

func main() {
	number := Generator()
	fmt.Println(<-number)
	fmt.Println(<-number)
	close(number)
	// …
}

/*
Output (non-deterministic interleaving)

 line 16 1
 line 22

 line 16  line 22
2
*/

/*
Code Explanation:
- Purpose: Simple generator goroutine sending incrementing ints on a channel
- Generator launches a goroutine that sends n on ch; prints markers; increments
- Reading twice gets two values; closing the channel causes the goroutine to return via the <-ch case
- Print output interleaves due to concurrency; exact ordering may vary
*/
