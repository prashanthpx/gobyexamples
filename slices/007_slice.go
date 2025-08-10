package main

import "fmt"

func main() {
	var iBuffer [10]int
	// Creating a slice of len = 0, but capacity is 10 as it's underlying array cap is 10
	slice := iBuffer[0:0]

	fmt.Println("len = ", len(slice), "cap =", cap(slice))
	for i := 0; i < 20; i++ {
		slice = Extend(slice, i)
		fmt.Println(slice)
	}
}

func Extend(slice []int, element int) []int {
	n := len(slice)
	// On every iteration, slice length is increased.
	// Once it cross capacity, it fails !!!
	slice = slice[0 : n+1]
	slice[n] = element
	fmt.Println("Extend: len = ", len(slice), "cap =", cap(slice))
	return slice
}

/*
Output (terminates with panic)
len =  0 cap = 10
Extend: len =  1 cap = 10
[0]
Extend: len =  2 cap = 10
[0 1]
Extend: len =  3 cap = 10
[0 1 2]
Extend: len =  4 cap = 10
[0 1 2 3]
Extend: len =  5 cap = 10
[0 1 2 3 4]
Extend: len =  6 cap = 10
[0 1 2 3 4 5]
Extend: len =  7 cap = 10
[0 1 2 3 4 5 6]
Extend: len =  8 cap = 10
[0 1 2 3 4 5 6 7]
Extend: len =  9 cap = 10
[0 1 2 3 4 5 6 7 8]
Extend: len =  10 cap = 10
[0 1 2 3 4 5 6 7 8 9]
panic: runtime error: slice bounds out of range [:11] with capacity 10
... stack trace omitted ...
*/

/*
Code Explanation:
- Purpose: Show manual slice growth and panic when exceeding capacity
- Extend grows the slice by 1 each time via reslicing; when n+1 > cap(slice), it panics
- Proper approach is to use append which grows capacity automatically
*/
