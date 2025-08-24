package bench

import (
	"hash/fnv"
	"sync/atomic"
	"testing"
)

type sharded struct{ shards []atomic.Int64 }

func newSharded(n int) *sharded { return &sharded{shards: make([]atomic.Int64, n)} }
func (s *sharded) add(key string) {
	h := fnv.New32a(); _, _ = h.Write([]byte(key))
	idx := int(h.Sum32()) % len(s.shards)
	s.shards[idx].Add(1)
}

func (s *sharded) total() int64 {
	var t int64
	for i := range s.shards { t += s.shards[i].Load() }
	return t
}

// Run with: go test -bench=Shard -benchmem ./atomic/bench -cpu=1,4
func BenchmarkShardedAdd(b *testing.B) {
	s := newSharded(64)
	keys := []string{"a","b","c","d","e","f","g","h"}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB){
		for pb.Next(){
			for _, k := range keys { s.add(k) }
		}
	})
	_ = s.total()
}

