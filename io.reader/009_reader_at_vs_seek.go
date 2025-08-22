package main

import (
	"bytes"
	"fmt"
)

// Run with: go run io.reader/009_reader_at_vs_seek.go
func main() {
	r := bytes.NewReader([]byte("ABCDEFG"))

	buf := make([]byte, 2)
	r.ReadAt(buf, 3) // doesn't change cursor
	fmt.Println("at:", string(buf))

	r.Read(buf) // reads from current pos (0 initially)
	fmt.Println("seeked read:", string(buf))
}

