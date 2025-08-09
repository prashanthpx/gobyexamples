/*
Program to demo annonymous structure
*/

package main

import "fmt"

func main() {
	//Here we declate annonymous stcuture and initialize the same
	building := struct {
		height, length int
		colour string
	} {
		colour: "white",
		height: 100,
		length: 500,
	}

	fmt.Println(" building : ", building)
}