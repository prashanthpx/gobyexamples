package main

import (
	"fmt"
)

func main() {
	sl := make([]int, 3)
	fmt.Println("len %d, cap %d", len(sl), cap(sl))
	sl[0] = 1
	//sl2 := []int{10}
	sl = append(sl, 20)
	fmt.Println("len %d, cap %d", len(sl), cap(sl))

	sl = append(sl, 30)
	fmt.Println("len %d, cap %d", len(sl), cap(sl))

	sl = append(sl, 40)
	fmt.Println("len %d, cap %d", len(sl), cap(sl))

	fmt.Println("len %d, cap %d", len(sl), cap(sl))
	//sl[2] = 400
	sl = append(sl, 50)
	fmt.Println("len %d, cap %d", len(sl), cap(sl))

	fmt.Printf("sl : %v", sl)
}