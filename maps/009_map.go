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
		sl = append(sl, strconv.Itoa(count))
		count++
		fmt.Printf("\n sl: %v", sl)
		ft["one"] = sl
	}

	//ft["two"] = []string{"100"}
	count1 := 100
	for count1 < 110 {
		sl := ft["two"]
		sl = append(sl, strconv.Itoa(count1))
		count1++
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

/*
Output
 sl: [0]
 sl: [0 1]
 sl: [0 1 2]
 sl: [0 1 2 3]
 sl: [0 1 2 3 4]
 sl: [0 1 2 3 4 5]
 sl: [0 1 2 3 4 5 6]
 sl: [0 1 2 3 4 5 6 7]
 sl: [0 1 2 3 4 5 6 7 8]
 sl: [0 1 2 3 4 5 6 7 8 9]
 sl: [100]
 sl: [100 101]
 sl: [100 101 102]
 sl: [100 101 102 103]
 sl: [100 101 102 103 104]
 sl: [100 101 102 103 104 105]
 sl: [100 101 102 103 104 105 106]
 sl: [100 101 102 103 104 105 106 107]
 sl: [100 101 102 103 104 105 106 107 108]
 sl: [100 101 102 103 104 105 106 107 108 109]map: map[one:[0 1 2 3 4 5 6 7 8 9] two:[100 101 102 103 104 105 106 107 108 109]]
 k: two
 100 101 102 103 104 105 106 107 108 109 k: one
 0 1 2 3 4 5 6 7 8 9
*/

/*
Code Explanation:
- Purpose: Build a map[string][]string by appending to slices for each key
- ft is a package-level map; ft["one"] and ft["two"] default to nil slices, safe to append
- Two loops append stringified numbers to slices; map updated each iteration
- Prints intermediate slices, final map, and iterates over map to print contents
- Note: Map iteration order is unspecified; observed order may vary
*/
