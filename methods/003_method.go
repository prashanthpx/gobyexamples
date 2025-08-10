package main

import "fmt"

type Value int

func (v Value) add() int {
	return int(v + 10)
}

func main() {
	//Passing value of 10
	val := Value(10)
	fmt.Printf(" calling add(): ret val = %d\n", val.add())

}

/*
Output
 calling add(): ret val = 20
*/

/*
Code Explanation:
- Purpose: Attach a method to a named type (Value) and use it
- type Value int defines a named type; (v Value) add() returns v+10
- main constructs Value(10) and prints the method result
*/
