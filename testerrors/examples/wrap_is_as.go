package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// Run with: go run testerrors/examples/wrap_is_as.go
func main() {
	_, err := os.Open("missing.txt")
	if err != nil {
		err = fmt.Errorf("open failed: %w", err)
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("not found (Is)")
		}
		var p *fs.PathError
		if errors.As(err, &p) {
			fmt.Println("path:", p.Path)
		}
	}
}

