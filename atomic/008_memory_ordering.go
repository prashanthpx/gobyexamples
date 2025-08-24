package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// Run with: go run atomic/008_memory_ordering.go
// Demonstrates publish/subscribe with atomics establishing happens-before.
func main() {
	var data atomic.Int64
	var ready atomic.Bool

	// Publisher
	go func() {
		data.Store(42)      // write the data first
		ready.Store(true)   // publish readiness (release)
	}()

	// Subscriber
	for !ready.Load() {   // acquire loop (spins briefly)
		time.Sleep(time.Millisecond)
	}
	v := data.Load()      // guaranteed to observe 42 after seeing ready=true
	fmt.Println("data:", v)
}

