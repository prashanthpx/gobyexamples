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

