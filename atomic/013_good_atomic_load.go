package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Run with: go run -race atomic/013_good_atomic_load.go
// Correct approach: use atomic.Load to read while others update atomically.
var count atomic.Int64

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ { count.Add(1) }
	}()

	go func(){
		defer wg.Done()
		var last int64
		for i := 0; i < 1_000_000; i++ {
			v := count.Load()
			if v < last { fmt.Println("unexpected backwards:", v, "<", last); return }
			last = v
		}
	}()

	wg.Wait()
	fmt.Println("final:", count.Load())
}

