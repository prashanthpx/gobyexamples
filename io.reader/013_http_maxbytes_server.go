package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Run with: go run io.reader/013_http_maxbytes_server.go
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Enforce 1MB max body
		r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "too large or read error", http.StatusRequestEntityTooLarge)
			return
		}
		fmt.Fprintf(w, "got %d bytes", len(b))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

