package bench

import (
	"sync"
	"sync/atomic"
	"testing"
)

// Run with: go test -bench=. -benchmem ./atomic/bench -cpu=1,4

func BenchmarkMutexInc(b *testing.B) {
	var mu sync.Mutex
	var n int64
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu.Lock(); n++; mu.Unlock()
		}
	})
}

func BenchmarkAtomicInc(b *testing.B) {
	var n atomic.Int64
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() { n.Add(1) }
	})
}

