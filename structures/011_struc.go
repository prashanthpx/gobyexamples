package main

import (
	"fmt"
)

type Person struct {
    Name string
}

func changeName(p Person) {
    p.Name = "Alice"
}

func main() {
    person := Person{Name: "Bob"}
    changeName(person)
    fmt.Println(person.Name) // Output: Bob (unchanged)
}

/*
Output
Bob
*/

/*
Code Explanation:
- Purpose: Show that Go passes struct values by copy, not by reference
- Person has a Name field; changeName modifies the copy, not the original
- main prints the original person.Name, which is unchanged
*/
