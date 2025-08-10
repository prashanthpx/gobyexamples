package main

import (
	"fmt"
	"os"
)

func main() {
	const name, age = "Kim", 22
	n, err := fmt.Fprintln(os.Stdout, name, "is", age, "years old.")

	// The n and err return values from Fprintln are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintln: %v\n", err)
	}
	fmt.Println(n, "bytes written.")

}

/*
Output
Kim is 22 years old.
21 bytes written.
*/

/*
Code Explanation:
- Purpose: Demonstrate fmt.Fprintln to an io.Writer and checking returned n, err
- Writes to os.Stdout, then prints the number of bytes written
- On error, logs to os.Stderr via fmt.Fprintf
*/
