package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Construct a string with invalid UTF-8 bytes:
	// 0x61 = 'a', 0xff = invalid, 0x62 = 'b', 0xc3 = leading byte without continuation
	b := []byte{0x61, 0xff, 0x62, 0xc3}
	s := string(b)

	fmt.Printf("bytes: % x\n", b)
	fmt.Println("utf8.ValidString:", utf8.ValidString(s))

	fmt.Println("range over s (invalid sequences produce utf8.RuneError):")
	for i, r := range s {
		_, w := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("i=%d r=%U %q width=%d\n", i, r, r, w)
	}

	fmt.Println("manual decode with utf8.DecodeRuneInString:")
	for i := 0; i < len(s); {
		r, w := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("i=%d r=%U %q width=%d\n", i, r, r, w)
		i += w
	}

	// Tip: To replace invalid bytes, you can use bytes.ToValidUTF8
	// cleaned := bytes.ToValidUTF8([]byte(s), []byte("?"))
	// fmt.Printf("cleaned: %q\n", string(cleaned))
}

