package main

import (
	"fmt"
)

func main() {
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialization and increment
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}
}

/*
Output:
- 10 * 1 = 10
- 11 * 2 = 22
- 12 * 3 = 36
- 13 * 4 = 52
- 14 * 5 = 70
- 15 * 6 = 90
- 16 * 7 = 112
- 17 * 8 = 136
- 18 * 9 = 162
- 19 * 10 = 190
*/

/*
Code Explanation:
- Purpose: Iterate two counters in lockstep and print their product each time
- Starting values: no = 10, i = 1
- Loop condition: continue while i <= 10 AND no <= 19
- Step each iteration: i, no = i+1, no+1 (both incremented together)
- Body: print "no * i = product"

Breakdown:
- Initialization: no, i := 10, 1
  - Declares and initializes both variables local to the loop scope
- Condition: i <= 10 && no <= 19
  - Loop runs as long as both conditions hold
- Post statement: i, no = i+1, no+1
  - Go evaluates the right-hand side first, then assigns both left-hand variables (safe multiple assignment)
- Print: fmt.Printf("%d * %d = %d\n", no, i, no*i)
  - %d formats integers; \n adds a newline
*/
