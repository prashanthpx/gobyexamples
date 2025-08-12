package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

// Run with: go run wait_group/examples/errgroup_pipeline.go
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)
	tasks := []int{1, 2, 3, 4}

	out := make(chan int)

	// Producer
	g.Go(func() error {
		defer close(out)
		for _, t := range tasks {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case out <- t:
			}
		}
		return nil
	})

	// Consumers
	for i := 0; i < 2; i++ {
		g.Go(func() error {
			for v := range out {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case <-time.After(200 * time.Millisecond):
					fmt.Print(v, " ")
				}
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil { fmt.Println("err:", err) }
}

