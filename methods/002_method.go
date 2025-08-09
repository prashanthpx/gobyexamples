/*
Pogram to demonstrate multiple methods for a type
*/

package main
import "fmt"

type User struct {
	Firstname, Lastname string
}

func (u *User) display() string {
	return fmt.Sprintf(" In display, hi %s %s", u.Firstname, u.Lastname)
}

func (u *User) greet() string {
	return fmt.Sprintf(" in greet, Hello %s %s", u.Firstname, u.Lastname)

}

func main() {
	u := User{"prashanth", "kumar"}
	fmt.Printf(" Dis : %s", u.display())

	// re-assign a different vakue to 'u'
	u = User{"Green", "World"}
	fmt.Printf("\n Greet: %s ", u.greet())
}

