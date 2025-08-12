package main

import (
	"io"
	"net/http"
)

// Run with: go run functions/mistakes/defer_loop_http.go
// Demonstrates why deferring inside loops can leak connections; shows a safe pattern.

var urls = []string{
	"https://example.com",
	"https://example.org",
}

// bad shows defers stacking until function return (not advisable for many iterations)
func bad() error {
	for _, u := range urls {
		resp, err := http.Get(u)
		if err != nil { return err }
		defer resp.Body.Close()
		io.Copy(io.Discard, resp.Body)
	}
	return nil
}

// good wraps each iteration in its own function so Close runs promptly
func good() error {
	for _, u := range urls {
		if err := func() error {
			resp, err := http.Get(u)
			if err != nil { return err }
			defer resp.Body.Close()
			_, _ = io.Copy(io.Discard, resp.Body)
			return nil
		}(); err != nil {
			return err
		}
	}
	return nil
}

func main() { _ = bad(); _ = good() }

