package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	fmt.Printf("line 9")
	for j := range jobs {
		fmt.Printf(" entering worker")
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		fmt.Printf("line 22")
		go worker(w, jobs, results)
	}
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= 9; a++ {
		<-results
	}

}

/*
Output (timing-dependent)
line 22line 22line 22line 9 entering workerworker 3 processing job 1
line 9 entering workerworker 1 processing job 2
line 9 entering workerworker 2 processing job 3
 entering worker entering workerworker 1 processing job 6
worker 2 processing job 4
 entering workerworker 3 processing job 5
 entering workerworker 3 processing job 7
 entering workerworker 1 processing job 8
 entering workerworker 2 processing job 9
*/

/*
Code Explanation:
- Purpose: Worker pool pattern with buffered job/result channels
- Launch 3 workers reading from jobs and writing to results; main enqueues 9 jobs and drains results
- Print interleaving varies with scheduling; Sleep simulates work
*/
