package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Run with: go run goroutine/examples/fanin_context.go
// Fan-in with context cancellation; all goroutines exit cleanly on cancel.
func producer(ctx context.Context, id int, out chan<- int, start int) {
	for n := start; ; n++ {
		select {
		case <-ctx.Done():
			return
		case out <- n:
		}
		time.Sleep(20 * time.Millisecond)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()

	out := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int){
			defer wg.Done()
			producer(ctx, id, out, id*100)
		}(i)
	}

	// closer
	go func(){ wg.Wait(); close(out) }()

	for v := range out { fmt.Print(v, " ") }
	fmt.Println("\nshut down cleanly")
}

