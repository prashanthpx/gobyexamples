package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	path := "demo_file.txt"

	// 1) Open (create if missing) for read+write; truncate to start clean.
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0o644)
	must(err)
	defer f.Close()

	// 2) Write some bytes.
	n, err := f.WriteString("Hello, file!\nThis is a second line.\n")
	must(err)
	fmt.Printf("wrote %d bytes\n", n)

	// 3) Read entire file from the beginning.
	_, err = f.Seek(0, io.SeekStart)
	must(err)
	all, err := io.ReadAll(f)
	must(err)
	fmt.Printf("full contents:\n%s", all)

	// 4) Seek to an offset and read a fixed number of bytes.
	// Offset 7 lands on the 'f' in "Hello, file!"
	_, err = f.Seek(7, io.SeekStart)
	must(err)
	buf := make([]byte, 4)
	_, err = io.ReadFull(f, buf) // read exactly 4 bytes
	must(err)
	fmt.Printf("\nbytes at offset 7 (len=4): %q\n", string(buf))

	// 5) Seek to end and append more data.
	_, err = f.Seek(0, io.SeekEnd)
	must(err)
	_, err = f.WriteString("APPENDED\n")
	must(err)

	// 6) Re-read to confirm final contents.
	_, err = f.Seek(0, io.SeekStart)
	must(err)
	final, err := io.ReadAll(f)
	must(err)
	fmt.Printf("\nfinal contents:\n%s", final)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

