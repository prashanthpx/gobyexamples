package main

import (
	"fmt"
)

// Run with: go run wait_group/mistakes/oneshot_worker.go
// Demonstrates the one-shot select bug vs proper loop over channel.

type Job int

func workerBad(in <-chan Job) {
	select {
	case j := <-in:
		fmt.Println("processed (bad)", j)
	}
}

func workerGood(in <-chan Job) {
	for j := range in {
		fmt.Println("processed (good)", j)
	}
}

func main() {
	jobs := make(chan Job, 3)
	for i := 0; i < 3; i++ { jobs <- Job(i) }
	close(jobs)

	workerBad(jobs)

	jobs2 := make(chan Job, 3)
	for i := 0; i < 3; i++ { jobs2 <- Job(i) }
	close(jobs2)
	workerGood(jobs2)
}

