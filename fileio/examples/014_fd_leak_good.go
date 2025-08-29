package main

import (
	"fmt"
	"os"
)

// GOOD: close each file at end of each iteration to avoid leaks.
func main() {
	for i := 0; i < 100; i++ {
		f, err := os.CreateTemp("", "noleak-*.txt")
		if err != nil { panic(err) }
		fmt.Fprintf(f, "record %d\n", i)
		f.Close() // explicit close per iteration
	}
	fmt.Println("Created and closed 100 temp files safely.")
}

