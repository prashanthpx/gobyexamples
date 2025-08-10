package main

import "fmt"

func main() {
	arr := [5]int{0, 1, 2, 3, 4}
	fmt.Println("len = ", len(arr), "cap = ", cap(arr))

	// Lets append an element
	slice := arr[0:2]
	fmt.Println("slice = ", slice, "len = ", len(slice), "cap = ", cap(slice))
	slice = append(slice, 5)
	/*
		When append() creates a new slice, it doesn't create a slice that's just one larger than the slice before.
		It actually creates a slice that is already a couple of elements larger than the previous one.
	*/
	fmt.Println("Appending one more element")
	fmt.Println("slice = ", slice, "len = ", len(slice), "cap = ", cap(slice))
	slice = append(slice, 50)

	fmt.Println("Appending one more element")
	fmt.Println("slice = ", slice, "len = ", len(slice), "cap = ", cap(slice))
	fmt.Println("arr = ", arr, "len = ", len(arr), "cap = ", cap(arr))

	// Lets append till we reach the capacity
	slice = append(slice, 100)
	fmt.Println("Appending on more element")
	fmt.Println("slice = ", slice, "len = ", len(slice), "cap = ", cap(slice))
	fmt.Println("arr = ", arr, "len = ", len(arr), "cap = ", cap(arr))

	slice = append(slice, 150)
	fmt.Println("Appending on more element")
	fmt.Println("slice = ", slice, "len = ", len(slice), "cap = ", cap(slice))
	fmt.Println("arr = ", arr, "len = ", len(arr), "cap = ", cap(arr))

	// Aft this point, we can see that append has exceeded the capacity and has created a new slice
	// This means any modification to new slice has not impact on the original array referece
	// now lets modifu first element of slice
	fmt.Println("Modifying the first element of the slice")
	slice[0] = 10
	fmt.Println("slice = ", slice, "len = ", len(slice), "cap = ", cap(slice))
	fmt.Println("arr = ", arr, "len = ", len(arr), "cap = ", cap(arr))

}

/*
Output
len =  5 cap =  5
slice =  [0 1] len =  2 cap =  5
Appending one more element
slice =  [0 1 5] len =  3 cap =  5
Appending one more element
slice =  [0 1 5 50] len =  4 cap =  5
arr =  [0 1 5 50 4] len =  5 cap =  5
Appending on more element
slice =  [0 1 5 50 100] len =  5 cap =  5
arr =  [0 1 5 50 100] len =  5 cap =  5
Appending on more element
slice =  [0 1 5 50 100 150] len =  6 cap =  10
arr =  [0 1 5 50 100] len =  5 cap =  5
Modifying the first element of the slice
slice =  [10 1 5 50 100 150] len =  6 cap =  10
arr =  [0 1 5 50 100] len =  5 cap =  5
*/

/*
Code Explanation:
- Purpose: Show how append can modify the backing array until capacity, then allocate a new one
- Initial appends mutate arr via shared backing; once capacity exceeded, a new backing array is allocated
- Subsequent changes to slice no longer affect arr
*/
