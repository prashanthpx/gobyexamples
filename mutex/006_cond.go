package main

import (
	"fmt"
	"sync"
)

// Run with: go run mutex/006_cond.go

func main() {
	mu := &sync.Mutex{}
	cond := sync.NewCond(mu)
	ready := false

	var wg sync.WaitGroup
	wg.Add(2)

	// Waiter
	go func(){
		defer wg.Done()
		mu.Lock()
		for !ready { cond.Wait() }
		fmt.Println("proceed")
		mu.Unlock()
	}()

	// Signaler
	go func(){
		defer wg.Done()
		mu.Lock()
		ready = true
		cond.Signal()
		mu.Unlock()
	}()

	wg.Wait()
}

