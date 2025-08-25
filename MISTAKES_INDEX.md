# Go Mistakes Index — Quick Lookup and Fixes

Use this index to quickly find high-impact Go mistakes and jump to the exact section in this repo that shows the bad pattern, the fix, and the reasoning. All links point to short, compilable examples.

How to use:
- Scan the category and pick the mistake that matches your situation
- Click the link to go to the relevant guide section with fixes and notes
- Run examples with `go run` and validate with `go vet` and `go test -race`

---

## Concurrency and Goroutines

- time.After in loops (alloc/leak); use Ticker or reusable Timer
  - channels/ChannelsGuide.md → Common Mistakes: time.After leak and busy default
    - channels/ChannelsGuide.md#9-common-mistakes-and-gotchas
  - time/TimeGuide.md → Common Mistakes: time.After inside loops
    - time/TimeGuide.md#6-common-mistakes-and-gotchas
- Busy loop with select default; add backoff or proper wait path
  - channels/ChannelsGuide.md#9-common-mistakes-and-gotchas
- Goroutine leaks and cancellation; always honor ctx.Done
  - goroutine/GoroutinesGuide.md#8-goroutine-leaks-and-how-to-avoid-them
- One-shot select: worker processes only one job; loop/range channel
  - wait_group/ConcurrencyGuide.md#6-common-mistakes-and-gotchas
- Context propagation anti-patterns; don’t store ctx in structs
  - goroutine/GoroutinesGuide.md#8-goroutine-leaks-and-how-to-avoid-them

## Channels

- Closing from the receiver side (wrong); only sender should close
  - channels/ChannelsGuide.md#9-common-mistakes-and-gotchas
- Send on closed channel (panic) and reading from closed
  - channels/ChannelsGuide.md#9-common-mistakes-and-gotchas
- Directional channels for API intent (chan<- / <-chan)
  - channels/ChannelsGuide.md#4-channel-directions-chan--chan

## Slices and Arrays

- Hidden retention: small slice keeps huge backing array alive; use 3-index slice
  - slices/SliceGuide.md#memory-optimization
- Copy vs share: append overwriting shared backing array; force copy
  - slices/SliceGuide.md#memory-optimization
- Range variable capture pitfalls (addresses of loop vars)
  - functions/FunctionsGuide.md#common-mistakes-and-gotchas

## Maps

- Nil map write panic; initialize before writes
  - maps/MapsGuide.md#common-mistakes-and-gotchas
- Iteration order is not deterministic; sort keys
  - maps/MapsGuide.md#iteration-and-ordering
- Preallocate with make(map, n) to reduce rehash cost
  - maps/MapsGuide.md#best-practices
- Map value modification requires reassign or pointer
  - maps/MapsGuide.md#common-mistakes-and-gotchas

## Time

- Monotonic vs wall clock; prefer time.Since for durations
  - time/TimeGuide.md#1-time-in-go-and-monotonic-clock
- Stop/Reset timers and tickers correctly
  - time/TimeGuide.md#3-stopping-and-resetting-timerstickers
- Parse/ParseInLocation and layout pitfalls
  - time/TimeGuide.md#4-timezones-locations-and-parsing

## Errors

- Wrap with %w; check with errors.Is/As
  - testerrors/ErrorHandlingGuide.md#2-wrapping-with-w-and-unwrapping-with-errorsisas
- Typed vs sentinel errors and guidance
  - testerrors/ErrorHandlingGuide.md#3-sentinel-vs-typed-errors
- Typed-nil interface pitfall
  - testerrors/ErrorHandlingGuide.md#4-typed-nil-interfaces-and-pitfalls

## Strings, Runes, and Bytes

- Indexing bytes vs runes (Unicode); iterate safely
  - strings/StringsGuide.md#3-indexing-iteration-and-unicode
- Excessive conversions string↔[]byte allocate; cache representation
  - strings/StringsGuide.md#4-conversions-and-allocations
- Build strings efficiently with strings.Builder
  - strings/StringsGuide.md#5-building-strings-efficiently

## Functions, Defer, and Resources

- Defer in loops keeps resources open; wrap loop body in a function
  - functions/FunctionsGuide.md#common-mistakes-and-gotchas
- Not closing HTTP response bodies; close and drain for reuse
  - functions/FunctionsGuide.md#common-mistakes-and-gotchas

## Interfaces and Methods

- T vs *T method sets for interface satisfaction
  - interface/InterfacesGuide.md#3-method-sets-and-interface-satisfaction-t-vs-t
- Nil interface vs typed-nil in interfaces
  - interface/InterfacesGuide.md#7-nil-interfaces-vs-typed-nil-in-interfaces
- Method values vs method expressions (receiver binding)
  - methods/MethodsGuide.md#4-method-values-vs-method-expressions

## YAML / JSON / Tags

- omitempty hides zero values; validate semantics
  - structures/StructuresGuide.md#5-tags-encodingdecoding-jsonyaml
- YAML strict decoding to catch unknown fields
  - yaml/YAMLGuide.md#3-strict-decoding-and-unknown-fields

## Packages and Modules

- Semantic import versioning for v2+ (module path suffix)
  - packages/PackagesGuide.md#2-import-paths-and-semantic-import-versioning-v2
- internal/ to enforce encapsulation boundaries
  - packages/PackagesGuide.md#3-internal-packages-and-visibility
- Using := at package scope (expected declaration); use var for globals and := inside functions
  - arrays/ArraysGuide.md#package-scope-vs-function-scope-initialization--vs-var


## Atomics

- Mixing atomic and non-atomic access; data races
- Plain read of an atomically-updated variable (racy); use atomic.Load
  - Bad: atomic/012_bad_plain_read.go; Good: atomic/013_good_atomic_load.go
  - atomic/AtomicGuide.md#mistakes-guide-with-runnable-examples

  - atomic/AtomicGuide.md#7-common-mistakes-and-gotchas
- Piecemeal updates to multi-field structs with atomics
  - atomic/AtomicGuide.md#7-common-mistakes-and-gotchas
- Spinning CAS loops without backoff
  - atomic/AtomicGuide.md#7-common-mistakes-and-gotchas
- Copying typed atomics after first use
  - atomic/AtomicGuide.md#7-common-mistakes-and-gotchas
- Using atomics where RWMutex is clearer/faster under contention
  - atomic/AtomicGuide.md#6-atomics-vs-locks-when-to-use-which
- Not initializing atomic.Value before first Load (panic) or mutating stored snapshots
  - Initialize with Store once; store immutable snapshots thereafter
  - atomic/AtomicGuide.md#5-read-mostly-data
- ABA pitfalls with CAS on pointers; no versioning
  - Example: atomic/007_aba_versioned_pointer.go
- Failing to establish happens-before (publish/subscribe)
  - Example: atomic/008_memory_ordering.go
- Choosing the wrong abstraction (atomic.Value vs atomic.Pointer)
  - Example trade-offs: atomic/009_pointer_vs_value_example.go


- Resetting counters with non-atomic writes (racy); prefer epoch swap with atomic.Swap(0)
  - Example: atomic/011_periodic_reset.go


---

## Tooling tips

- go vet: catch fmt.Printf verb/arg mismatches and more
  - printf/PrintfGuide.md#vetting-format-strings
- Race detector: find racy code paths quickly
  - run with `go test -race` on your package(s)

---

## Contributing / Extending this Index

- Found a recurring mistake? Add a short, compilable example to the appropriate guide and link it here
- Keep items succinct: problem → link to fix → single-sentence rationale
- Aim for breadth and practicality; prefer examples you’ve seen in real code

