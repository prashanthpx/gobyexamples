package bench

import (
	"sync"
	"testing"
)

type counter struct {
	mu sync.Mutex
	n  int
}

func (c *counter) inc() { c.mu.Lock(); c.n++; c.mu.Unlock() }
func (c *counter) get() int { c.mu.Lock(); defer c.mu.Unlock(); return c.n }

type rwcounter struct {
	mu sync.RWMutex
	n  int
}

func (c *rwcounter) inc() { c.mu.Lock(); c.n++; c.mu.Unlock() }
func (c *rwcounter) get() int { c.mu.RLock(); defer c.mu.RUnlock(); return c.n }

func BenchmarkMutex80R20W(b *testing.B) {
	var c counter
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%5 == 0 { c.inc() } else { _ = c.get() }
			i++
		}
	})
}

func BenchmarkRWMutex80R20W(b *testing.B) {
	var c rwcounter
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%5 == 0 { c.inc() } else { _ = c.get() }
			i++
		}
	})
}

func BenchmarkMutex50R50W(b *testing.B) {
	var c counter
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%2 == 0 { c.inc() } else { _ = c.get() }
			i++
		}
	})
}

func BenchmarkRWMutex50R50W(b *testing.B) {
	var c rwcounter
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%2 == 0 { c.inc() } else { _ = c.get() }
			i++
		}
	})
}

