package main
import "fmt"

func main() {
	arr := [...]int{10, 20, 30}
	fmt.Println("org arra ", arr)
	// passing array to function
	arrayPass(arr)
	// Pritning original array after function call
	fmt.Println("After func call, arr content:  ", arr)
}

func arrayPass(modarr [3]int) {
	fmt.Println("Passed array content :", modarr)
	// modify the content of array
	modarr[0] = 100
	// pritning modified array
	fmt.Println("modified Passed array content :", modarr)
}