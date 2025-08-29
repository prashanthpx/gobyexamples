package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

// sentenceSplit splits on '.', '!' or '?' and trims spaces.
func sentenceSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Find sentence terminator
	for i, b := range data {
		switch b {
		case '.', '!', '?':
			// include terminator
			adv := i + 1
			return adv, bytes.TrimSpace(data[:adv]), nil
		}
	}
	// If at EOF, return any remaining data
	if atEOF && len(data) > 0 {
		return len(data), bytes.TrimSpace(data), nil
	}
	// request more data
	return 0, nil, nil
}

func main() {
	path := "./fileio/examples/sample_sentences.txt"
	_ = os.WriteFile(path, []byte("Hello world. How are you? I am fine!"), 0644)

	f, err := os.Open(path)
	if err != nil { panic(err) }
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(sentenceSplit)
	for s.Scan() {
		fmt.Printf("sentence: %q\n", s.Text())
	}
	if err := s.Err(); err != nil { panic(err) }
}

