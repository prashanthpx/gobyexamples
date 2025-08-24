package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// Run with: go run atomic/002_flag.go
func main() {
	var stop atomic.Bool
	go func(){
		for !stop.Load() {
			fmt.Print(".")
			time.Sleep(50*time.Millisecond)
		}
		fmt.Println("\nstopped")
	}()
	time.Sleep(200*time.Millisecond)
	stop.Store(true)
	time.Sleep(50*time.Millisecond)
}

