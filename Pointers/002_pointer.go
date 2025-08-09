package main
import (
	"fmt"
)

func main() {
	a := 100
	var b *int

	if b == nil {
		fmt.Println(" b is ", b)
		b = &a
		fmt.Println("b is now initialized to ", b)
	}
}