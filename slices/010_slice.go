package main
import "fmt"

func main() {
	arr := [5] int {0, 1, 2, 3 ,4}
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