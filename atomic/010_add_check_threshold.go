package main

import (
	"errors"
	"fmt"
	"sync/atomic"
)

// Run with: go run atomic/010_add_check_threshold.go
// Increment and check threshold using atomic.AddInt64/atomic.Int64.
var count atomic.Int64

func incAndCheck(limit int64) error {
	v := count.Add(1) // atomic increment; returns new value
	if v >= limit {
		return fmt.Errorf("too many (count=%d)", v)
	}
	return nil
}

func main() {
	limit := int64(5)
	for i := 0; i < 7; i++ {
		err := incAndCheck(limit)
		if err != nil {
			// wrap to show error handling pattern
			fmt.Println(errors.Unwrap(fmt.Errorf("wrap: %w", err)))
		} else {
			fmt.Println(nil)
		}
	}
}

