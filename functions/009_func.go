package main

import (
	"fmt"
)

func multiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

func main() {
	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println(double(5)) // 10
	fmt.Println(triple(5)) // 15
}

/*
Output
10
15
*/

/*
Code Explanation:
- Purpose: Demonstrate closures - functions that capture and remember variables from their outer scope
- multiplier(factor int) is a higher-order function that returns another function
- The returned anonymous function func(n int) int captures the 'factor' variable from its enclosing scope
- Each call to multiplier creates a new closure with its own copy of the factor value
- double := multiplier(2) creates a closure where factor=2 is permanently captured
- triple := multiplier(3) creates a separate closure where factor=3 is permanently captured
- When double(5) is called, it uses the captured factor=2: 5 * 2 = 10
- When triple(5) is called, it uses the captured factor=3: 5 * 3 = 15
- This demonstrates how closures maintain state between function calls
- Each closure is independent - they don't share the factor variable
- Closures are useful for creating specialized functions, callbacks, and maintaining private state

But since the inner function references x, Go promotes x to the heap, and the closure keeps a pointer to it.
*/
