# Go by Examples and Mistakes — Guides Index

This repository teaches Go with runnable examples, intentional mistakes, and practical explanations. Each folder has a focused guide that:
- Explains fundamentals and advanced concepts
- Shows common mistakes and how to fix them
- Provides best practices and performance notes
- Keeps all examples compilable and exact

Use this index to jump to the guide for each topic.

---

## Completed Guides

- Mistakes Index: [MISTAKES_INDEX.md](MISTAKES_INDEX.md)

- Arrays: [arrays/ArraysGuide.md](arrays/ArraysGuide.md)
- Functions: [functions/FunctionsGuide.md](functions/FunctionsGuide.md)
- Maps: [maps/MapsGuide.md](maps/MapsGuide.md)
- Pointers: [Pointers/PointersGuide.md](Pointers/PointersGuide.md)
- Slices: [slices/SliceGuide.md](slices/SliceGuide.md)
- Structs: [structures/StructuresGuide.md](structures/StructuresGuide.md)
- Mutexes: [mutex/MutexGuide.md](mutex/MutexGuide.md)
- Types: [types/TypesGuide.md](types/TypesGuide.md)
- I/O Readers: [io.reader/IOReaderGuide.md](io.reader/IOReaderGuide.md)
- Atomics: [atomic/AtomicGuide.md](atomic/AtomicGuide.md)

  - Why lock-free matters: [atomic/AtomicGuide.md#why-lock-free-matters](atomic/AtomicGuide.md#why-lock-free-matters)

  - Jump to Benchmarks: [atomic/AtomicGuide.md#benchmarking](atomic/AtomicGuide.md#benchmarking)

- Goroutines: [goroutine/GoroutinesGuide.md](goroutine/GoroutinesGuide.md)
- Channels: [channels/ChannelsGuide.md](channels/ChannelsGuide.md)
- Concurrency Patterns: [wait_group/ConcurrencyGuide.md](wait_group/ConcurrencyGuide.md), [job_queues/JobQueuesGuide.md](job_queues/JobQueuesGuide.md), [worker_queue/WorkerQueuesGuide.md](worker_queue/WorkerQueuesGuide.md)
- Interfaces: [interface/InterfacesGuide.md](interface/InterfacesGuide.md)
- Methods: [methods/MethodsGuide.md](methods/MethodsGuide.md)
- Operators: [Operators/OperatorsGuide.md](Operators/OperatorsGuide.md)
- Packages & Modules: [packages/PackagesGuide.md](packages/PackagesGuide.md)
- Printf & Formatting: [printf/PrintfGuide.md](printf/PrintfGuide.md)
- Time: [time/TimeGuide.md](time/TimeGuide.md)
- Conditionals & Control Flow: [conditional/ConditionalsGuide.md](conditional/ConditionalsGuide.md)
- Validation: [validator/ValidatorGuide.md](validator/ValidatorGuide.md)
- YAML & Encoding: [yaml/YAMLGuide.md](yaml/YAMLGuide.md)
- Error Handling: [testerrors/ErrorHandlingGuide.md](testerrors/ErrorHandlingGuide.md)
- Strings: [strings/StringsGuide.md](strings/StringsGuide.md)

- Linked Lists: [linkedlist/LinkedListGuide.md](linkedlist/LinkedListGuide.md)

- File I/O (files and processing): [fileio/FileIOGuide.md](fileio/FileIOGuide.md)


---

## In Progress / Planned

- Add more runnable examples and microbenchmarks where useful
- Expand profiling coverage (pprof CPU/heap, traces) and link from guides
- Add more “Mistakes Index” cross-links across guides
- Deepen generics coverage (constraint design, performance)
- Add unsafe.Pointer cautionary appendix (advanced; opt-in)
- Expand YAML/JSON custom marshal/unmarshal recipes
- More concurrency patterns (errgroup variants, backpressure, retries)
- More IO Reader decorators and HTTP defensive patterns (MaxBytesReader, timeouts)

(Hint: If you want a specific topic prioritized, open an issue and we’ll queue it up.)

---

## Guide Structure (What to Expect)

Each guide follows a consistent structure so you can read quickly or study in depth:
1. Introduction & Why It Matters
2. Core Concepts with runnable examples
3. Common Mistakes (and how to fix them)
4. Best Practices and performance implications
5. Advanced Challenge Questions (with explanations)

All examples are designed to compile. Where a mistake is shown, the fixed version immediately follows.

---

## Tips for Using These Guides

- Run examples with:
  - `go run path/to/file.go` (each example is self‑contained)
- Compare outputs with the Output blocks appended to many examples in the repo
- Read “Common Mistakes” and “Challenge Questions” to sharpen understanding and articulation

---

## Contributing / Extending

- Keep examples small, runnable, and focused on a single idea
- Pair every intentional mistake with a corrected version and an explanation
- Prefer keyed struct literals and explicit types in guides for clarity
- Include performance notes (allocation, GC, bounds checks) when relevant

If you want a top‑level README that links these guides in addition to this index, we can add that too.

