package main

import (
	"encoding/json"
	"fmt"
)

// Run with: go run structures/examples/json_omitempty.go
// Demonstrates omitempty behavior for zero values.

type Item struct {
	Name  string   `json:"name"`
	Count int      `json:"count,omitempty"`
	Tags  []string `json:"tags,omitempty"`
}

func main() {
	// Count=0 and Tags=nil will be omitted
	b1, _ := json.Marshal(Item{Name: "book"})
	fmt.Println(string(b1)) // {"name":"book"}

	// Empty slice is omitted as well
	b2, _ := json.Marshal(Item{Name: "book", Tags: []string{}})
	fmt.Println(string(b2)) // {"name":"book"}

	// Non-zero Count appears
	b3, _ := json.Marshal(Item{Name: "book", Count: 1})
	fmt.Println(string(b3)) // {"name":"book","count":1}

	// Force presence by using pointers or custom marshaling if required
	type WithPtr struct { Count *int `json:"count,omitempty"` }
	zero := 0
	b4, _ := json.Marshal(WithPtr{Count: &zero})
	fmt.Println(string(b4)) // {"count":0}
}

