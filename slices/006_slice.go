package main
import "fmt"

func SubtractOneFromLength(slice []int) []int {
	slice = slice[0 : len(slice)-1]
	//Here we modified the length of passed slice which is a copy of the original slice
	// We cannot modify the slice header length
	fmt.Println("SubtractOneFromLength: slice = ", slice)
    return slice
}

func main() {
	buffer := [10]int{0, 1, 2, 3 ,4 ,5,6 ,7 ,8, 9}
	slice := buffer[3:8]
    fmt.Println("Before: len(slice) =", len(slice), " slice = ", slice)
    newSlice := SubtractOneFromLength(slice)
    fmt.Println("After:  len(slice) =", len(slice))
    fmt.Println("After:  len(newSlice) =", len(newSlice))
}