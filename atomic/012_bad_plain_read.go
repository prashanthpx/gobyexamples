package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Run with: go run -race atomic/012_bad_plain_read.go
// Demonstrates a data race when mixing atomic writes with plain reads.
var count int64

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ { atomic.AddInt64(&count, 1) }
	}()

	go func(){
		defer wg.Done()
		var last int64
		for i := 0; i < 1_000_000; i++ {
			if count < last { // plain read -> race
				fmt.Println("saw backwards:", count, "<", last)
				return
			}
			last = count
		}
	}()

	wg.Wait()
	fmt.Println("final:", count)
}

