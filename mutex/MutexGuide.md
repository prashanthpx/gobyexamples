# sync.Mutex and sync.RWMutex: Advanced Guide with Runnable Examples

Run these examples
- Basic lock/unlock: go run mutex/001_basic_mutex.go
- RWMutex: go run mutex/002_rwmutex.go
- Lock ordering (deadlock vs fix): go run mutex/003_deadlock_lock_order.go
- Copying a struct with a mutex (pitfall): go run mutex/004_copy_mutex.go
- sync.Cond producer/consumer: go run mutex/006_cond.go
- Goroutine-safe map: go run mutex/007_map_with_mutex.go
- Embedded mutex API: go run mutex/009_embedded_mutex.go
- Benchmarks: go test -bench=. -benchmem ./mutex/bench

---

## Table of Contents
1. [Why Mutexes (vs channels/atomic)](#toc-1-why-mutexes)
2. [sync.Mutex: semantics and patterns](#toc-2-mutex)
3. [sync.RWMutex: read-mostly optimization](#toc-3-rwmutex)
4. [Lock Ordering and Deadlocks](#toc-4-lock-ordering)
5. [Copying and Addressability Rules (do not copy a mutex)](#toc-5-copying)
6. [Unlocking: defer vs explicit and hot-path performance](#toc-6-unlocking)
7. [sync.Cond with a Mutex](#toc-7-cond)
8. [Maps and Mutexes (vs sync.Map)](#toc-8-maps)
9. [Embedding Mutexes in Types (API design)](#toc-9-embedding)
10. [Common Mistakes and Gotchas](#toc-10-mistakes)
11. [Best Practices](#toc-11-best)
12. [Advanced Interview Questions](#toc-12-advanced)
13. [Cheat-sheet (quick reminders)](#toc-13-cheatsheet)

---

<a id="toc-1-why-mutexes"></a>

## 1) Why Mutexes

- Mutexes protect shared, mutable state with minimal overhead; they are often the simplest correct tool
- Prefer channels for ownership transfer; prefer atomic for single-word counters/flags
- RWMutex helps when reads dominate writes and critical sections are short

---

<a id="toc-2-mutex"></a>

## 2) sync.Mutex: semantics and patterns

- Zero value is usable; a Mutex starts unlocked
- Lock blocks the caller until it acquires the lock; Unlock must be called from the goroutine that owns the lock and only when locked

Pattern: guard a counter
```go
// See: mutex/001_basic_mutex.go
```

Defer vs explicit Unlock
- Defer ensures Unlock in the presence of panics/early returns (safety)
- In extremely hot loops, avoid defer inside loop bodies for perf; unlock explicitly

---

<a id="toc-3-rwmutex"></a>

## 3) sync.RWMutex: read-mostly optimization

- RLock allows multiple readers concurrently; Lock excludes both readers and writers
- Use when reads are frequent and sections are short. Under heavy contention or long reads, RWMutex may hurt

```go
// See: mutex/002_rwmutex.go
```

---

<a id="toc-4-lock-ordering"></a>

## 4) Lock Ordering and Deadlocks

- Always acquire locks in a consistent global order across your codebase
- Never hold a lock while calling into unknown code (callbacks) unless documented
- Tip: the race detector (-race) won’t detect deadlocks; use it to find data races, and use timeouts/logging to surface potential deadlocks in tests

```go
// See: mutex/003_deadlock_lock_order.go
```

---

<a id="toc-5-copying"></a>

## 5) Copying and Addressability Rules

- Do not copy a value containing a sync.Mutex or sync.RWMutex once it’s in use
- Store structs with mutex by pointer if they will be passed around or put in maps/slices

```go
// See: mutex/004_copy_mutex.go
```

---

<a id="toc-6-unlocking"></a>

## 6) Unlocking: defer vs explicit and hot-path performance

- Safe default: defer mu.Unlock() immediately after mu.Lock()
- Hot-path: prefer explicit Unlock to avoid per-iteration defer cost; still ensure Unlock on errors

---

<a id="toc-7-cond"></a>

## 7) sync.Cond with a Mutex

- Coordinates goroutines waiting for a condition protected by a mutex
- Always check the predicate in a for loop to handle spurious wakeups

```go
// See: mutex/006_cond.go
```

---

<a id="toc-8-maps"></a>

## 8) Maps and Mutexes (vs sync.Map)

- For domain-specific access patterns, a map+mutex is often faster and clearer
- sync.Map is optimized for append-only or read-mostly with infrequent updates; benchmark before choosing

```go
// See: mutex/007_map_with_mutex.go
```

---

<a id="toc-9-embedding"></a>

## 9) Embedding Mutexes in Types (API design)

- Embed an unexported mutex to control access in methods
- Don’t export a struct field of type sync.Mutex; it invites misuse (copying)

```go
// See: mutex/009_embedded_mutex.go
```

---

<a id="toc-10-mistakes"></a>

## 10) Common Mistakes and Gotchas

1) Copying a struct that contains a Mutex after first use
- Can lead to panics or data races; never copy after locking even once

2) Locking the same mutex twice in the same goroutine (not re-entrant)
- Deadlocks; design to avoid re-entrancy or split critical sections

3) Unlocking a mutex you don’t hold
- Panics; pair Lock/Unlock carefully

4) Holding a lock across slow I/O/long operations
- Increases contention and latency; shorten critical sections

5) Using RWMutex when writes are frequent or read sections are long
- Can be slower than Mutex; measure first

6) Defer in hot loops
- Convenient but adds overhead; consider explicit Unlock

---

<a id="toc-11-best"></a>

## 11) Best Practices

- Keep critical sections small; do work outside the lock
- Document lock ordering and ownership
- Prefer composition: methods own locking, callers don’t touch the mutex directly
- Avoid exporting mutex fields; keep them unexported and non-copyable by convention
- Consider atomic for counters/flags; consider channels for ownership transfer

---

<a id="toc-12-advanced"></a>

## 12) Advanced Interview Questions

1) Explain when RWMutex outperforms Mutex and when it underperforms
- Reads must dominate, read sections short; otherwise contention hurts

2) Why is copying a struct containing a Mutex unsafe? How do you prevent it?
- Internal state becomes inconsistent; use pointers, avoid copying, hide mutex

3) How do you prevent deadlocks with multiple locks?
- Global lock order; acquire in increasing order; release in reverse

4) Can you upgrade a read lock to a write lock with RWMutex?
- No. You must release RLock and then Lock; design APIs to avoid upgrade

5) When would you avoid defer mu.Unlock()?
- In hot loops; use explicit Unlock but ensure all paths unlock

6) How would you design a threadsafe type with internal invariants?
- Embed an unexported mutex, expose methods that lock internally, avoid leaking references that break invariants

---

## 13) Cheat-sheet

- sync.Mutex (exported type): zero value is usable (unlocked). As a value, it’s a lock.
- *sync.Mutex: pointer to a mutex; share the same lock instance by reference (e.g., in a map or passed around).
- sync.Mutex{}: composite literal; equivalent to var mu sync.Mutex — both zero (unlocked) and ready.
- log = &sync.Mutex: doesn’t compile; you can’t take the address of a type name. Allocate as `mu := &sync.Mutex{}` or `mu := new(sync.Mutex)`. Also, avoid shadowing the log package; prefer a name like mu.

