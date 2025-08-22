package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Run with: go run io.reader/008_http_post_reader.go
func main() {
	payload := map[string]string{"name":"Alice"}
	b, _ := json.Marshal(payload)
	body := bytes.NewReader(b)
	_, _ = http.Post("https://example.com", "application/json", body)
}

