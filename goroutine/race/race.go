package main

import (
	"fmt"
	"sync"
)

// Run with: go run -race goroutine/race/race.go
func main() {
	var x int
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			x++ // data race; detected with -race
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("x=", x)
}

