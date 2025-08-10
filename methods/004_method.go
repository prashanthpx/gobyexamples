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

/*
Output
 calling method In display()
*/

/*
Code Explanation:
- Purpose: Method with no parameters on a basic named type
- type Val int, method (v Val) display() prints a message
- main calls the method on a zero-value Val
*/
