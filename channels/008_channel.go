package main

import (
	"fmt"
	"time"
)

func consume(ch chan int) {
	for i := 0; i < 5; i++ {
		fmt.Printf(" \n line 11 val written to chan: %v", i)
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int, 2)
	go consume(ch)
	time.Sleep(2 * time.Second)

	for val := range ch {
		fmt.Printf(" \n line 23 ch val: %v", val)
		time.Sleep(2 * time.Second)
	}
}

/*
Output (timing-dependent)

 line 11 val written to chan: 0
 line 11 val written to chan: 1
 line 11 val written to chan: 2
 line 23 ch val: 0
 line 11 val written to chan: 3
 line 11 val written to chan: 4
 line 23 ch val: 1
 line 23 ch val: 2
 line 23 ch val: 3
 line 23 ch val: 4
*/

/*
Code Explanation:
- Purpose: Producer fills buffered channel; consumer ranges and sleeps between reads
- Producer writes 0..4 then closes; consumer prints each value with delay
- Interleaving varies based on timing and buffer capacity
*/
