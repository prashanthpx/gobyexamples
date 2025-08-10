package main

import (
	"fmt"
)

func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

func main() {
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}

/*
Output
[Go world playground]
[Go world]
*/

/*
Code Explanation:
- Purpose: Passing a slice into a variadic function modifies the backing array element but not caller’s slice length
- s[0] = "Go" updates the first element; appending to s doesn’t change caller’s length
- After change, welcome shows updated first element but original length
*/
