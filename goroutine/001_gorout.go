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