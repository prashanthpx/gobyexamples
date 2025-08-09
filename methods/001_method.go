/*
Program to demonstrate usage of methods
*/

package main

import "fmt"

type User struct {
	Firstname, LastName string
}

func (u User) Display() string {
	return fmt.Sprintf("hi %s %s", u.Firstname, u.LastName)
}
func main() {
	u := User{"Prashanth", "Kumar"}
	// invoking method display using receiver
	fmt.Println(u.Display())
}