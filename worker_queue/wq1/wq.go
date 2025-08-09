package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// Job - interface for job processing
type Job interface {
	Run()
}

// Worker - the worker threads that actually process the jobs
type Worker struct {
	done      *sync.WaitGroup
	readyPool chan Job
	terminate chan bool
}
// JobQueue - a queue for enqueueing jobs to be processed
type JobQueue struct {
	inputQueue       chan Job
	readyPool        chan Job
	workers          []*Worker
	schedulerStopped sync.WaitGroup
	workersStopped   *sync.WaitGroup
	terminate        chan bool
}

// NewJobQueue - creates a new job queue
func NewJobQueue(maxWorkers int) *JobQueue {
	workersStopped := sync.WaitGroup{}
	readyPool := make(chan Job, maxWorkers)
	//readyPool := make(chan Job, 5)

	//workers := make([]*Worker, maxWorkers)
	workers := make([]*Worker, 3)

	// for i := 0; i < maxWorkers; i++ {
	for i := 0; i < 3; i++ {
		workers[i] = NewWorkerPool(readyPool, &workersStopped)
	}
	return &JobQueue{
		inputQueue:       make(chan Job),
		readyPool:        readyPool,
		workers:          workers,
		schedulerStopped: sync.WaitGroup{},
		workersStopped:   &workersStopped,
		terminate:        make(chan bool),
	}
}

// Start - starts the worker routines and dispatcher routine
func (q *JobQueue) Start() {
	for i := 0; i < len(q.workers); i++ {
		fmt.Println(" line 56")
		q.workers[i].Start()
	}
	go q.schedule()
}

// Stop - stops the workers and dispatcher routine
func (q *JobQueue) Stop() {
	q.terminate <- true
	q.schedulerStopped.Wait()
}

func (q *JobQueue) schedule() {
	q.schedulerStopped.Add(1)
	for {
		select {
		case job := <-q.inputQueue:
			// Fetch from internal queue and push to readyPool
			fmt.Printf("line 75 Passed to readyPool: %v\n", job)
			q.readyPool <- job
		case <-q.terminate:
			for i := 0; i < len(q.workers); i++ {
				q.workers[i].Stop()
			}
			q.workersStopped.Wait()
			q.schedulerStopped.Done()
			return
		}
	}
}

// Submit - adds a new job to be processed
func (q *JobQueue) Submit(job Job) {
	q.inputQueue <- job
}

// NewWorkerPool - creates a new worker pool
func NewWorkerPool(readyPool chan Job, done *sync.WaitGroup) *Worker {
	return &Worker{
		done:      done,
		readyPool: readyPool,
		terminate: make(chan bool),
	}
}

// Start - begins the job processing loop for the worker
func (w *Worker) Start() {
	w.done.Add(1)
	go func() {
		for {
			select {
			case job := <-w.readyPool:
				job.Run()
			case <-w.terminate:
				w.done.Done()
				return
			}
		}
	}()
}

// Stop - stops the worker
func (w *Worker) Stop() {
	w.terminate <- true
}

// ================================== //
// TestJob - holds only an ID to show state
type TestJob struct {
	ID string
}

// Process - test process function
func (t *TestJob) Run() {
	fmt.Printf("Processing job '%s'\n", t.ID)
	time.Sleep(1 * time.Minute)
	fmt.Printf("Completed job '%s'\n", t.ID)
}

func main() {
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	// queue := NewJobQueue(runtime.NumCPU())
	queue := NewJobQueue(10)

	queue.Start()
	fmt.Println(" line 142")
	defer queue.Stop()
	
	// for i := 0; i < 4*runtime.NumCPU()*2; i++ {
	for i := 0; i < 20; i++ {
		fmt.Printf("Submitting job %d\n", i)
		queue.Submit(&TestJob{strconv.Itoa(i)})
	}
	fmt.Println(" Completed Q submissions")
}
