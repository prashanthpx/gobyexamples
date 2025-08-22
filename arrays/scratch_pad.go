package main

import (
	"fmt"
	"sync"
)

func main() {
	var num int = 10
	var fl float32 = 12.34
	_ = [3]int{0,1,2}
	// comb := num + fl // Won't work
	comb := float32(num) + fl

	var n int = 10
	fmt.Println(comb)

	wg := sync.WaitGroup{}

	go add()
	wg.Add(1)

	wg.Done()
	wg.Wait()

}


