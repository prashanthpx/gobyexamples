package main

import (
	"fmt"
	"os"
)

func main() {
	path := "./fileio/examples/sample.txt"
	_ = os.WriteFile(path, []byte("hello\nworld\n"), 0644)

	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", b)
}

