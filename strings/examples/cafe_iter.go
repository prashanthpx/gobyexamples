package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "café"
	fmt.Printf("s=%q\n", s)
	fmt.Printf("bytes: % x\n", []byte(s))
	fmt.Println("len(s) bytes:", len(s))
	fmt.Println("runes:", utf8.RuneCountInString(s))

	for i, r := range s {
		fmt.Printf("i=%d  r=%U %q\n", i, r, r)
	}
}

