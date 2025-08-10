package main

import (
	"fmt"
)

func prod(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
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

/*
Output
ch = 0x14000020150v =  0
v =  1
v =  2
v =  3
v =  4
v =  5
v =  6
v =  7
v =  8
v =  9
*/

/*
Code Explanation:
- Purpose: Produce ints on a channel and range over them until close
- prod sends 0..9 then closes; for v := range ch reads until channel is closed
*/
