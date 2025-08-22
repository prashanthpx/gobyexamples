package main

import (
	"fmt"
)

func main() {
	str := []string{"h", "e", "l", "l", "o"}
	fmt.Println(str)
	str[0] = "H"
	fmt.Println(str)

	cStr := ""
	// Address
	fmt.Printf("Address of cStr is %p\n", &cStr)  
	cStr = cStr + "E" + "A"
	fmt.Printf("Address of cStr is %p\n", &cStr)
	fmt.Println(cStr)
	// Address

	var p *int32
	p = new(int32)
	fmt.Printf("Address of p is %p\n", p)

	var z *int32 = new(int32)
	fmt.Printf("Address of z is %p\n", z)

}