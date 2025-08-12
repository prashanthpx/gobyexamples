package main

import (
	"fmt"
	"unicode/utf8"
)

// Run with: go run strings/examples/grapheme_demo.go
// Demonstrates combining characters vs single code points.
func main() {
	base := "e\u0301" // e + combining acute accent
	a := "é"           // precomposed
	fmt.Println("bytes base,a:", len(base), len(a))
	fmt.Println("runes base,a:", utf8.RuneCountInString(base), utf8.RuneCountInString(a))
	for i, r := range base { fmt.Printf("base[%d]=%U %q\n", i, r, string(r)) }
	for i, r := range a { fmt.Printf("a[%d]=%U %q\n", i, r, string(r)) }
}

