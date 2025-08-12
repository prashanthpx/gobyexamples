package main

import (
	"context"
	"fmt"
	"time"
)

// Run with: go run goroutine/examples/cancel_pipeline.go
// Demonstrates manual cancellation propagation in a channel pipeline.

func gen(ctx context.Context, n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			select {
			case <-ctx.Done():
				return
			case out <- i:
			}
		}
	}()
	return out
}

func square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			select {
			case <-ctx.Done():
				return
			case out <- v * v:
			}
		}
	}()
	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := gen(ctx, 10)
	sq := square(ctx, c)

	for v := range sq {
		fmt.Print(v, " ")
		if v >= 9 { // cancel early
			cancel()
		}
	}
	// All goroutines exit because they select on ctx.Done()
}

