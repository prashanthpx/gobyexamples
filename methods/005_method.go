/*
Demonstrating pointer receiver methods
*/

package main

import "fmt"

type User struct {
	FirstName, LastName string
}

func (u *User) display() {
	fmt.Printf(" user name: %s %s", u.FirstName, u.LastName)
}

func main() {
	u := &User{"Prashanth", "Kumar"}
	u.display()
}