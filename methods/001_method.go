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

/*
Output
hi Prashanth Kumar
*/

/*
Code Explanation:
- Purpose: Demonstrate defining and invoking a method on a struct type
- type User has fields Firstname and LastName
- func (u User) Display() string is a value receiver method returning a greeting
- main constructs a User and calls u.Display()
*/
