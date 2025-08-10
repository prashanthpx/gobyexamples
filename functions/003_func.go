// variadic functions
package main

import "fmt"

func main() {
	printPara(1, 10, 20, 30)
	fmt.Println("Printing next para list...")
	printPara(2, 100, 200, 300, 400, 500)
	fmt.Println("Printing next para list...")
	printPara(1)
}

/*
varidaic functions parameter gets passed as slice.
Hence they can be accessed based on index
*/
func printPara(first int, nums ...int) {
	fmt.Printf("First para = %d\n", first)
	for i := range nums {
		fmt.Printf("Index = %d, content = %d\n ", i, nums[i])
	}
}

/*
Output
First para = 1
Index = 0, content = 10
 Index = 1, content = 20
 Index = 2, content = 30
 Printing next para list...
First para = 2
Index = 0, content = 100
 Index = 1, content = 200
 Index = 2, content = 300
 Index = 3, content = 400
 Index = 4, content = 500
 Printing next para list...
First para = 1
*/

/*
Code Explanation:
- Purpose: Variadic function parameters are received as a slice
- printPara(first int, nums ...int) iterates over nums and prints indices and values
- Demonstrates zero-argument varargs call (only first parameter)
*/
