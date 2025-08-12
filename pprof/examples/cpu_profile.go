package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

// Run with:
//   go run pprof/examples/cpu_profile.go -cpuprofile cpu.out
// Then:
//   go tool pprof cpu.out
//   (pprof) top
//   (pprof) web  # if graphviz installed

func work() {
	var x int
	for i := 0; i < 1e7; i++ { x += i }
	_ = x
}

func main() {
	var cpu string
	flag.StringVar(&cpu, "cpuprofile", "", "write cpu profile to file")
	flag.Parse()

	if cpu != "" {
		f, err := os.Create(cpu)
		if err != nil { log.Fatal(err) }
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil { log.Fatal(err) }
		defer pprof.StopCPUProfile()
	}

	work()
}

