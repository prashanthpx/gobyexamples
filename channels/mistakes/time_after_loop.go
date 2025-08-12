package main

import (
	"context"
	"fmt"
	"time"
)

// Demonstrates time.After leak pattern in loops and the fixed Ticker version.
// Run with: go run channels/mistakes/time_after_loop.go

func bad() {
	for i := 0; i < 3; i++ {
		select {
		case <-time.After(10 * time.Millisecond):
			fmt.Print("tick ")
		}
	}
	fmt.Println("(bad done)")
}

func good(ctx context.Context) {
	t := time.NewTicker(10 * time.Millisecond)
	defer t.Stop()
	for i := 0; i < 3; i++ {
		select {
		case <-t.C:
			fmt.Print("tick ")
		case <-ctx.Done():
			return
		}
	}
	fmt.Println("(good done)")
}

func main() {
	bad()
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	good(ctx)
}

