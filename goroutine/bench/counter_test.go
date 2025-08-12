package counter

import (
	"sync"
	"testing"
)

// Run with: go test -bench=. -cpu=1,4 -benchtime=200ms ./goroutine/bench

func BenchmarkMutex(b *testing.B) {
	var mu sync.Mutex
	var n int
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu.Lock(); n++; mu.Unlock()
		}
	})
}

func BenchmarkChannel(b *testing.B) {
	ch := make(chan int, 1024)
	done := make(chan struct{})
	go func(){
		for {
			select {
			case <-done: return
			case <-ch:
			}
		}
	}()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() { ch <- 1 }
	})
	close(done)
}

