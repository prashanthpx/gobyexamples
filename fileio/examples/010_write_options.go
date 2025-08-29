package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func must(err error) { if err != nil { panic(err) } }

func main() {
	// 1) Truncate/overwrite
	f1, err := os.OpenFile("./fileio/examples/out_trunc.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	must(err)
	_, err = f1.WriteString("overwritten contents\n")
	must(err)
	f1.Close()

	// 2) Append
	f2, err := os.OpenFile("./fileio/examples/out_append_opts.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	must(err)
	for i := 0; i < 2; i++ { fmt.Fprintf(f2, "line %d\n", i) }
	f2.Close()

	// 3) Create new file, fail if exists (O_EXCL)
	f3, err := os.OpenFile("./fileio/examples/out_new_only.txt", os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err == nil {
		f3.Write([]byte("first write\n"))
		f3.Close()
	} else {
		fmt.Println("out_new_only.txt already exists; O_EXCL prevented overwrite")
	}

	// 4) Read/Write combination and buffered writer
	f4, err := os.OpenFile("./fileio/examples/out_buf.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	must(err)
	bw := bufio.NewWriter(f4)
	bw.WriteString("hello ")
	bw.Write([]byte("world\n"))
	bw.Flush() // important
	f4.Close()

	// 5) Write []byte via io.Copy from a bytes.Reader
	data := []byte("copied bytes to file\n")
	f5, err := os.OpenFile("./fileio/examples/out_copy_bytes.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	must(err)
	br := bytes.NewReader(data)
	_, err = br.WriteTo(f5) // equivalent to io.Copy(f5, br)
	must(err)
	f5.Close()

	fmt.Println("write options demo complete")
}

