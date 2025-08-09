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