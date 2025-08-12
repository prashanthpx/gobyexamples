package main

import (
	"fmt"
	"time"
)

// Run with: go run time/examples/parse_location.go
func main() {
	layout := "2006-01-02 15:04"
	u, _ := time.Parse(layout, "2024-11-05 09:30")
	loc, _ := time.LoadLocation("America/New_York")
	n, _ := time.ParseInLocation(layout, "2024-11-05 09:30", loc)
	fmt.Println("Parse loc:", u.Location())
	fmt.Println("ParseInLocation loc:", n.Location())
}

