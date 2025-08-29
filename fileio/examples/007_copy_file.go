package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	srcPath := "./fileio/examples/src.txt"
	dstPath := "./fileio/examples/dst.txt"
	_ = os.WriteFile(srcPath, []byte("copy me\n"), 0644)

	src, err := os.Open(srcPath)
	if err != nil { panic(err) }
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil { panic(err) }
	defer dst.Close()

	n, err := io.Copy(dst, src)
	if err != nil { panic(err) }
	fmt.Println("copied bytes:", n)
}

