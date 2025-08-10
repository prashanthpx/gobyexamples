package main

import "fmt"

func main() {
	c := make(chan int, 2) // a buffered channel
	c <- 3
	c <- 5
	close(c)
	fmt.Println(len(c), cap(c)) // 2 2
	x, ok := <-c
	fmt.Println(x, ok)          // 3 true
	fmt.Println(len(c), cap(c)) // 1 2
	x, ok = <-c
	fmt.Println(x, ok)          // 5 true
	fmt.Println(len(c), cap(c)) // 0 2
	x, ok = <-c
	fmt.Println(x, ok) // 0 false
	x, ok = <-c
	fmt.Println(x, ok)          // 0 false
	fmt.Println(len(c), cap(c)) // 0 2
	close(c)                    // panic!
	// The send will also panic if the above
	// close call is removed.
	c <- 7
}

/*
Output (panic)
2 2
3 true
1 2
5 true
0 2
0 false
0 false
0 2
panic: close of closed channel
*/

/*
Code Explanation:
- Purpose: Explore buffered channel length/capacity, reading after close, and panics
- After close, reads drain remaining values and then return zero value with ok=false
- Closing a closed channel panics; sending on a closed channel also panics
*/
