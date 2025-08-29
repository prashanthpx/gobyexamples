package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	path := "./fileio/examples/sample_wc.txt"
	_ = os.WriteFile(path, []byte("alpha beta\ngamma delta alpha\n"), 0644)

	f, err := os.Open(path)
	if err != nil { panic(err) }
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)
	counts := map[string]int{}
	for s.Scan() { counts[s.Text()]++ }
	if err := s.Err(); err != nil { panic(err) }

	for k, v := range counts {
		fmt.Printf("%s: %d\n", k, v)
	}
}

