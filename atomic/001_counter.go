package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Run with: go run atomic/001_counter.go
func main() {
	var n atomic.Int64
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(){
			for j := 0; j < 1000; j++ { n.Add(1) }
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count:", n.Load())
}

