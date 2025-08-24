package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// Run with: go run atomic/011_periodic_reset.go
// Periodically reset a counter safely by swapping epochs.

type epochCounter struct {
	cur atomic.Int64
	old atomic.Int64
}

func (e *epochCounter) Inc() { e.cur.Add(1) }

func (e *epochCounter) SnapshotAndReset() int64 {
	// Swap epochs: move current to old and zero current
	prev := e.cur.Swap(0)
	e.old.Store(prev) // keep last window if needed
	return prev
}

func main() {
	var ec epochCounter
	// incrementing goroutines
	stop := make(chan struct{})
	for i := 0; i < 4; i++ {
		go func(){ for { select { case <-stop: return; default: ec.Inc(); time.Sleep(time.Millisecond) } } }()
	}

	t := time.NewTicker(100 * time.Millisecond)
	for i := 0; i < 5; i++ {
		<-t.C
		fmt.Println("window count:", ec.SnapshotAndReset())
	}
	t.Stop(); close(stop)
}

