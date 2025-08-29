package main

import (
	"fmt"
	"os"
	"time"
)

// BAD: deferring Close inside a loop leaks file descriptors until the function returns.
func main() {
	for i := 0; i < 100; i++ {
		f, err := os.CreateTemp("", "leak-*.txt")
		if err != nil { panic(err) }
		defer f.Close() // BAD: defers all 100 closes until main returns
		fmt.Fprintf(f, "record %d\n", i)
	}

	fmt.Println("Created 100 temp files with deferred closes (bad). On Unix, you could run 'lsof -p <pid>' now to observe many open files.")
	time.Sleep(2 * time.Second) // Give you a moment to inspect with external tools
}

