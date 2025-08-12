package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

// Run with:
//   go run pprof/examples/heap_profile.go -memprofile mem.out
// Then:
//   go tool pprof mem.out
//   (pprof) top

func leak() [][]byte {
	var hold [][]byte
	for i := 0; i < 1e4; i++ {
		b := make([]byte, 1024) // 1KB each
		hold = append(hold, b)
	}
	return hold
}

func main() {
	var mem string
	flag.StringVar(&mem, "memprofile", "", "write heap profile to file")
	flag.Parse()

	if mem != "" {
		f, err := os.Create(mem)
		if err != nil { log.Fatal(err) }
		defer f.Close()
		pprof.WriteHeapProfile(f)
	}

	_ = leak()
}

