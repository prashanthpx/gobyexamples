package main

import (
	"bytes"
	"fmt"
	"io"
)

// Run with: go run io.reader/005_limit_tee_multi.go
func main() {
	// LimitReader: cap reads
	lr := io.LimitReader(bytes.NewReader([]byte("1234567890")), 4)
	b, _ := io.ReadAll(lr)
	fmt.Println("limited:", string(b))

	// TeeReader: duplicate stream to a sink
	src := bytes.NewReader([]byte("ABC"))
	var sink bytes.Buffer
	tee := io.TeeReader(src, &sink)
	io.ReadAll(tee) // drains src and writes to sink
	fmt.Println("tee sink:", sink.String())

	// MultiReader: concatenate
	m := io.MultiReader(bytes.NewReader([]byte("HELLO ")), bytes.NewReader([]byte("WORLD")))
	all, _ := io.ReadAll(m)
	fmt.Println(string(all))
}

