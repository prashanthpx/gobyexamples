package main

import (
    "fmt"
)

func main() {
    // Define the function type
    type MathOp func(int, int) int

    // Create the map
    ops := map[string]MathOp{
        "add": func(a, b int) int { return a + b },
        "sub": func(a, b int) int { return a - b },
        "mul": func(a, b int) int { return a * b },
    }

    // Call a function from the map
    fmt.Println("3 + 4 =", ops["add"](3, 4)) // Output: 3 + 4 = 7
    fmt.Println("10 - 2 =", ops["sub"](10, 2)) // Output: 10 - 2 = 8
}

/*
Output
3 + 4 = 7
10 - 2 = 8
*/

/*
Code Explanation:
- Purpose: Demonstrate a map with function values for a calculator
- MathOp is a type alias for func(int, int) int
- ops map[string]MathOp stores named operations as functions
- ops["add"](3, 4) calls the function stored under "add" with args 3 and 4
- Prints the result of each operation
*/
