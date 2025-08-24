package bench

import (
	"sync"
	"sync/atomic"
	"testing"
)

// Epoch-based atomic counter: reset via Swap(0)
type epochCounter struct{ cur atomic.Int64 }

func (e *epochCounter) Inc() { e.cur.Add(1) }
func (e *epochCounter) Reset() { _ = e.cur.Swap(0) }

// Mutex-protected counter: reset under the lock
type mutexCounter struct {
	mu sync.Mutex
	n  int64
}

func (m *mutexCounter) Inc()  { m.mu.Lock(); m.n++; m.mu.Unlock() }
func (m *mutexCounter) Reset() { m.mu.Lock(); m.n = 0; m.mu.Unlock() }

// Run with: go test -bench=Reset -benchmem ./atomic/bench -cpu=1,4

func BenchmarkAtomicEpochReset(b *testing.B) {
	var c epochCounter
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			c.Inc()
			i++
			if i&1023 == 0 { c.Reset() }
		}
	})
}

func BenchmarkMutexReset(b *testing.B) {
	var c mutexCounter
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			c.Inc()
			i++
			if i&1023 == 0 { c.Reset() }
		}
	})
}

