package main

import "fmt"

type building struct {
	height, length int
	colour         string
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

/*
Output

  address newMall: 0x140000ac080 -- mall: 0x140000ac060
 new mall : &{10 20 blue}
 line 32 new mall : &{100 20 blue}
 line 33 mall : &{10 20 blue}
 line 38 mall : &{10 20 orange}
 line 39 new mall : &{100 20 blue}
*/

/*
Code Explanation:
- Purpose: Demonstrate copying pointed-to structs vs separate pointers
- mall and newMall are separate pointers; *newMall = *mall copies underlying value
- Subsequent field changes on one pointer don't change the other’s fields
*/
