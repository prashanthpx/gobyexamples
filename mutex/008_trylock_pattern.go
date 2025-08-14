package main

import (
	"fmt"
	"sync"
	"time"
)

// Run with: go run mutex/008_trylock_pattern.go
// Demonstrates a TryLock-like pattern using a channel to limit concurrency.
// Note: sync.Mutex has no TryLock in the standard library.

type token struct{}

func main() {
	sem := make(chan token, 1) // capacity 1 acts like a lock

	// tryLock returns whether it acquired the token without blocking
	tryLock := func() bool {
		select {
		case sem <- token{}:
			return true
		default:
			return false
		}
	}

	unlock := func() { <-sem }

	if tryLock() {
		fmt.Println("acquired")
		unlock()
	}

	// Show non-blocking behavior
	if tryLock() {
		fmt.Println("acquired again")
		// do some work
		time.Sleep(10 * time.Millisecond)
		unlock()
	} else {
		fmt.Println("busy")
	}
}

