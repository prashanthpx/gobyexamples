package main

import (
	"fmt"
	"os"
)

func main() {
	path := "./fileio/examples/out_overwrite.txt"
	if err := os.WriteFile(path, []byte("initial\n"), 0644); err != nil { panic(err) }
	if err := os.WriteFile(path, []byte("overwritten\n"), 0644); err != nil { panic(err) }
	b, _ := os.ReadFile(path)
	fmt.Printf("%s", b)
}

