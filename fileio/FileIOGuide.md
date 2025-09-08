# File I/O in Go: Practical Guide (open, read, process, write) with runnable examples

Run these examples
- Read whole file: go run fileio/examples/001_read_all.go
- Stream file (Reader): go run fileio/examples/002_read_stream.go
- Scan lines robustly: go run fileio/examples/003_scan_lines.go
- Scan sentences (custom split): go run fileio/examples/004_scan_sentences.go
- Write (overwrite): go run fileio/examples/005_write_overwrite.go
- Append safely: go run fileio/examples/006_append.go
- Copy file (io.Copy): go run fileio/examples/007_copy_file.go
- Pitfall: defer Close in loop: go run fileio/examples/008_defer_in_loop_pitfall.go
- Word count (processing): go run fileio/examples/009_word_count.go
- Read CSV: go run fileio/examples/011_read_csv.go
- Chunked checksum (64KB): go run fileio/examples/012_chunk_checksum.go

- End-to-end demo: go run fileio/examples/015_end_to_end_demo.go

- End-to-end demo (buffered I/O): go run fileio/examples/016_end_to_end_buffered.go

- BAD: defer Close in loop (leaks fds): go run fileio/examples/013_fd_leak_bad.go
- GOOD: close per iteration: go run fileio/examples/014_fd_leak_good.go


---

## Table of Contents
1. [Basics: Opening and Closing Files](#toc-1-basics-open-close)
2. [Convenience vs Streaming APIs](#toc-2-convenience-vs-stream)
3. [Reading as Lines and Sentences](#toc-3-lines-sentences)
4. [Processing Content (examples)](#toc-4-processing)
5. [Writing, Appending, and Copying](#toc-5-writing)
6. [Common Gotchas (and fixes)](#toc-6-gotchas)
7. [Best Practices](#toc-7-best)
8. [Performance Notes](#toc-8-perf)
9. [Advanced: Custom SplitFunc and Large Files](#toc-9-advanced)
10. [End-to-end file demo](#toc-10-demo)

---

<a id="toc-1-basics-open-close"></a>

## 1) Basics: Opening and Closing Files

- Open for reading: f, err := os.Open(path) // read-only
- Open for create/write/truncate: f, err := os.Create(path)
- Always Close: use defer f.Close() right after successful Open/Create
- Check errors!

Pattern
```go
f, err := os.Open("data.txt")
if err != nil { return err }
defer f.Close()
// read from f (an *os.File implements io.Reader)
```

Closing nuances
- Close errors: On files, Close typically has no error on success; still check if you need to detect flush errors (writers)
- Do not defer Close inside a tight loop over many files (see Pitfalls)

---

<a id="toc-2-convenience-vs-stream"></a>

## 2) Convenience vs Streaming APIs

- Convenience (reads all bytes): os.ReadFile(path) ([]byte, error)
  - Simple, great for small files
- Streaming (progressive read): wrap *os.File with bufio.Reader or read in chunks
  - Scales to large files, supports backpressure/processing-on-the-fly

Examples to run
- Read-all: fileio/examples/001_read_all.go
- Stream-chunks: fileio/examples/002_read_stream.go

---

<a id="toc-3-lines-sentences"></a>

## 3) Reading as Lines and Sentences

- bufio.Scanner: convenient tokenization (lines, words) but has a default token limit (~64K)
- Increase buffer for long lines: scanner.Buffer(make([]byte, 0, 1024), 1024*1024) // 1MB
- Custom SplitFunc: implement sentence splitting on '.', '!' or '?'

Examples to run
- Scan lines robustly: fileio/examples/003_scan_lines.go
- Scan sentences (custom): fileio/examples/004_scan_sentences.go

---

<a id="toc-4-processing"></a>

## 4) Processing Content (examples)

- Word count with Scanner (ScanWords)
- Transform while streaming (upper-case lines and write)
- Filter lines (contains substring), count matches

Examples to run
- Word count: fileio/examples/009_word_count.go
- Transform+write (see write examples for patterns)

---

<a id="toc-5-writing"></a>

## 5) Writing, Appending, and Copying

- Overwrite: os.WriteFile or os.Create + Write
- Append: os.OpenFile with O_APPEND|O_WRONLY (create with O_CREATE if missing)
- Copy: io.Copy(dst, src) efficiently copies data
- Permissions: final argument to os.WriteFile/OpenFile is a permission mask (affected by umask)

Examples to run

Write options (OpenFile flags)
- Truncate/overwrite: O_CREATE|O_WRONLY|O_TRUNC
- Append: O_CREATE|O_WRONLY|O_APPEND
- Create new only (fail if exists): O_CREATE|O_EXCL|O_WRONLY
- Read/Write combo: O_CREATE|O_RDWR (pair with O_TRUNC to clear)

Write []byte to a file
- Small files: `os.WriteFile(path, data, perm)`
- Using an *os.File: `f.Write(data)` or `bufio.NewWriter(f).Write(data)` then `Flush`
- Using a reader source: `io.Copy(f, bytes.NewReader(data))` or `br.WriteTo(f)`

Example to run: fileio/examples/010_write_options.go

- Write overwrite: fileio/examples/005_write_overwrite.go
- Append: fileio/examples/006_append.go
- Copy: fileio/examples/007_copy_file.go

---

<a id="toc-6-gotchas"></a>

## 6) Common Gotchas (and fixes)

1) Forgetting to Close files
- Fix: defer f.Close() right after successful Open/Create
- In loops: close explicitly at end of iteration (do NOT defer inside the loop)

2) Using Scanner on very long lines
- Fix: increase buffer via scanner.Buffer(..., max)

3) Assuming Read fills the buffer
- io.Reader may return fewer bytes than requested without EOF; loop until done

4) Mixing text encodings (CRLF vs LF)
- Normalize if needed; Scanner splits on '\n' and handles '\r\n' too

5) Permissions surprises
- Remember umask can reduce effective permissions; test on target OS

6) Not flushing buffered writers
- bufio.Writer must be Flush()'d (defer w.Flush())

7) Defer in loops (fd leaks)
- See: fileio/examples/008_defer_in_loop_pitfall.go

---

<a id="toc-7-best"></a>

## 7) Best Practices

- Prefer os.ReadFile for small config/text files; prefer streaming for large inputs
- Place defer f.Close() immediately after a successful Open/Create
- For line/word parsing, start with bufio.Scanner; increase buffer for long tokens
- For large copies, io.Copy is optimized; optionally use io.CopyBuffer with a reusable buffer
- Wrap writers with bufio.NewWriter for many small writes; remember to Flush
- Make examples self-contained by creating sample input in tmp dir (as in this repo)

---

<a id="toc-8-perf"></a>

## 8) Performance Notes

- Avoid reading entire huge files into memory; stream and process in chunks
- Reuse buffers to reduce allocations (pool with sync.Pool for hot paths)
- Use io.Copy (uses sendfile/Splice on some OSes via internal optimizations)
- Use larger Scanner buffers or switch to bufio.Reader for very long records

---

<a id="toc-9-advanced"></a>

## 9) Advanced: Custom SplitFunc and Large Files

- Custom sentence splitter: implement bufio.SplitFunc to split on punctuation
- Very large records: prefer bufio.Reader.ReadSlice/ReadString/ReadBytes with delimiter to control memory
- Memory map (advanced, platform-specific): often unnecessary; benchmark before choosing


<a id="toc-10-demo"></a>

## 10) End-to-end file demo

Add this complete example at the end to see open, read, seek, and append in one place.

```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	path := "demo_file.txt"

	// 1) Open (create if missing) for read+write; truncate to start clean.
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0o644)
	must(err)
	defer f.Close()

	// 2) Write some bytes.
	n, err := f.WriteString("Hello, file!\nThis is a second line.\n")
	must(err)
	fmt.Printf("wrote %d bytes\n", n)

	// 3) Read entire file from the beginning.
	_, err = f.Seek(0, io.SeekStart)
	must(err)
	all, err := io.ReadAll(f)
	must(err)
	fmt.Printf("full contents:\n%s", all)

	// 4) Seek to an offset and read a fixed number of bytes.
	// Offset 7 lands on the 'f' in "Hello, file!"
	_, err = f.Seek(7, io.SeekStart)
	must(err)
	buf := make([]byte, 4)
	_, err = io.ReadFull(f, buf) // read exactly 4 bytes
	must(err)
	fmt.Printf("\nbytes at offset 7 (len=4): %q\n", string(buf))

	// 5) Seek to end and append more data.
	_, err = f.Seek(0, io.SeekEnd)
	must(err)
	_, err = f.WriteString("APPENDED\n")
	must(err)

	// 6) Re-read to confirm final contents.
	_, err = f.Seek(0, io.SeekStart)
	must(err)
	final, err := io.ReadAll(f)
	must(err)
	fmt.Printf("\nfinal contents:\n%s", final)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
```

See: fileio/examples/004_scan_sentences.go, 002_read_stream.go

