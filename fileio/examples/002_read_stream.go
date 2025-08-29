package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	path := "./fileio/examples/sample_stream.txt"
	f, err := os.Create(path)
	if err != nil { panic(err) }
	for i := 0; i < 5; i++ {
		fmt.Fprintf(f, "line %d\n", i)
	}
	f.Close()

	f, err = os.Open(path)
	if err != nil { panic(err) }
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			if len(line) > 0 { fmt.Print(line) }
			break
		}
		if err != nil { panic(err) }
		fmt.Print(line)
	}
}

