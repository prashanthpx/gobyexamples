package main

import (
	"fmt"
	"io"
	"strings"
)

// Run with: go run io.reader/002_strings_reader_basic.go
func main() {
	r := strings.NewReader("abcde")
	b := make([]byte, 2)
	for {
		n, err := r.Read(b)
		if n > 0 { fmt.Printf("got: %q\n", b[:n]) }
		if err == io.EOF { break }
		if err != nil { panic(err) }
	}
}

