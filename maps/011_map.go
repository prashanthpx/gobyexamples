package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	processedMap      = make(map[string]map[string]bool)
	processedMapMutex sync.Mutex
)

func main() {
	key := "cluster-1"

	// Goroutine 1: initializes and writes to map
	go func() {
		for {
			processedMapMutex.Lock()
			processedMap[key] = make(map[string]bool) // Overwrites inner map
			fmt.Println("G1: initialized inner map")
			processedMapMutex.Unlock()
			time.Sleep(10 * time.Millisecond)
		}
	}()

	// Goroutine 2: writes to inner map
	go func() {
		for {
			processedMapMutex.Lock()
			processedMap[key]["controller-A"] = true // Writing to inner map
			fmt.Println("G2: wrote to inner map")
			processedMapMutex.Unlock()
			time.Sleep(10 * time.Millisecond)
		}
	}()

	// Let it run
	select {}
}

/*
Output (non-terminating; sample)
G1: initialized inner map
G2: wrote to inner map
G1: initialized inner map
G2: wrote to inner map
... repeats indefinitely ...
*/

/*
Code Explanation:
- Purpose: Demonstrate concurrent map access with synchronization and logical overwrite behavior
- Two goroutines run in loops protected by a mutex:
  - G1: acquires lock and reinitializes processedMap[key] to a new empty inner map
  - G2: acquires lock and writes processedMap[key]["controller-A"] = true
- Because G1 overwrites the inner map repeatedly, G2’s writes can be lost between iterations
- The mutex ensures there is no data race on the maps, but the logic still leads to non-deterministic output and lost updates
- The program blocks forever (select {}), so it prints continuously
*/
