package main

import (
	"bytes"
	"fmt"
	"io"
)

// Run with: go run io.reader/001_bytes_reader_basic.go
func main() {
	r := bytes.NewReader([]byte("Hello, Go Reader!"))
	buf := make([]byte, 5)
	n, _ := r.Read(buf)
	fmt.Printf("read %d: %q\n", n, buf)
	n, _ = r.Read(buf)
	fmt.Printf("read %d: %q\n", n, buf)
	// rewind and read all
	r.Seek(0, io.SeekStart)
	all, _ := io.ReadAll(r)
	fmt.Printf("all: %q\n", all)
}

