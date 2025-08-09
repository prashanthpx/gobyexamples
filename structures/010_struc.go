package main

import "fmt"

type building struct {
	height, length int
	colour string
}

func main() {
	/*
	mall := building {
		height: 10,
		length: 20,
		colour: "blue",
	}

	newMall := mall
	fmt.Printf(" new mall : %v", newMall)
	newMall.height = 100
	fmt.Printf(" new mall : %v", mall)
	*/
	mall := &building{
		height: 10,
		length: 20,
		colour: "blue",
	}
	
	var newMall = &building{}
	//*newMall = *mall
	*newMall = *mall
	fmt.Printf("\n  address newMall: %p -- mall: %p", newMall, mall)
	fmt.Printf(" \n new mall : %v", newMall)
	newMall.height = 100
	fmt.Printf(" \n line 32 new mall : %v", newMall)
	fmt.Printf(" \n line 33 mall : %v", mall)
	mall.colour = "orange"
	fmt.Printf(" \n line 38 mall : %v", mall)
	fmt.Printf(" \n line 39 new mall : %v", newMall)



}