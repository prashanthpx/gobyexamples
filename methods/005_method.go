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

/*
Output
 user name: Prashanth Kumar
*/

/*
Code Explanation:
- Purpose: Demonstrate pointer receiver methods mutating or accessing via pointer
- User has FirstName/LastName; display uses a *User receiver
- main constructs a *User and calls the method
*/
