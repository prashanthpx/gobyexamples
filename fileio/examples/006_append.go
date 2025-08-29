package main

import (
	"fmt"
	"os"
)

func main() {
	path := "./fileio/examples/out_append.txt"
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil { panic(err) }
	defer f.Close()

	for i := 0; i < 3; i++ {
		fmt.Fprintf(f, "line %d\n", i)
	}

	b, _ := os.ReadFile(path)
	fmt.Printf("%s", b)
}

