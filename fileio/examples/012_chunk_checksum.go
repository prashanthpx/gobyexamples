package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func main() {
	path := "./fileio/examples/large.txt"
	f, err := os.Create(path)
	if err != nil { panic(err) }
	// Write ~1MB sample
	for i := 0; i < 1024; i++ {
		fmt.Fprintf(f, "line %06d: %s\n", i, "abcdefghijklmnopqrstuvwxyz0123456789")
	}
	f.Close()

	f, err = os.Open(path)
	if err != nil { panic(err) }
	defer f.Close()

	h := sha256.New()
	buf := make([]byte, 64*1024) // 64KB chunk buffer
	for {
		n, err := f.Read(buf)
		if n > 0 {
			if _, werr := h.Write(buf[:n]); werr != nil { panic(werr) }
		}
		if err == io.EOF { break }
		if err != nil { panic(err) }
	}
	fmt.Printf("SHA256: %x\n", h.Sum(nil))
}

