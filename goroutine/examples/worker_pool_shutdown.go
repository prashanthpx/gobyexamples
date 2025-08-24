package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Run with: go run goroutine/examples/worker_pool_shutdown.go
// Worker pool with bounded parallelism and graceful shutdown.
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan int)
	results := make(chan string)

	const N = 3 // workers
	var wg sync.WaitGroup
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(id int){
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case j, ok := <-jobs:
					if !ok { return }
					time.Sleep(30*time.Millisecond)
					results <- fmt.Sprintf("w%d:%d", id, j)
				}
			}
		}(i)
	}

	// results closer
	go func(){ wg.Wait(); close(results) }()

	// producer
	go func(){
		for i := 0; i < 10; i++ { jobs <- i }
		close(jobs)
	}()

	for r := range results { fmt.Print(r, " ") }
	fmt.Println("\ndone")
}

