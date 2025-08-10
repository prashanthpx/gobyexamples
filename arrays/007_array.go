package main

import (
	"fmt"
)

func main() {
	arr := [3]int{10, 20 ,30}
	fmt.Println("Before mod arr content: ", arr)
	modArray(&arr)
	fmt.Println("Modified arr content: ", arr)
}

func modArray(marr *[3]int) {
	marr[0] = 100
	fmt.Println("modArray: ", *marr)
}

/*
Output
Before mod arr content:  [10 20 30]
modArray:  [100 20 30]
Modified arr content:  [100 20 30]
*/