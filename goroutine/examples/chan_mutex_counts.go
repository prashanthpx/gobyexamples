package main

import (
	"fmt"
	"sync"
)

// Run with: go run goroutine/examples/chan_mutex_counts.go
// Jobs delivered over a channel; workers update shared map under a mutex.
func main() {
	jobs := make(chan string)
	counts := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	// workers
	for w := 0; w < 3; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				mu.Lock()
				counts[j]++
				mu.Unlock()
			}
		}()
	}

	// producer
	go func() {
		for _, k := range []string{"a","b","a","c","b","a"} { jobs <- k }
		close(jobs)
	}()

	wg.Wait()
	fmt.Println(counts)
}

