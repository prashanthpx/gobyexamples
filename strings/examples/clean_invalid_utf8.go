package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte{0x61, 0xff, 0x62, 0xc3} // a, invalid, b, dangling lead byte
	s := string(b)
	fmt.Println("valid?", utf8.ValidString(s))

	cleaned := bytes.ToValidUTF8([]byte(s), []byte{"?"[0]})
	fmt.Printf("cleaned: %q (bytes: % x)\n", string(cleaned), cleaned)
}

