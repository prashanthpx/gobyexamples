package main

import (
	"fmt"
	"time"
)

var jobs = make(chan int, 2)
var noOfJobs = 10

func worker(ch chan int) {
	fmt.Printf("\n line 12 len: %v", len(jobs))
	for i := 0; i < noOfJobs; i++ {
		jobs<-i
		fmt.Printf("\n line 15")
	}
	close(jobs)
}

func main() {
	go worker(jobs)
	for job := range jobs {
		fmt.Printf(" \n line 24 job : %v, time: %v", job, time.Now())
	}
	// For loop blocks until the channel is closed.
	// For channels, the iteration values produced are the successive values sent on the channel until the channel is closed. 
	// If the channel is nil, the range expression blocks forever.
	fmt.Printf(" \n line 25 ")
	//time.Sleep(4 * time.Second)
}