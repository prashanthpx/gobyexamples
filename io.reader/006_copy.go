package main

import (
	"bytes"
	"fmt"
	"io"
)

// Run with: go run io.reader/006_copy.go
func main() {
	src := bytes.NewReader([]byte("copy me"))
	var dst bytes.Buffer
	n, err := io.Copy(&dst, src)
	if err != nil { panic(err) }
	fmt.Println("copied:", n, dst.String())
}

