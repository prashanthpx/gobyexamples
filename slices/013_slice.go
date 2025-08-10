package main

import "fmt"

type info struct {
	name string
	age  int
}

func main() {
	infos := []*info{
		{
			name: "pk",
			age:  41,
		},
		{
			name: "kumar",
			age:  40,
		},
		{
			name: "k8s",
			age:  10,
		},
	}

	for _, inf := range infos {
		if inf.name == "k8s" {
			inf.name = "mango"
			inf.age = 100
		}
	}
	for _, inf := range infos {
		fmt.Printf("infos: %+v", inf)
	}
}

/*
Output
infos: &{name:pk age:41}infos: &{name:kumar age:40}infos: &{name:mango age:100}
*/

/*
Code Explanation:
- Purpose: Demonstrate mutating structs via pointers inside a slice
- infos is a slice of *info; modifying inf inside range updates underlying objects
- Shows that changing one element’s fields is reflected in the final print
*/
