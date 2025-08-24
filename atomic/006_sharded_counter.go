package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Run with: go run atomic/006_sharded_counter.go
// ShardedCounter reduces contention by spreading increments across shards.

type ShardedCounter struct {
	shards []atomic.Int64
}

func NewShardedCounter(n int) *ShardedCounter {
	s := &ShardedCounter{shards: make([]atomic.Int64, n)}
	return s
}

// IncShard increments the shard at index i (0..len-1).
// In real apps, choose shard by hashing goroutine/task/owner to spread load.
func (s *ShardedCounter) IncShard(i int) { s.shards[i%len(s.shards)].Add(1) }

func (s *ShardedCounter) Load() int64 {
	var total int64
	for i := range s.shards { total += s.shards[i].Load() }
	return total
}

func main() {
	s := NewShardedCounter(16)
	var wg sync.WaitGroup
	G := 8
	for g := 0; g < G; g++ {
		wg.Add(1)
		sh := g % len(s.shards)
		go func(idx int){
			for i := 0; i < 10000; i++ { s.IncShard(idx) }
			wg.Done()
		}(sh)
	}
	wg.Wait()
	fmt.Println("count:", s.Load())
}

