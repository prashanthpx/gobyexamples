package main

import (
	"fmt"
)

type test struct {
	name string
	age  int
}

func main() {
	arr := []test{
		{"pk", 40},
		{"kum", 50},
		{"blr", 60},
		{"mang", 70},
	}
	// fill some data

	fmt.Printf("%v", arr[1])
	resp := make([][]interface{}, len(arr))

	for i, val := range arr {
		resp[i] = []interface{}{
			val.name,
			val.age,
		}
		if val.name == "mang" {
			resp[i] = append(resp[i], "kar")
		}
	}
	fmt.Printf("len of resp: %v", len(resp))
	fmt.Printf("content of resp: %v", (resp))
}

/*
Output
{kum 50}len of resp: 4content of resp: [[pk 40] [kum 50] [blr 60] [mang 70 kar]]
*/

/*
Code Explanation:
- Purpose: Build a [][]interface{} from a slice of structs with conditional extra field
- For each test, create a slice with name and age; add "kar" when name=="mang"
- Prints an element, then the length and complete data structure
*/
