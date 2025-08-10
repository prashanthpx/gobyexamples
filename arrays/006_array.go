package main

import "fmt"

func main() {
	arr := [...]int{10, 20, 30}
	fmt.Println("org array arr content: ", arr)
	// Passing array to function
	arrayPass(arr)
	// Printing original array after function call
	fmt.Println("After func call, arr content:  ", arr)
}

func arrayPass(modarr [3]int) {
	fmt.Println("Passed array content :", modarr)
	// modify the content of array
	modarr[0] = 100
	// pritning modified array
	fmt.Println("modified Passed array content :", modarr)
}

/*
Output
org array arr content:  [10 20 30]
Passed array content : [10 20 30]
modified Passed array content : [100 20 30]
After func call, arr content:   [10 20 30]
*/
