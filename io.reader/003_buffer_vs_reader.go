package main

import (
	"bytes"
	"fmt"
)

// Run with: go run io.reader/003_buffer_vs_reader.go
func main() {
	data := []byte("payload")

	buf := bytes.NewBuffer(nil)
	buf.WriteString("pre-")
	buf.Write(data)
	fmt.Println("buffer:", buf.String())

	r := bytes.NewReader(data)
	b := make([]byte, 4)
	r.Read(b)
	fmt.Printf("reader first: %q\n", b)
}

