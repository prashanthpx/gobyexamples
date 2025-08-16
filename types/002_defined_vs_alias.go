package main

import "fmt"

// Run with: go run types/002_defined_vs_alias.go

type UserID int           // defined type (distinct)
type MyInt = int          // alias (exact same type)

type Email string         // defined type

func (e Email) Domain() string {
	for i := range e {
		if e[i] == '@' { return string(e[i+1:]) }
	}
	return ""
}

func main() {
	var id UserID = 10
	var x int = int(id) // explicit conversion required
	fmt.Println(x, Email("a@b.com").Domain())
}

