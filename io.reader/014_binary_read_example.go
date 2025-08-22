package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Run with: go run io.reader/014_binary_read_example.go
func main() {
	// Read a big-endian uint32 from an in-memory []byte using binary.Read
	buf := []byte{0x00, 0x00, 0x00, 0x2A} // 42
	r := bytes.NewReader(buf)
	var n uint32
	if err := binary.Read(r, binary.BigEndian, &n); err != nil {
		panic(err)
	}
	fmt.Println(n)
}

