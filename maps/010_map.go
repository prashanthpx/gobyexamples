package main

import (
	"fmt"
)

func main() {
	mp := make(map[string]int)
	mp["one"] = 1
	mp["two"] = 2
	modify(mp)
	fmt.Printf("mp : %v", mp)
	dummy := map[string]bool{}
	dummy["one"] = true
	if !dummy["two"] {
		fmt.Printf("line 15")
	} else {
		fmt.Printf("line 17")
	}
}

func modify(mp1 map[string]int) {
	mp1["three"] = 3
	mp1["one"] = 10
}

/*
Output
mp : map[one:10 three:3 two:2]line 15
*/

/*
Code Explanation:
- Purpose: Show that maps are reference types and basic existence checks
- modify modifies the passed map (adds key "three", updates "one")
- dummy := map[string]bool{}; dummy["two"] is false (absent -> zero value)
- The if prints "line 15" because !dummy["two"] is true
*/
