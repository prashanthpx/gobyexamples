package main

import (
	"bytes"
	"fmt"
	"io"
)

// Run with: go run io.reader/007_pipe.go
func main() {
	pr, pw := io.Pipe()

	// writer goroutine
	go func() {
		defer pw.Close()
		io.Copy(pw, bytes.NewBufferString("streamed"))
	}()

	// reader
	b, _ := io.ReadAll(pr)
	fmt.Println(string(b))
}

