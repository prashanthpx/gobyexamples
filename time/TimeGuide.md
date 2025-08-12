# Time in Go: Advanced Developer Guide

## Table of Contents
1. time.Time and Monotonic Clock
2. Timers vs Tickers vs time.After
3. Stopping and Resetting Timers/Tickers
4. Timezones, Locations, and Parsing
5. Deadlines and Context Integration
6. Common Mistakes and Gotchas
7. Best Practices
8. Performance Considerations
9. Advanced Challenge Questions

---

## 1) time.Time and Monotonic Clock

Go’s time.Time can carry a monotonic clock reading in addition to wall time. `time.Since(t)` and `t.Sub(u)` prefer monotonic components when present.

```go
start := time.Now()     // carries monotonic
work()
elapsed := time.Since(start) // uses monotonic; immune to wall clock changes
```

If you parse time from a string, it won’t have a monotonic component:
```go
u, _ := time.Parse(time.RFC3339, "2024-01-01T00:00:00Z")
_ = u.Sub(start) // wall clock arithmetic
```

---

## 2) Timers vs Tickers vs time.After

- time.After(d) creates a new timer each call and returns a channel that fires once
- time.NewTimer(d) returns a *Timer you can Stop/Reset
- time.NewTicker(d) returns a *Ticker that ticks repeatedly until Stop

```go
// One-shot delay
<-time.After(50*time.Millisecond)

// Repeated events
tick := time.NewTicker(100*time.Millisecond)
defer tick.Stop()
for i := 0; i < 3; i++ { <-tick.C }
```

Avoid time.After in loops; reuse Timer or Ticker instead to prevent allocations/leaks.

---

## 3) Stopping and Resetting Timers/Tickers

Always Stop timers/tickers you create.

```go
t := time.NewTimer(100*time.Millisecond)
if !t.Stop() {
  <-t.C // drain if already fired
}

// Reset for another interval
t.Reset(200*time.Millisecond)
<-t.C
```

```go
k := time.NewTicker(time.Second)
defer k.Stop() // prevent leak
```

---

## 4) Timezones, Locations, and Parsing

Parse vs ParseInLocation:
```go
layout := "2006-01-02 15:04"
// Input without zone: interpreted in UTC by Parse
u, _ := time.Parse(layout, "2024-11-05 09:30")
// Parsed in specific location
loc, _ := time.LoadLocation("America/New_York")
n, _ := time.ParseInLocation(layout, "2024-11-05 09:30", loc)
fmt.Println(u.Location(), n.Location()) // UTC America/New_York
```

DST edge cases:
```go
loc, _ := time.LoadLocation("America/New_York")
// Spring forward: some local times do not exist
missing := time.Date(2024, 3, 10, 2, 30, 0, 0, loc)
fmt.Println(missing) // normalized by time package
```


```go
loc, _ := time.LoadLocation("America/New_York")
t := time.Date(2024, 3, 10, 2, 30, 0, 0, loc) // DST boundary example
fmt.Println(t.String())

// Parsing with layouts (layout is an example date: Mon Jan 2 15:04:05 MST 2006)
when, _ := time.Parse("2006-01-02 15:04", "2024-11-05 09:30")
```

Notes:
- Layouts use the reference time “2006-01-02 15:04:05 -0700 MST”
- Use ParseInLocation when strings lack timezone info
- Beware DST transitions; some local times don’t exist or repeat

---

## 5) Deadlines and Context Integration

```go
ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
defer cancel()
select {
case <-ctx.Done():
  // deadline exceeded or cancelled
case <-time.After(50*time.Millisecond):
}
```

Prefer context deadlines for request-scoped operations; derive timeouts rather than hardcoding sleeps.

---

## 6) Common Mistakes and Gotchas

1) Using time.After in loops (allocates and can leak)
```go
for {
  select {
  case <-time.After(100*time.Millisecond):
    work()
  case <-ctx.Done():
    return
  }
}
// ✅ Use a single Ticker or a reusable Timer
```

2) Forgetting to Stop a Ticker/Timer
```go
k := time.NewTicker(time.Second)
// ❌ missing k.Stop(): goroutine leak
```

3) Misusing wall time for durations
```go
// ✅ Use time.Since for durations; immune to wall clock jumps
start := time.Now(); defer func(){ fmt.Println(time.Since(start)) }()
```

4) Parsing layouts incorrectly
```go
// ❌ Using strftime-like patterns
// ✅ Use Go’s reference layout 2006-01-02 15:04:05
```

---

## 7) Best Practices

- Prefer monotonic durations for measuring time (time.Since)
- Use context deadlines over ad-hoc timers when possible
- Always Stop timers and tickers; drain channels when Stop returns false
- Parse with explicit locations; document timezone assumptions

---

## 8) Performance Considerations

- Allocation: time.After allocates per call; prefer reusable Timer/Ticker in loops
- Avoid unnecessary time.Now calls in hot paths; batch or hoist
- Reuse buffers when formatting times with AppendFormat

---

## 9) Advanced Challenge Questions

1) Why does time.Since avoid issues with NTP jumps?
- It uses the monotonic component when available.

2) When should you use ParseInLocation vs Parse?
- When the input string lacks timezone info and you want a specific location applied.

3) How do you safely reset a timer that may have fired?
- Stop it and drain its channel if needed before Reset.

4) Why is time.After problematic inside select loops?
- It creates a new timer each iteration, increasing allocations and can keep goroutines alive if not read.

