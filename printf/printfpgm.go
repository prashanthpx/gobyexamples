package main

import "fmt"

func main() {
	x := 100
	pi := 3.1412435
	isbool := true
	str := "golang"

	// String concatination while declaring
	concatstr := "go" + "lang"

	fmt.Printf(" x = %d", x)
	fmt.Printf(" pi =  %f", pi)

	// Priting needed decimal spaces
	fmt.Printf(" pi = %.3f", pi)

	//Printing boolen using %t
	fmt.Printf(" isbool = %t ", isbool)

	fmt.Println(len(str))

	// String concatination
	fmt.Println(str + "language")

	fmt.Printf(" concatstr : %s", concatstr)

	// Printing binary
	fmt.Printf(" binary of x = %b\n", x)

	// char printing
	fmt.Printf(" keycode %c\n", 65)

	//Printing hexa
	fmt.Printf(" hex of x = %x\n", x)

	// Printing in scientific notation
	fmt.Printf(" scientific pi = %e\n", pi)
}