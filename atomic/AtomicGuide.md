# Go Atomics: Advanced Developer Guide

## **Table of Contents**
1. [Atomic Fundamentals](#atomic-fundamentals)
2. [Memory Model and Happens-Before](#memory-model-and-happens-before)
3. [The Two APIs: Function-Style vs Typed](#the-two-apis-function-style-vs-typed)
4. [Common Patterns](#common-patterns)
5. [Read-Mostly Data with atomic.Value](#read-mostly-data-with-atomicvalue)
6. [Atomics vs Locks: When to Choose Which](#atomics-vs-locks-when-to-choose-which)
7. [Common Mistakes and Gotchas](#common-mistakes-and-gotchas)
8. [Best Practices](#best-practices)
9. [Performance Considerations](#performance-considerations)
10. [Advanced Challenge Questions](#advanced-challenge-questions)
11. [Benchmarking](#benchmarking)
12. [Atomic Cheat Sheet (Go)](#atomic-cheat-sheet-go)

**Run these examples:**
- Atomic counter: `go run atomic/001_counter.go`
- Atomic flag/shutdown: `go run atomic/002_flag.go`
- CAS loop (min/once-style): `go run atomic/003_cas_spin.go`
- Read-mostly config with atomic.Value: `go run atomic/004_value_config.go`
- Pointer swap with atomic.Pointer[*T]: `go run atomic/005_pointer_swap.go`
- Sharded counter: `go run atomic/006_sharded_counter.go`
- Versioned pointer (ABA mitigation): `go run atomic/007_aba_versioned_pointer.go`
- Memory ordering (publish/subscribe): `go run atomic/008_memory_ordering.go`
- Pointer vs Value trade-offs: `go run atomic/009_pointer_vs_value_example.go`
- Atomic threshold check: `go run atomic/010_add_check_threshold.go`
- Periodic reset via Swap (epochs): `go run atomic/011_periodic_reset.go`
- Benchmarks: `go test -bench=. -benchmem ./atomic/bench`

---

## Quick Patterns Index
- Counters: [1) Counters](#1-counters)
- Threshold check: [2) Threshold Check After Increment](#2-threshold-check-after-increment)
- Flags: [3) Shutdown Flags](#3-shutdown-flags)
- CAS loops: [4) CAS Loops (Compare-and-Swap)](#4-cas-loops-compare-and-swap)
- Read-mostly snapshots: [atomic.Value](#read-mostly-data-with-atomicvalue)
- Pointer vs Value trade-offs: [atomic.Pointer[T] vs atomic.Value](#atomicpointert-vs-atomicvalue)
- Memory ordering: [Memory Model and Happens-Before](#memory-model-and-happens-before)
- Benchmarks: [Benchmarking](#benchmarking)


## Atomic Fundamentals

### **What are "atomics"?**

Operations that happen indivisibly—no other goroutine can see a half-done update. Use them for tiny shared pieces of state (counters, flags) without a mutex.

Atomics provide lock-free operations on single words (int64, pointer, etc.) that are:
- **Indivisible**: No other goroutine can observe a partial update
- **Ordered**: Establish happens-before relationships in the memory model
- **Race-free**: Safe for concurrent access without additional synchronization

Go gives you:
- **Typed atomics** (Go 1.19+): `atomic.Int64`, `atomic.Uint32`, etc.
- **Functions**: `atomic.AddInt64`, `atomic.LoadInt64`, `atomic.CompareAndSwapInt64`, etc.
- **atomic.Value**: store/load any type safely (as a whole)


> Why lock-free matters
> - Atomics avoid lock bookkeeping and goroutine parking on the fast path
> - Reads can be a single atomic load; writes can be a single atomic store or CAS
> - Under read-heavy workloads, this dramatically reduces contention and latency
> - See [Benchmarking](#benchmarking) for measured results and interpretation


> Terminology note
> - Lock-free: system as a whole makes progress without requiring any specific goroutine to run; individual goroutines may starve
> - Wait-free: every goroutine completes its operation in a bounded number of steps (stronger guarantee, rarer in practice)
> - Go’s sync/atomic provides lock-free primitives; most examples here are lock-free but not strictly wait-free

**Rule**: Don't mix plain reads/writes with atomic ops on the same variable. Use Load/Store.

### **1) Atomic counter (increment + read)**

```go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    var count atomic.Int64 // typed atomic (Go 1.19+)
    var wg sync.WaitGroup

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            count.Add(1) // atomic increment
        }()
    }
    wg.Wait()
    fmt.Println("count =", count.Load()) // safe read
}
```

**Why**: Classic case—many goroutines bump a shared number.

### **2) Atomic flag (turn "on" once)**

```go
package main

import (
    "fmt"
    "sync/atomic"
)

func main() {
    var ready atomic.Uint32

    // set flag
    ready.Store(1) // 1 = true

    // read flag
    if ready.Load() == 1 {
        fmt.Println("ready!")
    }
}
```

**Why**: Replace bool with Uint32 for a lock-free on/off flag.

### **3) Threshold check (add then test)**

```go
package main

import (
    "fmt"
    "sync/atomic"
)

func main() {
    var count int64
    newVal := atomic.AddInt64(&count, 1) // add returns the new value
    if newVal >= 5 {
        fmt.Println("limit reached:", newVal)
    }
}
```

**Why**: Safe "increment then check" without a mutex.

### **4) Compare-And-Swap (CAS) — update only if value matches**

```go
package main

import (
    "fmt"
    "sync/atomic"
)

func main() {
    var x int64

    // set to 10 only if it's currently 0
    swapped := atomic.CompareAndSwapInt64(&x, 0, 10)
    fmt.Println("swapped?", swapped, "x =", x)

    // retry loop (typical CAS pattern)
    for {
        cur := atomic.LoadInt64(&x)
        if cur >= 100 {
            break
        }
        if atomic.CompareAndSwapInt64(&x, cur, cur+1) {
            break
        }
    }
    fmt.Println("final x =", atomic.LoadInt64(&x))
}
```

**Why**: Conditional updates without locks. If CAS fails, reload and retry.

### **What atomics are NOT**

- **Not general mutual exclusion**: Only protect single variables
- **Not magic**: Still need careful design for complex state
- **Not always faster**: Under heavy contention, locks may perform better

---

## Memory Model and Happens-Before

### **Why use Load/Store instead of plain reads/writes?**

If a variable is written with atomic ops, all concurrent reads must also be atomic:
- **Data races**: Mixing plain reads causes undefined behavior
- **Torn reads**: On 32-bit arch, 64-bit values may tear without atomic loads
- **Memory ordering**: Atomics establish happens-before; plain reads may see stale values

### **Race Example: Bad vs Good**

**❌ Bad (data race):**
```go
var count int64

// Writer goroutine
atomic.AddInt64(&count, 1)

// Reader goroutine
fmt.Println(count) // RACE: plain read while others write atomically
```

**✅ Good (atomic reads):**
```go
var count atomic.Int64

// Writer goroutine
count.Add(1)

// Reader goroutine
fmt.Println(count.Load()) // Safe: atomic read
```

**Run the examples:**
- Bad: `go run -race atomic/012_bad_plain_read.go`
- Good: `go run -race atomic/013_good_atomic_load.go`

### **32-bit/Alignment Note**

On 32-bit platforms, 64-bit values require proper alignment. Prefer typed atomics (`atomic.Int64`) or ensure 64-bit fields are naturally aligned (e.g., placed first in structs) to avoid torn reads.

---

## The Two APIs: Function-Style vs Typed

### **Function-Style (Legacy)**

```go
var count int64

atomic.AddInt64(&count, 1)
v := atomic.LoadInt64(&count)
atomic.StoreInt64(&count, 42)
```

### **Typed Atomics (Go 1.19+, Recommended)**

```go
var count atomic.Int64

count.Add(1)
v := count.Load()
count.Store(42)
```

**Benefits of typed atomics:**
- Type safety and clarity
- Automatic alignment on all platforms
- Method-based API is more readable
- Cannot accidentally mix atomic/non-atomic access

---

## Common Patterns

### **1. Counters**

```go
var requests atomic.Int64

func handleRequest() {
    requests.Add(1)
    // ... handle request
}

func getStats() int64 {
    return requests.Load()
}
```

### **2. Threshold Check After Increment**

```go
func incAndCheck(limit int64) error {
    v := count.Add(1) // atomic increment; returns new value
    if v >= limit {
        return fmt.Errorf("too many (count=%d)", v)
    }
    return nil
}
```

**Run:** `go run atomic/010_add_check_threshold.go`

### **3. Shutdown Flags**

```go
var shutdown atomic.Bool

func worker() {
    for !shutdown.Load() {
        // do work
        time.Sleep(100 * time.Millisecond)
    }
}

func stop() {
    shutdown.Store(true)
}
```

**Run:** `go run atomic/002_flag.go`

### **4. CAS Loops (Compare-and-Swap)**

```go
func atomicMin(addr *atomic.Int64, val int64) {
    for {
        old := addr.Load()
        if val >= old {
            return // no update needed
        }
        if addr.CompareAndSwap(old, val) {
            return // successfully updated
        }
        // retry if CAS failed
    }
}
```

**Run:** `go run atomic/003_cas_spin.go`

---

## Read-Mostly Data with atomic.Value

### **What is atomic.Value?**

`atomic.Value` is a special type that lets you store and load any Go value atomically (indivisibly). Unlike other atomic functions that only work on specific types (int32, int64, etc.), `atomic.Value` can hold any type.

**Think of it as:** "A box where you can replace the whole object in one atomic operation."

### **Why do we need atomic.Value?**

**Without atomic.Value (using mutex):**
```go
var mu sync.RWMutex
var cfg *Config

// Readers
mu.RLock()
c := cfg
mu.RUnlock()

// Writers
mu.Lock()
cfg = newCfg
mu.Unlock()
```

**With atomic.Value (lock-free):**
```go
var cfg atomic.Value
cfg.Store(&Config{Rate: 10})

// Readers (no locks needed)
c := cfg.Load().(*Config)  // safe, always consistent

// Writers
cfg.Store(&Config{Rate: 20})  // atomic swap
```

### **Key Guarantees**

- `Store()` replaces the value atomically
- `Load()` always returns a complete value (never partially written)
- Multiple goroutines can safely call `Load()` and `Store()` simultaneously
- All stored values must be of the same concrete type (first stored value decides)

### **Where does Store() actually store the value?**

**Memory allocation:**
- The `&Config{...}` struct is allocated on the heap (normal Go allocation)
- `atomic.Value` doesn't own separate storage; it atomically publishes a reference to your object
- Inside, `atomic.Value` holds an interface pair: (dynamic type pointer, data pointer)
- `Store()` swaps that pair atomically

**Mental model:**
```
Before:
config (atomic.Value) -> (type=*Config, ptr=0xA1B2)  // points to old config on heap

Store(&Config{...})  // allocate new *Config at 0xC3D4, atomically swap pointer

After:
config (atomic.Value) -> (type=*Config, ptr=0xC3D4)  // now points to new config
```

**Readers:**
```go
c := config.Load().(*Config)  // returns the same pointer (0xA1B2 or 0xC3D4)
```

### **Real-World Patterns**

**1. Configuration Hot Reload**
```go
type Config struct {
    Timeout time.Duration
    MaxConns int
}

var config atomic.Value

func GetConfig() *Config {
    return config.Load().(*Config)
}

func UpdateConfig(newCfg *Config) {
    config.Store(newCfg)  // atomic swap
}

// Usage
func handleRequest() {
    cfg := GetConfig()  // fast read, no locks
    // use cfg.Timeout, cfg.MaxConns...
}
```

**2. Fast-Path Cache**
```go
var cache atomic.Value

func GetCache() map[string]string {
    return cache.Load().(map[string]string)
}

func RebuildCache() {
    newCache := make(map[string]string)
    // ... populate newCache
    cache.Store(newCache)  // atomic replacement
}
```

**3. Copy-on-Write Pattern**
```go
// Writer path (immutable updates)
old := config.Load().(*Config)
next := *old              // make a copy (value copy)
next.Timeout = 3 * time.Second
config.Store(&next)       // publish new snapshot


### FAQ: What does cfg.Store(&Config{...}) actually do? Which snapshot do readers see?

What `cfg.Store(&Config{...})` does
- `&Config{Rate: 20, Name: "beta"}` creates a new Config instance (a new object) and returns its pointer
- `cfg.Store(ptr)` does not clone or allocate; it just publishes exactly that pointer atomically
- Each `Store(&Config{...})` publishes a new snapshot (a different `*Config`) to all readers

What `Load` returns
- `cfg.Load()` returns whichever snapshot was current at the instant of the `Load` — the exact pointer most recently stored at that moment
- With concurrent writers, two back-to-back `Store`s can race; different goroutines may observe different snapshots depending on timing. Each `Load` is atomic and consistent (no partial writes)

“How do I know which store I got?”
- `atomic.Value` doesn’t tell you which writer/store produced the snapshot. If you need to identify it:
  - Add a version/timestamp field to the struct and check it:
    ```go
    type Config struct {
        Version int64
        Rate    int
        Name    string
    }
    // writer:
    cfg.Store(&Config{Version: v+1, Rate: 20, Name: "beta"})
    // reader:
    c := cfg.Load().(*Config)
    fmt.Println(c.Version, c.Rate, c.Name)
    ```
  - Compare pointers if you’ve saved the previous pointer:
    ```go
    p := cfg.Load().(*Config)
    if p == lastSeenPtr { /* same snapshot */ }
    ```

Important best practices
- Store immutable snapshots. Don’t mutate a `*Config` after you’ve stored it. Create a new one and `Store` that
- Keep type consistent. The first `Store` sets the allowed concrete type; subsequent stores must use the same type (here `*Config`)
- Single writer is simplest. With one writer goroutine, readers will eventually see the latest snapshot; at any instant, a reader sees either the old or the new, never a mix

By-value vs by-pointer
- You can also store by value (not pointer): `cfg.Store(Config{...})`. That copies the struct and avoids accidental mutation through shared pointers. For small structs, this is often cleaner
- If the struct is large, pointer snapshots are fine — just keep them immutable

Bottom line
- Each `Store(&Config{...})` creates and publishes a new instance. A `Load()` gets whichever complete snapshot was current at that moment; if you must identify it, include a `Version` (or timestamp) field or compare pointers

// Readers always see either old or new—never half-updated
```

### **Mutex vs atomic.Value Comparison**

**With sync.RWMutex:**
```
Readers:
┌─────────────────────────────┐
│ mu.RLock()                  │
│   c := cfg                  │  // copy pointer under RLock
│ mu.RUnlock()                │
└─────────────────────────────┘

Writers:
┌─────────────────────────────┐
│ mu.Lock()                   │
│   cfg = newCfg              │  // publish update
│ mu.Unlock()                 │
└─────────────────────────────┘
```

**With atomic.Value:**
```
Readers (no locks):
┌─────────────────────────────┐
│ c := config.Load().(*Config)│  // atomic load of pointer
└─────────────────────────────┘

Writers:
┌─────────────────────────────┐
│ config.Store(newCfg)        │  // atomic swap
└─────────────────────────────┘
```

### **When to use atomic.Value**

**Choose atomic.Value when:**
- Read-mostly workload (configs, lookup tables, feature flags)
- You can replace the whole object in one go (immutable snapshot style)
- All stored values share the same concrete type
- Many readers, few writers

**Choose mutex when:**
- Need to mutate parts of a shared structure in place
- Multiple fields must be updated with invariants across them
- Read-modify-write sequences span more than one atomic swap

### **5) atomic.Value — swap whole objects safely**

```go
package main

import (
    "fmt"
    "sync/atomic"
)

type Config struct {
    Rate int
    Name string
}

func main() {
    var cfg atomic.Value

    // initial config
    cfg.Store(&Config{Rate: 10, Name: "alpha"})

    // reader goroutine
    load := func() *Config {
        return cfg.Load().(*Config) // always atomic + consistent
    }

    // writer swaps config atomically
    cfg.Store(&Config{Rate: 20, Name: "beta"})

    c := load()
    fmt.Println(c.Rate, c.Name) // 20 beta
}
```

**Why**: Multiple fields updated together as one atomic pointer swap (readers never see a half-updated struct).

### **Common Gotchas (Quick Hits)**

- Always use `Load/Store` for the same variable you update atomically—no plain `x++` or `fmt.Println(x)` reads
- On 32-bit systems, `int64` must be properly aligned; typed atomics (e.g., `atomic.Int64`) handle this safely
- Atomics are great for simple state (counters/flags). For multi-variable invariants or longer critical sections → use a mutex
- Heavy contention on a single atomic can cause cache ping-pong; for super-hot paths consider sharded counters

### **atomic.Pointer[T] vs atomic.Value**

**Use atomic.Pointer[T] when:**
- Working with pointers to specific types
- Want type safety without type assertions
- Go 1.19+ available

**Use atomic.Value when:**
- Need to store different types in the same variable
- Working with value types (structs, not pointers)
- Backward compatibility with older Go versions


#### Store vs Swap on atomic.Pointer[T]

- Store(newPtr): updates the pointer; does not return the previous value
- Swap(newPtr): atomically replaces the pointer and returns the previous pointer

Why Swap exists (and when to prefer it):
- Atomicity across “get old + set new”: a single indivisible operation, safe under heavy concurrency
- Need the old value: to log, free resources, or compute deltas/rollbacks

Example:
```go
var cfg atomic.Pointer[Config]
// store initial
cfg.Store(&Config{Rate: 10, Name: "alpha"})

// swap and capture old
old := cfg.Swap(&Config{Rate: 20, Name: "beta"})
fmt.Println("old:", old.Rate, old.Name)                // 10 alpha
fmt.Println("new:", cfg.Load().Rate, cfg.Load().Name) // 20 beta
```

Store vs Swap under races:
- With Store, you cannot reliably retrieve the previous pointer without a race (Load then Store can be interleaved by another writer)
- With Swap, you atomically update and fetch the old pointer in one step

**Run these examples:**
- Configuration snapshots: `go run atomic/004_value_config.go`
- Pointer vs Value trade-offs: `go run atomic/009_pointer_vs_value_example.go`

---

## Atomics vs Locks: When to Choose Which

### **Use atomics when:**
- Simple counters, gauges, flags; one variable at a time
- Fast paths avoiding lock contention/overhead
- Copy-on-write snapshots so readers only Load a pointer/value

### **Use locks when:**
- You must update multiple fields together or enforce invariants across values
- The critical section spans more than a single atomic op or includes I/O/complex logic
- You need mutual exclusion across a sequence of operations

### **Performance Comparison**

Under very high contention, a single atomic can still bottleneck (cache-line ping-pong). Consider:
- Sharded counters: `go run atomic/006_sharded_counter.go`
- RWMutex for complex read-heavy structures


## Benchmarking

Here’s a micro-benchmark you can run with `go test -bench . -benchmem` to compare read-heavy config access using `sync.RWMutex` vs `atomic.Value`.

It measures:
- Reads (many readers, no writes in the hot path)
- Writes (occasional config swaps)

Expect `atomic.Value` to shine on reads (one atomic load) and be comparable on writes (single atomic swap).

Run:
- `go test -bench=. -benchmem ./atomic/bench`

Code: `atomic/bench/bench_atomic_vs_mutex_test.go`

### How to read the results

A typical line looks like:

```
BenchmarkRead_Atomic-8        20000000   65.0 ns/op    0 B/op   0 allocs/op
```

- `-8`: number of OS threads (GOMAXPROCS) used
- `ns/op`: average nanoseconds per operation (lower is better)
- `B/op`: bytes allocated per operation (lower is better)
- `allocs/op`: number of allocations per operation (lower is better)

### What to expect (rules of thumb)

Reads (many readers, no/rare writes)
- `BenchmarkRead_Atomic` should be faster (lower ns/op) than `BenchmarkRead_Mutex`
- Both should show `0 B/op`, `0 allocs/op` (we’re just reading pointers)
- Typical ranges (vary by machine):
  - Read_Atomic: ~10–80 ns/op
  - Read_Mutex: ~30–200 ns/op

Why: `atomic.Value.Load()` is a single atomic load; `RLock/RUnlock` adds overhead and contention.

Writes (occasional swaps)
- `BenchmarkWrite_Atomic` and `BenchmarkWrite_Mutex` are often close
- Both publish a new pointer; mutex adds lock/unlock, atomic adds a single atomic store
- Expect similar allocations: `0 B/op`, `0 allocs/op`

Mixed (many reads, few writes)
- Atomic version usually wins overall throughput; advantage grows with read bias

### Interview-ready interpretation

Core conclusion:
- For read-mostly access patterns (configs, feature flags, lookup tables), `atomic.Value` gives lock-free reads (one atomic load), significantly reducing overhead and contention compared to `RWMutex`.

Why atomic reads are faster:
- An RLock still touches shared state (mutex bookkeeping) and may contend
- An atomic load is essentially a single CPU instruction with memory ordering constraints

Why writes are similar:
- Both approaches ultimately publish a new pointer:
  - Mutex: lock, write pointer, unlock
  - Atomic: single atomic pointer swap

When not to use `atomic.Value`:
- If you must mutate fields in place or maintain invariants across multiple variables, use a mutex. `atomic.Value` is best for replace-whole-object (copy-on-write) patterns.

### Fairness and tips

- CPU & OS matter: results vary by CPU architecture, caches, and scheduler
- Parallelism: try different GOMAXPROCS
  - `GOMAXPROCS=1 go test -bench . -benchmem`
  - `GOMAXPROCS=8 go test -bench . -benchmem`
- Stability: increase benchtime
  - `go test -bench . -benchmem -benchtime=3s`
- Avoid allocations in hot loops (the provided code already does)
- If adding counters, pad structs to avoid false sharing

Sound bite:
- “On my machine, atomic reads were ~2–4× faster than RLock/RUnlock reads with zero allocations, and writes were comparable. That’s why for read-heavy config access, `atomic.Value` is ideal: readers do a single atomic load, and writers publish a new snapshot with one atomic store. If I need to update multiple fields with invariants or mutate in place, I’d switch to a mutex.”

### Sample results (Apple M1 Pro)

Command: `go test -bench=. -benchmem -benchtime=2s ./atomic/bench`

- Read-heavy (config access)
  - Read_Mutex: 118.3 ns/op, 0 B/op, 0 allocs/op
  - Read_Atomic: 0.2232 ns/op, 0 B/op, 0 allocs/op
- Writes (publish new config)
  - Write_Mutex: 34.48 ns/op, 16 B/op, 1 allocs/op
  - Write_Atomic: 16.78 ns/op, 16 B/op, 1 allocs/op
- Mixed workload (many reads, few writes)
  - Mixed_Mutex: 124.3 ns/op, 0 B/op, 0 allocs/op
  - Mixed_Atomic: 57.94 ns/op, 0 B/op, 0 allocs/op
- Read-mostly store
  - RWMutexReadMostly: 73.56 ns/op, 0 B/op, 0 allocs/op
  - AtomicValueReadMostly: 0.3915 ns/op, 0 B/op, 0 allocs/op
- Resets (epoch swap vs mutex)
  - AtomicEpochReset: 58.27 ns/op, 0 B/op, 0 allocs/op
  - MutexReset: 124.9 ns/op, 0 B/op, 0 allocs/op
- Counter increment
  - MutexInc: 114.2 ns/op, 0 B/op, 0 allocs/op
  - AtomicInc: 58.27 ns/op, 0 B/op, 0 allocs/op

Interpretation: Atomic reads were ~2–4× faster than RLock/RUnlock in typical cases, and up to orders of magnitude faster in the microbench here. Writes were comparable, with atomic slightly faster. Mixed workloads favored atomic as read bias increased.

**Benchmarks:**
- `go test -bench=. -benchmem ./atomic/bench`
- `go test -bench=ReadMostly -benchmem ./atomic/bench` (atomic.Value vs RWMutex)

---

## Common Mistakes and Gotchas

### **Mistakes Index (quick scan)**
- Mixing atomic and non-atomic access to same variable
- Piecemeal updates to multi-field structs with atomics (use locks or snapshots)
- Busy CAS loops without backoff (burn CPU)
- Expecting atomic.Value to merge fields (it swaps whole snapshots)
- Copying typed atomics after use (undefined behavior)
- Using atomics where a lock is clearer/faster under contention

### **Detailed Examples**

**1. Mixing atomic and non-atomic access**
```go
// BAD
var count int64
atomic.AddInt64(&count, 1) // atomic write
fmt.Println(count)         // plain read - RACE!

// GOOD
var count atomic.Int64
count.Add(1)               // atomic write
fmt.Println(count.Load())  // atomic read
```

**2. Resetting counters safely**
```go
// BAD
count = 0  // plain write while others do atomic ops

// GOOD - epoch swap
old := count.Swap(0)  // atomic reset, returns previous value
```

**Run:** `go run atomic/011_periodic_reset.go`

---

## Best Practices

- **Prefer typed atomics** (`atomic.Int64/Bool/Pointer[T]`) for clarity and type safety
- **On 32-bit platforms**, ensure 64-bit fields are aligned; prefer `atomic.Int64` to avoid torn reads
- **Use immutable snapshots** with `atomic.Value` / `atomic.Pointer` to avoid reader locks
- **Encapsulate atomics** in your types; expose methods, not raw fields
- **Document** whether fields are updated atomically or under a lock
- **Measure**: prefer the simplest correct approach; atomics help only in specific hotspots

---

## Performance Considerations

### **When atomics shine:**
- High-frequency counters/flags with many readers
- Lock-free fast paths in hot code
- Read-mostly data with occasional updates

### **When to reconsider:**
- Complex multi-field updates (prefer locks)
- Heavy contention on single atomic (consider sharding)
- Simple, infrequent operations (locks may be clearer)

### **Sharding Hot Counters**

```go
type ShardedCounter struct {
    shards []atomic.Int64
}

func (s *ShardedCounter) Inc(key string) {
    h := hash(key) % len(s.shards)
    s.shards[h].Add(1)
}

func (s *ShardedCounter) Total() int64 {
    var total int64
    for i := range s.shards {
        total += s.shards[i].Load()
    }
    return total
}
```

**Run:** `go run atomic/006_sharded_counter.go`

---

## Advanced Challenge Questions

### **Atomic add+check vs mutex block**
- Show that `Add()` returns new value atomically; branch on it for thresholds
- Discuss when this is sufficient vs when a lock is still required (multi-field invariants)

### **Are atomics faster than locks?**
- Usually for single-word ops (no kernel parking, minimal bookkeeping)
- Under heavy contention, cache-line ping-pong can be costly; RWMutex may compete or win; measure

### **FAQ: Designing a lock-free-ish counter service**
- Shard counters per key across N shards to reduce contention
- Choose shard = hash(key) % N; each shard uses `atomic.Int64`
- Reads: sum across shards or maintain periodic aggregated snapshot in `atomic.Value`
- Consider trade-offs vs RWMutex and pick based on contention profile

---

## Atomic Cheat Sheet (Go)

### 0) Golden rule / mnemonic

- Using plain types (int64, uint32, etc.) → call package functions and pass &var
- Using atomic wrapper types (atomic.Int64, atomic.Uint32, atomic.Pointer[T]) → call methods on the value (Go auto-addresses)

Think: “plain needs address; atomic type just works.”

### 1) Load / Store

Operation | Old API (plain var) | New API (atomic type)
---|---|---
Load | `v := atomic.LoadInt64(&x)` | `v := at.Load()`
Store | `atomic.StoreInt64(&x, 42)` | `at.Store(42)`

Types:
- Old: `LoadInt32/64`, `StoreInt32/64`, `LoadUint32/64`, `StoreUint32/64`, `LoadPointer`, `StorePointer`
- New: `atomic.Int32/Int64/Uint32/Uint64/Bool`, `atomic.Pointer[T]`

### 2) Add / Swap / CompareAndSwap (CAS)

Operation | Old API (plain var) | New API (atomic type)
---|---|---
Add | `nv := atomic.AddInt64(&x, 1)` | `nv := at.Add(1)`
Swap | `ov := atomic.SwapInt64(&x, 7)` | `ov := at.Swap(7)`
CAS | `ok := atomic.CompareAndSwapInt64(&x, old, new)` | `ok := at.CompareAndSwap(old, new)`

### 3) Pointers

Operation | Old API (unsafe.Pointer) | New API (generic)
---|---|---
Load | `p := (*T)(atomic.LoadPointer(&ptr))` | `p := ap.Load()` (where `ap := atomic.Pointer[T]{}`)
Store | `atomic.StorePointer(&ptr, unsafe.Pointer(p))` | `ap.Store(p)`
Swap | `op := (*T)(atomic.SwapPointer(&ptr, unsafe.Pointer(np)))` | `op := ap.Swap(np)`
CAS | `ok := atomic.CompareAndSwapPointer(&ptr, oldP, newP)` | `ok := ap.CompareAndSwap(oldP, newP)`

Prefer `atomic.Pointer[T]` in new code; it’s typesafe and nicer.

### 4) Initialization & zero values

- Old API: just a plain zeroed var (`var x int64`)
- New API: zero value is ready to use (`var at atomic.Int64`), or embed as a field

### 5) When to use which

- Prefer the new type-based API in modern code (Go 1.19+): clearer and safer
- Use the old API if:
  - You’re working with legacy code bases
  - You must operate on plain variables (not feasible to switch types)
  - You need operations not wrapped (rare nowadays)

### 6) Gotchas (important!)

- Don’t copy a value that contains atomic fields after first use (same rule as sync.Mutex). Store them in structs and pass pointers to those structs instead of copying
- Methods on `atomic.*` types have pointer receivers; calling `v.Load()` works because addressable variables are auto-addressed. Avoid storing atomic values in interfaces or copying structs around
- Keep 64-bit atomics sensibly placed in structs to avoid alignment surprises or false sharing (e.g., group frequently-updated atomics, consider padding if contended)

### 7) Quick reference (by operation)

- Load
  - Old: `atomic.Load<Type>(&x)`
  - New: `a.Load()`
- Store
  - Old: `atomic.Store<Type>(&x, v)`
  - New: `a.Store(v)`
- Add
  - Old: `atomic.Add<Type>(&x, d)`
  - New: `a.Add(d)`
- Swap
  - Old: `atomic.Swap<Type>(&x, nv)`
  - New: `a.Swap(nv)`
- CAS
  - Old: `atomic.CompareAndSwap<Type>(&x, old, new)`
  - New: `a.CompareAndSwap(old, new)`

## FAQ
- Are atomics always faster than locks?
  - Often for single-word ops and read-mostly snapshots, yes. See [Benchmarking](#benchmarking). Under heavy contention or multi-field invariants, measure and consider locks.
- When should I use atomic.Pointer[T] vs atomic.Value?
  - Pointer: type-safe pointer swaps without assertions; store only pointers to a single concrete type.
  - Value: can hold any value type but all Stores must share the same concrete type; great for copy-on-write snapshots. See [Pointer vs Value](#atomicpointert-vs-atomicvalue).
- Can I mix atomic writes with plain reads?
  - No. Use atomic Load/Store on the same variable. See [Memory Model](#memory-model-and-happens-before) and the bad/good runnable examples.
- How do I reset a counter safely while others increment?
  - Use `Swap(0)` (epoch rotation). See `atomic/011_periodic_reset.go` and [Mistakes](#common-mistakes-and-gotchas).
- Lock-free vs wait-free?
  - Lock-free: system makes progress without requiring any particular goroutine to run. Wait-free: each goroutine completes in bounded steps. Most patterns here are lock-free, not strictly wait-free.


## Further Reading
- Go Memory Model: https://go.dev/ref/mem
- Package sync/atomic docs: https://pkg.go.dev/sync/atomic
- Blog: Copy-on-write with atomic.Value (various articles)
  - Example overview: replacing whole snapshots for read-mostly workloads


*For more examples and benchmarks, explore the `atomic/` directory and run the provided test suites.*
