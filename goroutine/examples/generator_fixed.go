package main

import (
	"context"
	"fmt"
	"time"
)

// Run with: go run goroutine/examples/generator_fixed.go
// Properly cancellable generator: sender closes the channel; receiver never closes it.
func Generator(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		n := 1
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- n:
				n++
			}
		}
	}()
	return ch
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Millisecond)
	defer cancel()
	ch := Generator(ctx)
	for v := range ch { fmt.Print(v, " ") }
}

