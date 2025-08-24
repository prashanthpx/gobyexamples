# Atomics in Go (sync/atomic): Practical Guide with Runnable Examples

Run these examples
- Memory ordering (publish/subscribe): go run atomic/008_memory_ordering.go
- Pointer vs Value trade-offs: go run atomic/009_pointer_vs_value_example.go

- Atomic counter: go run atomic/001_counter.go
- Atomic flag/shutdown: go run atomic/002_flag.go
- CAS loop (min/once-style): go run atomic/003_cas_spin.go
- Read-mostly config with atomic.Value: go run atomic/004_value_config.go
- Pointer swap with atomic.Pointer[*T]: go run atomic/005_pointer_swap.go
- Sharded counter: go run atomic/006_sharded_counter.go
- Versioned pointer (ABA mitigation): go run atomic/007_aba_versioned_pointer.go
- Mutex vs atomic inc, sharded adds, resets (bench): go test -bench=. -benchmem ./atomic/bench
- Atomic threshold check: go run atomic/010_add_check_threshold.go
- Periodic reset via Swap (epochs): go run atomic/011_periodic_reset.go

---

## Table of Contents
1. What atomics are (and aren’t)
2. Go memory model and atomic happens-before
3. The two APIs: function-style and typed atomics
4. Common patterns: counters, flags, CAS loops
5. Read-mostly data: atomic.Value and atomic.Pointer
6. Atomics vs Locks: when to choose which
7. Common mistakes and gotchas
8. Best practices
9. Performance notes
10. Advanced interview questions

---

## 1) What atomics are (and aren’t)

- Atomics provide lock-free, word-sized operations with well-defined ordering
- They are great for single variables (counters, flags, pointers)
- They are not a drop-in replacement for locks when multiple fields must be updated atomically or invariants span several values

---

## 2) Go memory model and atomic happens-before

- sync/atomic operations establish ordering (happens-before) between goroutines
- Loads/Stores from sync/atomic include the necessary acquire/release semantics; do not mix atomic and non-atomic access to the same variable
- For compound state, prefer locks or copy-on-write with atomics pointing to immutable data

---

## 3) Two APIs

Function-style (available since early Go):
- atomic.AddInt64, atomic.LoadUint64, atomic.StorePointer, atomic.CompareAndSwapUint32, etc.

Typed atomics (Go 1.19+):
- atomic.Int64, atomic.Uint64, atomic.Bool, atomic.Pointer[T], etc. with methods Load/Store/Add/CAS/Swap

See: atomic/001_counter.go, 002_flag.go, 005_pointer_swap.go

---

## 4) Common patterns

- Counters: use Add/Load; avoid non-atomic increments across goroutines
- Threshold check after increment: use Add()’s return value to branch without races
- Flags: atomic.Bool for shutdown/ready; compare with channels for cancellation signaling
- CAS loop: implement min/max, once-like initialization, lock-free state machines (simple ones)

See: atomic/001_counter.go, 002_flag.go, 003_cas_spin.go, 010_add_check_threshold.go

---

## 5) Read-mostly data

- atomic.Value holds any value (store/load) with guaranteed consistency (readers see whole values)
- Store immutable snapshots and swap them; readers Load and read without locks
- atomic.Pointer[T] to swap pointers to immutable structs with zero-copy

See: atomic/004_value_config.go, 005_pointer_swap.go

---

## 6) Atomics vs Locks (when to use which)

Use atomics when:
- Simple counters, gauges, flags; one variable at a time
- Fast paths avoiding lock contention/overhead
- Copy-on-write snapshots so readers only Load a pointer/value

Use locks when:
- You must update multiple fields together or enforce invariants across values
- The critical section spans more than a single atomic op or includes I/O/complex logic
- You need mutual exclusion across a sequence of operations

Note: Under very high contention, a single atomic can still bottleneck (cache-line ping-pong). Shard hot counters or consider a mutex/RWMutex — measure both.

---

## 7) Common mistakes and gotchas

Mistakes Index (quick scan)
- Mixing atomic and non-atomic access to same variable
- Piecemeal updates to multi-field structs with atomics (use locks or snapshots)
- Busy CAS loops without backoff (burn CPU)
- Expecting atomic.Value to merge fields (it swaps whole snapshots)
- Copying typed atomics after use (undefined behavior)
- Using atomics where a lock is clearer/faster under contention


1) Mixing atomic and non-atomic access to the same variable (data race)
2) Using atomics to mutate multi-field structs piecemeal (torn writes/inconsistent reads)
3) Spinning with CAS without backoff/yield (burning CPU)
4) Assuming atomic.Value merges fields (it swaps whole values; make snapshots immutable)
5) Copying typed atomics after first use (undefined behavior) — never copy atomic.Int64/Bool/etc once in use
6) Building complex algorithms with atomics when a simple lock would be clearer/faster under contention

---

## 8) Best practices

- Prefer typed atomics (atomic.Int64/Bool/Pointer[T]) for clarity and type safety
- Use immutable snapshots with atomic.Value / atomic.Pointer to avoid reader locks
- Encapsulate atomics in your types; expose methods, not raw fields
- Document whether fields are updated atomically or under a lock
- Measure: prefer the simplest correct approach; atomics help only in specific hotspots

---

## 9) Performance notes

- Atomics avoid kernel-level blocking but still serialize on cache lines
- Contention can be worse than a lock if many goroutines hammer a single variable
- Shard counters to reduce contention; combine periodically

### Design exercise: Lock-free counter service (sharding + aggregation)

Prompt (interview-ready)
- Design a high-throughput counter service supporting Increment(key) and Read(key|total) under heavy concurrency on multi-core.
- Minimize lock contention; discuss correctness, consistency, and scalability. Compare with a mutex-based approach.

Expected approach
- Shard counters per key (or per service) across N shards to reduce contention.
  - Choose shard = hash(key) % N (e.g., FNV/xxhash). Each shard uses atomic.Int64.
  - Increment(key): atomic add on shard counter for that key (or per-shard map guarded by independent locks if key space is large; the lock is now hot only per-shard).
- Reads
  - Total: sum across shards (iterate N atomics). For very hot read paths, maintain a periodic aggregated snapshot in atomic.Value.
  - Per-key: either (a) same sharding with per-shard key buckets (adds small mutex per bucket), or (b) maintain per-key sharded atomics in a fixed array if key space is bounded.
- Aggregation snapshot (optional)
  - A goroutine ticks every T (e.g., 10–100ms), loads shard values, computes a total map or scalar, and atomically publishes via atomic.Value.
  - Readers can choose snapshot (fast, slightly stale) vs on-demand sum (fresh, slower).

Correctness and memory model
- atomic.Add/Load on each shard ensures word-level atomicity and establishes happens-before between writers and readers.
- Publishing snapshots with atomic.Value provides whole-value visibility (readers never see partial aggregates).

Trade-offs vs locks
- Atomics: avoids contended global mutex; scales with CPU cores; simple increments; total read is O(N_shards).
- RWMutex design: simpler to reason about per-key maps; may outperform atomics when keys are sparse and read sections are short; measure.

Hot keys and skew
- If a single key is hot, consider further sharding that key’s counter (per-key sharded array) and sum on read.
- Adaptive sharding: detect hot keys and allocate more shards dynamically (rebalancing strategy required).

Operational considerations
- Choose shard count as a power of two (masking) for speed; tune via benchmark.
- Snapshot interval balances freshness vs overhead; expose both fresh and snapshot reads.
- Persistence: periodically flush snapshot to storage; ensure flush is monotonic.
- Resizing shards requires a migration phase (double-publish two sets and flip pointer, or stop-the-world briefly).

Pitfalls to call out
- Mixing atomic and non-atomic access to the same counters (data race).
- Using atomics for complex multi-field state; prefer immutable snapshots + pointer/value swap.
- Starvation with busy-spinning CAS loops; prefer Add/Load or backoff.

One-liner summary
- “Use sharded atomic counters to spread contention, optionally publish periodic atomic.Value snapshots for fast reads; provide fresh reads by summing shards. Compare throughput vs RWMutex and pick based on contention profile.”

- Bench with realistic concurrency (go test -bench -cpu=N)

---

## 10) Advanced interview questions

Atomic add+check vs mutex block
- Show that Add() returns new value atomically; branch on it for thresholds (see 010_add_check_threshold.go)
- Discuss when this is sufficient vs when a lock is still required (multi-field invariants)

Are atomics faster than locks?
- Usually for single-word ops (no kernel parking, minimal bookkeeping)
- Under heavy contention, cache-line ping-pong can be costly; RWMutex may compete or win; measure

Sharded counters
- Describe sharding under hot contention and aggregation strategies; trade-offs vs single atomic


1) Compare atomics vs locks for a read-mostly configuration store; design both
2) Explain why mixing atomic and non-atomic reads/writes to the same variable is a data race
3) Show a CAS loop to implement min() and discuss ABA problems
4) Design a hot-path counter: single atomic vs sharded counters; when is RWMutex better?
5) How does atomic.Value provide whole-value visibility? What are its limitations?
6) When would you choose atomic.Pointer[T] over atomic.Value?

