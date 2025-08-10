package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"sort"
)

func main() {
	values := make([]byte, 32*1024*1024)
	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	done := make(chan struct{}) // can be buffered or not

	// The sorting goroutine
	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		// Notify sorting is done.
		done <- struct{}{}
		fmt.Printf("len: %v, cap: %v", len(done), cap(done))
	}()

	// do some other things ...

	<-done // waiting here for notification
	fmt.Println(values[0], values[len(values)-1])
}

/*
Output
len: 0, cap: 00 255
*/

/*
Code Explanation:
- Purpose: Use a (possibly unbuffered) done channel to signal completion of a long operation
- After sorting 32MB of random bytes, send on done; main waits <-done, then prints min/max byte
- len(done), cap(done) on an unbuffered channel are 0
*/
