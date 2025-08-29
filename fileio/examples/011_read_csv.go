package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	path := "./fileio/examples/sample.csv"
	// Create a sample CSV file
	_ = os.WriteFile(path, []byte("name,age,city\nAlice,30,NY\nBob,25,SF\n"), 0644)

	f, err := os.Open(path)
	if err != nil { panic(err) }
	defer f.Close()

	r := csv.NewReader(f)
	r.TrimLeadingSpace = true
	// r.Comma = ',' // default
	records, err := r.ReadAll() // small files; for large files prefer streaming with Read()
	if err != nil { panic(err) }

	if len(records) == 0 { return }
	header := records[0]
	fmt.Println("header:", strings.Join(header, ", "))
	for i := 1; i < len(records); i++ {
		row := records[i]
		fmt.Printf("row %d: name=%s age=%s city=%s\n", i, row[0], row[1], row[2])
	}
}

