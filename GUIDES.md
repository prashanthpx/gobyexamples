# Go by Examples and Mistakes — Guides Index

This repository teaches Go with runnable examples, intentional mistakes, and interview‑grade explanations. Each folder has a focused guide that:
- Explains fundamentals and advanced concepts
- Shows common mistakes and how to fix them (with interview‑ready language)
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

---

## In Progress / Planned (many are now completed)

- Channels: channels/ChannelsGuide.md
- Concurrency Patterns: wait_group/ConcurrencyGuide.md, job_queues/JobQueuesGuide.md, worker_queue/WorkerQueuesGuide.md
- Goroutines: goroutine/GoroutinesGuide.md
- Interfaces: interface/InterfacesGuide.md
- Methods: methods/MethodsGuide.md
- Operators: Operators/OperatorsGuide.md
- Packages & Modules: packages/PackagesGuide.md
- Printf & Formatting: printf/PrintfGuide.md
- Time: time/TimeGuide.md
- Conditionals & Control Flow: conditional/ConditionalsGuide.md
- Validation: validator/ValidatorGuide.md
- YAML & Encoding: yaml/YAMLGuide.md
- Error Handling: testerrors/ErrorHandlingGuide.md

(These will appear as the files are added. Check back or watch the repo for updates.)

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

