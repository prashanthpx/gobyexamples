package bench

import (
	"sync"
	"sync/atomic"
	"testing"
)

type cfg struct{ A, B, C int }

// RWMutex-protected read-mostly store
type rwStore struct {
	mu  sync.RWMutex
	val cfg
}

func (s *rwStore) Load() cfg { s.mu.RLock(); v := s.val; s.mu.RUnlock(); return v }
func (s *rwStore) Store(v cfg) { s.mu.Lock(); s.val = v; s.mu.Unlock() }

// atomic.Value-based read-mostly store
type valStore struct{ v atomic.Value } // stores cfg

func newValStore() *valStore { s := &valStore{}; s.v.Store(cfg{}); return s }
func (s *valStore) Load() cfg { return s.v.Load().(cfg) }
func (s *valStore) Store(v cfg) { s.v.Store(v) }

// Run with: go test -bench=ReadMostly -benchmem ./atomic/bench -cpu=1,4

func BenchmarkRWMutexReadMostly(b *testing.B) {
	s := &rwStore{}
	s.Store(cfg{A:1,B:2,C:3})
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB){
		var i int
		for pb.Next() {
			_ = s.Load()               // read path
			i++
			if i&1023 == 0 {           // ~1 write per 1024 reads
				s.Store(cfg{A:i,B:i,C:i})
			}
		}
	})
}

func BenchmarkAtomicValueReadMostly(b *testing.B) {
	s := newValStore()
	s.Store(cfg{A:1,B:2,C:3})
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB){
		var i int
		for pb.Next() {
			_ = s.Load()               // read path
			i++
			if i&1023 == 0 {           // ~1 write per 1024 reads
				s.Store(cfg{A:i,B:i,C:i})
			}
		}
	})
}

