# Go I/O Readers and Friends: Practical Guide (bytes.Reader, strings.Reader, SectionReader, Buffer, and more)

Run these examples
- bytes.Reader basics: go run io.reader/001_bytes_reader_basic.go
- strings.Reader basics: go run io.reader/002_strings_reader_basic.go
- Buffer vs Reader: go run io.reader/003_buffer_vs_reader.go
- SectionReader & ReaderAt: go run io.reader/004_section_reader.go
- Tee/Limited/Multi Reader/Writer: go run io.reader/005_limit_tee_multi.go
- io.Copy fundamentals: go run io.reader/006_copy.go
- io.Pipe streaming: go run io.reader/007_pipe.go
- HTTP body from memory: go run io.reader/008_http_post_reader.go
- ReaderAt vs Seek cursor: go run io.reader/009_reader_at_vs_seek.go
- PNG-like header parsing: go run io.reader/010_png_header_reader.go
- Line-counting decorator: go run io.reader/011_line_counting_reader.go
- Benchmark: strings.NewReader vs bytes.NewReader([]byte): go test -bench=. -benchmem ./io.reader
- Server-side max body size: go run io.reader/013_http_maxbytes_server.go

---

## Table of Contents
1. Reader family overview (mental model)
2. bytes.Reader vs strings.Reader vs bytes.Buffer
3. io.SectionReader (fixed window), io.LimitReader, io.TeeReader
4. io.Copy, MultiReader/MultiWriter, Pipe
5. ReaderAt, Seeker, and cursor semantics
6. Choosing the right abstraction (cheat sheet)
7. Common mistakes and gotchas
8. Best practices and performance notes
9. Advanced interview questions
10. Interview-ready advantages of Readers
11. “File-like” usage of bytes.Reader and limitations

---

## 1) Reader family overview (mental model)

- io.Reader: minimal pull-based stream primitive; Read(p []byte) (n int, err error)
- io.Writer: dual abstraction for push-based writes; Write(p []byte) (n int, err error)
- io.Seeker: move a cursor within a stream; used with random access
- io.ReaderAt: random access reads that do not change the cursor
- io.Closer: close underlying resource (files, network, compressed streams)

Create Readers from different sources:
- In-memory bytes: bytes.NewReader([]byte)
- In-memory strings: strings.NewReader(string)
- Files: os.Open returns *os.File (Reader, Writer, Seeker, ReaderAt)
- Network: net.Conn implements Reader/Writer

---

## 2) bytes.Reader vs strings.Reader vs bytes.Buffer

- bytes.Reader: read-only view over []byte with Seeker, ReaderAt, Byte/Rune scanners. Lightweight; perfect when you already have bytes
- strings.Reader: same as bytes.Reader but for string; avoids []byte conversion/allocations
- bytes.Buffer: growable read/write buffer; no Seek. Best for building data dynamically, then reading

See: io.reader/001_bytes_reader_basic.go, 002_strings_reader_basic.go, 003_buffer_vs_reader.go

---

## 3) io.SectionReader, io.LimitReader, io.TeeReader

- SectionReader: fixed-length window over a ReaderAt (e.g., a sub-file). Useful for parsing headers/regions
- LimitReader: cap stream to N bytes, protecting decoders from over-reads
- TeeReader: mirror reads into a Writer (e.g., logging, checksums) while consuming

See: io.reader/004_section_reader.go, 005_limit_tee_multi.go

---

## 4) io.Copy, MultiReader/MultiWriter, Pipe

- io.Copy(dst, src) efficiently copies until EOF (uses specialized interfaces when available)
- MultiReader concatenates readers; MultiWriter fans out writes to multiple writers
- Pipe connects an io.Reader to an io.Writer in memory with backpressure; great for streaming producer/consumer in the same process

See: io.reader/005_limit_tee_multi.go, 006_copy.go, 007_pipe.go

---

## 5) ReaderAt vs Seeker and cursor semantics

- ReaderAt reads at offsets without mutating the current position; safe for concurrent reads
- Seeker moves a shared cursor; not safe to use from multiple goroutines without external synchronization
- A bytes.Reader supports both; files usually support both

See: io.reader/009_reader_at_vs_seek.go, 004_section_reader.go

---

## 6) Choosing the right abstraction (cheat sheet)

- Already have []byte → bytes.NewReader (read/seek) or bytes.NewBuffer (read/write)
- Already have string → strings.NewReader (no []byte conversion)
- Need fixed window → io.NewSectionReader over a ReaderAt source
- Need to cap input size → io.LimitReader
- Need to duplicate stream to a Writer → io.TeeReader
- Need to concatenate readers → io.MultiReader
- Need to broadcast writes → io.MultiWriter
- Need to stream between goroutines → io.Pipe

---

## 7) Common mistakes and gotchas

1) Assuming Read fills the buffer
- Read may return with n < len(p) without EOF; loop until you have enough or hit EOF

2) Forgetting to close
- Close readers that own resources (files, HTTP response bodies, compressors). bytes.Reader/strings.Reader don’t need closing

3) Converting strings ↔ []byte unnecessarily
- Prefer strings.NewReader for strings, bytes.NewReader for []byte to avoid allocations

4) Using bytes.Buffer when only reading
- If you only need to read, bytes.Reader is lighter and supports Seek/ReaderAt

5) Concurrent use of a single Reader without coordination
- Use ReaderAt for parallel reads or coordinate with a mutex around Seeker

6) Reading huge bodies without bounds
- Wrap with io.LimitReader and/or enforce MaxBytesReader in HTTP servers

---

## 8) Best practices and performance notes

- Prefer io.Copy/CopyN over manual loops when appropriate; they use optimized paths
- Reuse buffers (sync.Pool) in tight loops to reduce allocations
- For logging/metrics while reading, use io.TeeReader to avoid double reads
- For large files, prefer ReaderAt with goroutines to parallelize segments (be mindful of storage)
- Avoid ReadAll on untrusted/unbounded inputs; stream/process incrementally

---

## 9) Advanced interview questions

Bonus: interview-ready points about bytes.Reader and “file-like” usage
- See runnable: io.reader/014_binary_read_example.go


1) bytes.Reader vs bytes.Buffer vs strings.Reader — when do you pick each? Cost trade-offs?
2) Explain Read contract (short reads), EOF handling, and when Read returns (0, nil)
3) ReaderAt vs Seeker semantics; concurrency safety and typical use cases
4) How would you cap an HTTP request body size defensively? Show both client and server
5) How does io.Copy optimize using interfaces like WriterTo/ReaderFrom?
6) Implement a line-counting reader decorator; how would you structure it?
7) When would you prefer io.Pipe over a temporary file or in-memory buffer?




## 10) Interview-ready advantages of Readers (bytes.Reader in focus)

1) Turn raw []byte into a stream (no copy)
- Many APIs expect io.Reader
- Example: http.Post(url, "text/plain", bytes.NewReader([]byte("hello")))

2) Automatic cursor management
- Successive Read calls advance internally (no manual indices)

3) Random access with Seek
- Example: r.Seek(5, io.SeekStart) to jump to an offset

4) Random access without moving cursor via ReaderAt
- Example: r.ReadAt(buf, off)

5) Re-read without rebuild
- Seek(0, io.SeekStart) to replay the same bytes (e.g., retries)

6) Interoperates with io utilities
- io.Copy, compressors, encoders/decoders, binary.Read all consume io.Reader

7) Great for tests/mocking
- Wrap literals/slices: bytes.NewReader([]byte("a,b,c\n"))

8) Lightweight and allocation-friendly
- Wraps an existing slice; doesn’t grow like bytes.Buffer

9) Safer than manual slicing for streaming
- Fewer off-by-one/index bugs; clear stream semantics

10) Works like an in-memory file for read paths
- Implements Reader, Seeker, ReaderAt, scanners

See also runnable examples in this folder; for binary.Read demo, run: go run io.reader/014_binary_read_example.go

---

## 11) “File-like” usage of bytes.Reader and limitations

- Many file APIs operate on interfaces (Reader/Seeker/ReaderAt), not concrete *os.File
- bytes.Reader satisfies read-side interfaces; you can plug it into most read-only file workflows

Example (binary.Read works on files or in-memory data):
<code>
var num uint32
binary.Read(bytes.NewReader([]byte{0,0,0,42}), binary.BigEndian, &num)
// prints 42
</code>

Caveats
- bytes.Reader is read-only: it does not implement io.Writer or methods like Truncate, Sync
- No filesystem semantics (paths, permissions, rename)
- For writing/mutating, use *os.File or an in-memory buffer (bytes.Buffer) as appropriate

Interview soundbite
- “bytes.Reader is an in-memory, read-only, file-like view over a []byte. It implements the key read/seek interfaces so I can reuse all of Go’s I/O ecosystem without touching disk.”