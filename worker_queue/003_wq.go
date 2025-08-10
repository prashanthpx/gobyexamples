package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	//"context"
)

type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup, wid int) {
	fmt.Printf(" \n line 35 worker id : %v - len: %v", wid, len(jobs))
	select {
	case job := <-jobs:
		output := Result{job, digits(job.randomno)}
		fmt.Printf(" \n line 38 worker id: %v - id: %v, job rand no: %v ", wid, job.id, job.randomno)
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	//ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		fmt.Printf(" \n line 54 allocate - job id: %v, rand no: %v", i, randomno)
		jobs <- job
		fmt.Printf("  len(jobs): %v, cap(jobs): %v", len(jobs), cap(jobs))
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("\n Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func main() {
	startTime := time.Now()
	noOfJobs := 10
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 3
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

/*
Output (timing/random-dependent)
 ... allocate logs ...
 line 35 worker id : <id> - len: <n>
 ... some jobs processed (others dropped due to single receive) ...
 total time taken  2.00xxxx seconds
*/

/*
Code Explanation:
- Purpose: Illustrate a flawed worker that only receives one job via select
- The worker uses select to take a single job, process it, and then returns; many jobs remain unprocessed
- Contrast with 001_wq.go where workers range over jobs to drain the queue
*/
