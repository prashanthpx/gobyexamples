package main

import "fmt"

func main() {
	s := "👋🏽" // waving hand + skin tone modifier (two runes)
	fmt.Printf("s=%q\n", s)
	fmt.Printf("bytes: % x\n", []byte(s))
	for i, r := range s {
		fmt.Printf("i=%d  r=%U %q\n", i, r, r)
	}
}

