package main

import "fmt"

func main() {
	elements := map[string]string {
		"drive": "ssd",
		"name": "pk",
	}
	for i, j := range elements {
		fmt.Println(i, j)
	}

}