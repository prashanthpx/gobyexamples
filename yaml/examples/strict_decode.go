package main

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v3"
)

// Run with: go run yaml/examples/strict_decode.go
// Demonstrates strict decoding that fails on unknown fields.

type Item struct {
	Name  string   `yaml:"name"`
	Count int      `yaml:"count"`
	Tags  []string `yaml:"tags"`
}

func main() {
	input := []byte("name: book\ncount: 1\nextra: oops\n")
	dec := yaml.NewDecoder(bytes.NewReader(input))
	dec.KnownFields(true)
	var it Item
	if err := dec.Decode(&it); err != nil {
		fmt.Println("strict error:", err) // reports unknown field "extra"
	}
}

