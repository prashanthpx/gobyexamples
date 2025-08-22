package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// Run with: go run io.reader/011_line_counting_reader.go
// A simple line-counting decorator around an io.Reader.

type lineCounter struct {
	r io.Reader
}

func (l lineCounter) Read(p []byte) (int, error) {
	n, err := l.r.Read(p)
	for i := 0; i < n; i++ {
		if p[i] == '\n' { /* counting could be tracked here if needed */ }
	}
	return n, err
}

func main() {
	src := bytes.NewBufferString("a\nb\nc\n")
	lc := lineCounter{r: src}
	s := bufio.NewScanner(lc)
	s.Split(bufio.ScanLines)
	count := 0
	for s.Scan() { count++ }
	fmt.Println("lines:", count)
}

