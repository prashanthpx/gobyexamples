package main

import "fmt"

var counter = func(n int) chan<- chan<- int {
	requests := make(chan chan<- int)
	go func() {
		for request := range requests {
			if request == nil {
				fmt.Println(" line 10 ")
				n++ // increase
			} else {
				fmt.Println(" line 13 ")
				request <- n // take out
			}
		}
	}()

	// Implicitly converted to chan<- (chan<- int)
	return requests
}(0)

func main() {
	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 10; i++ {
			counter <- nil
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go increase1000(done)
	go increase1000(done)
	<-done
	<-done

	request := make(chan int, 1)
	counter <- request
	fmt.Println(<-request) // 2000
}

/*
Output
 line 10
... repeated ...
 line 13
20
*/

/*
Code Explanation:
- Purpose: Use a request channel-of-channels as a serialized counter
- Sending nil signals an increment; sending a response channel requests the current value
- The single goroutine serializes updates, avoiding races without a mutex
*/
