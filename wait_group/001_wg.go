package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	//wg.Wait()
	fmt.Println("All go routines finished executing")
}

/*
Output
All go routines finished executing
*/

/*
Code Explanation:
- Purpose: Demonstrate WaitGroup usage (but wg.Wait() is commented out)
- Three goroutines are launched, each sleeps 2s and calls Done
- Since wg.Wait() is commented, main exits immediately after printing the message; goroutines may not complete
- To wait for all goroutines, uncomment wg.Wait()
*/
