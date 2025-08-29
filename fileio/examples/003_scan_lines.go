package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	path := "./fileio/examples/sample_lines.txt"
	_ = os.WriteFile(path, []byte("first\nsecond\nthird\n"), 0644)

	f, err := os.Open(path)
	if err != nil { panic(err) }
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Buffer(make([]byte, 0, 64*1024), 1024*1024) // allow long lines up to 1MB
	for s.Scan() {
		fmt.Println("line:", s.Text())
	}
	if err := s.Err(); err != nil { panic(err) }
}

