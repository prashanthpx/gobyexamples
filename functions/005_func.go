package main
import "fmt"

func main() {
	slice := []int{10, 20, 30}

	// when we want to pass a slice to function, we inclide
	// ... after the slice
	print(slice...)
	// printing the modified slice
	fmt.Println(slice)

}

func print(para ...int) {
	fmt.Println(" Inside func print")
	fmt.Println(para)
	// modiyfing the slice
	para[0] = 100
}