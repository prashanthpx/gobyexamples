package main
import "fmt"

func main() {
	a := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun", "1", "2", "3"}

	slice1 := a[0:]
	slice2 := a[3:]

	fmt.Println("------- Before Modifications -------")
	fmt.Println("a  = ", a)
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)

	//slice1[0] = "modTUE"
	//slice1[1] = "modWED"
	//slice1[2] = "modTHU"

	//slice2[1] = "FRIDAY"

	fmt.Println("\n-------- After Modifications --------")
	fmt.Println("a  = ", a)
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
	//for i := 0; i < len(slice1); i++ {
	//for i, val := range slice1 {
	for len(slice1) >= 1 {
		index := 0
		//if len(slice1) > 1 {
			//fmt.Printf("i = %v, val: %v ", i, val)
			copy(slice1[index:], slice1[index+1:])
			//fmt.Println("line 27 slice1 = ", slice1)
			slice1 = slice1[:len(slice1) - 1]
			fmt.Println("line 29 slice1 = ", slice1)
			fmt.Printf("line 30 len: %d", len(slice1))
			fmt.Printf(" ************ \n")
	}
}
