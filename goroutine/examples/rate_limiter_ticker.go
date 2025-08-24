package main

import (
	"fmt"
	"time"
)

// Run with: go run goroutine/examples/rate_limiter_ticker.go
// Token bucket using time.Ticker and a buffered channel as the bucket.
func main() {
	bucket := make(chan struct{}, 5) // capacity = burst size
	// fill tokens periodically
	t := time.NewTicker(50 * time.Millisecond)
	defer t.Stop()
	go func() {
		for range t.C {
			select {
			case bucket <- struct{}{}:
				// added a token
			default:
				// bucket full; drop token
			}
		}
	}()

	// use tokens to rate-limit work
	for i := 0; i < 12; i++ {
		<-bucket
		fmt.Printf("req %d at %v\n", i, time.Now().Format("15:04:05.000"))
	}
}

