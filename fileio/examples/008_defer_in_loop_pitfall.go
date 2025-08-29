package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < 3; i++ {
		f, err := os.CreateTemp("", "tmp-*.txt")
		if err != nil { panic(err) }

		// PITFALL: do NOT defer in loops (defers run at function return)
		// defer f.Close()
		// Correct: close explicitly at end of iteration
		fmt.Fprintf(f, "hello %d\n", i)
		f.Close()
	}
	fmt.Println("done; no leaked descriptors")
}

