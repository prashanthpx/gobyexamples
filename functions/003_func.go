//variadic functions
package main
import "fmt"

func main() {
	printPara(1, 10,20, 30)
	fmt.Println("Printing next para list...")
	printPara(2, 100,200, 300, 400, 500)
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
