/*
Pogram to demonstrate passing no parameter to receiver just calling
the method
*/
package main
import "fmt"

type Val int

func (v Val) display() {
	fmt.Println("In display()")
}

func main() {
	var num Val
	fmt.Printf(" calling method ")
	num.display()
}