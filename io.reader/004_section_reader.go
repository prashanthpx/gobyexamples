package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

// Run with: go run io.reader/004_section_reader.go
func main() {
	blob := []byte("HEADER|BODY-DATA|FOOTER")
	r := bytes.NewReader(blob)

	// SectionReader over the BODY-DATA region
	s := io.NewSectionReader(r, int64(len("HEADER|")), int64(len("BODY-DATA")))
	b, _ := ioutil.ReadAll(s)
	fmt.Println(string(b))
}

