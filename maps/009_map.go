package main

import (
	"fmt"
	"strconv"
)

var (
	ft = map[string][]string{}
)


func main() {
	//ft["one"] = []string{"1"}
	count := 0
	for count < 10 {
		sl := ft["one"]
		sl = append(sl,  strconv.Itoa(count))
		count++;
		fmt.Printf("\n sl: %v", sl)
		ft["one"] = sl
	}

	//ft["two"] = []string{"100"}
	count1 := 100
	for count1 < 110 {
		sl := ft["two"]
		sl = append(sl,  strconv.Itoa(count1))
		count1++;
		fmt.Printf("\n sl: %v", sl)
		ft["two"] = sl
	}
	fmt.Printf("map: %v\n", ft)

	for k, val := range ft {
		fmt.Printf(" k: %v\n", k)
		for _, sl := range val {
			fmt.Printf(" %v", sl)
		}
	}
}