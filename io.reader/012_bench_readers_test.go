package main

import (
	"bytes"
	"strings"
	"testing"
)

// Run with: go test -bench=. -benchmem ./io.reader

func BenchmarkStringsReader(b *testing.B) {
	s := strings.Repeat("x", 1<<16)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := strings.NewReader(s)
		buf := make([]byte, 1024)
		for {
			_, err := r.Read(buf)
			if err != nil { break }
		}
	}
}

func BenchmarkBytesReaderFromBytes(b *testing.B) {
	bts := bytes.Repeat([]byte{'x'}, 1<<16)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := bytes.NewReader(bts)
		buf := make([]byte, 1024)
		for {
			_, err := r.Read(buf)
			if err != nil { break }
		}
	}
}

