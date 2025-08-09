package main
import "fmt"

func main() {
	slice := make([]int, 5, 5)
	//populating data
	for i:=0; i<5; i++ {
		slice[i] = i
	}

	newSlice := make([]int, 5, 5)
	fmt.Println(slice)
	// Here as the new slice has length upto 5, which matches with the existing slice, it copies all data
	copy(newSlice, slice)
	fmt.Println(newSlice)
	/*
	The copy function is smart. It only copies what it can, paying attention to the lengths of both arguments. 
	In other words, the number of elements it copies is the minimum of the lengths of the two slices. 
	This can save a little bookkeeping. 
	Also, copy returns an integer value, the number of elements it copied, although it's not always worth checking.
	*/

	smallSlice := make([]int, 3, 5)
	copy(smallSlice, slice)
	fmt.Println(smallSlice)
}