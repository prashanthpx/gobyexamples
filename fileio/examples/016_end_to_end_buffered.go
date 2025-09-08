package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	path := "demo_buffered.txt"

	// 1) Open (create if missing) for read+write; truncate to start clean.
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0o644)
	must(err)
	defer f.Close()

	// 2) Buffered write some bytes (remember to Flush).
	bw := bufio.NewWriter(f)
	_, err = bw.WriteString("Hello, buffered file!\nThis is a second line.\n")
	must(err)
	must(bw.Flush())
	fmt.Println("buffered write complete")

	// 3) Read entire file from the beginning using a buffered reader.
	_, err = f.Seek(0, io.SeekStart)
	must(err)
	br := bufio.NewReader(f)
	all, err := io.ReadAll(br)
	must(err)
	fmt.Printf("full contents:\n%s", all)

	// 4) Seek to an offset and read a fixed number of bytes (recreate reader after Seek).
	_, err = f.Seek(7, io.SeekStart)
	must(err)
	br = bufio.NewReader(f)
	buf := make([]byte, 6)
	_, err = io.ReadFull(br, buf)
	must(err)
	fmt.Printf("\nbytes at offset 7 (len=6): %q\n", string(buf))

	// 5) Seek to end and append more data using a buffered writer.
	_, err = f.Seek(0, io.SeekEnd)
	must(err)
	bw = bufio.NewWriter(f)
	_, err = bw.WriteString("BUFFERED_APPEND\n")
	must(err)
	must(bw.Flush())

	// 6) Re-read to confirm final contents.
	_, err = f.Seek(0, io.SeekStart)
	must(err)
	br = bufio.NewReader(f)
	final, err := io.ReadAll(br)
	must(err)
	fmt.Printf("\nfinal contents:\n%s", final)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

