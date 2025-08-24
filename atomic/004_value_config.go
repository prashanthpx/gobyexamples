package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// Run with: go run atomic/004_value_config.go
// Read-mostly config using atomic.Value (immutable snapshots)

type config struct{
	Name string
	Rate int
}

func main() {
	var cfg atomic.Value
	cfg.Store(config{Name:"init", Rate:1})

	// reader goroutine
	done := make(chan struct{})
	go func(){
		for {
			select {
			case <-done: return
			default:
				c := cfg.Load().(config)
				_ = c
				// read without locks
			}
		}
	}()

	// writer swaps entire snapshot occasionally
	for i := 0; i < 3; i++ {
		time.Sleep(50*time.Millisecond)
		cfg.Store(config{Name: fmt.Sprintf("v%d", i), Rate: i})
	}
	close(done)
	fmt.Println("final:", cfg.Load().(config))
}

