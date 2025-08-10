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

/*
Output
len %d, cap %d 3 3
len %d, cap %d 4 6
len %d, cap %d 5 6
len %d, cap %d 6 6
len %d, cap %d 6 6
len %d, cap %d 7 12
sl : [1 0 0 20 30 40 50]
*/

/*
Code Explanation:
- Purpose: Show how append grows capacity over time
- Starts with len=3, cap=3; after appends, capacity grows (implementation-defined growth factor)
- Final slice contents demonstrate appended values
*/
