package main

import (
	"fmt"
)

type Person struct {
    Name string
}

func changeName(p *Person) {
    p.Name = "Alice"
}

func main() {
    person := Person{Name: "Bob"}
    changeName(&person)
    fmt.Println(person.Name) // Output: Alice (changed)
}


/*
Output
Alice
*/

/*
Code Explanation:
- Purpose: Show that Go passes struct pointers by copy, not by reference
*/
