package main

import (
	"fmt"
)

func main() {
	mp := make(map[string]int)
	mp["one"] = 1
	mp["two"] = 2
	modify(mp)
	fmt.Printf("mp : %v",mp)
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