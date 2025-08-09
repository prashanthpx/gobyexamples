package main

import "fmt"

func main() {
	var age map[string]int
	if age == nil {
		fmt.Println("map is nil....")
		age = make(map[string]int)
		fmt.Println(age)
	}

	// adding items to map
	age["prashanth"] = 39
	age["kumar"] = 25
	age["krish"] = 30

	// printing map contents
	fmt.Println(age)
}