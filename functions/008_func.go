package main

import (
	"fmt"
)

func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    next := counter()
    fmt.Println(next()) // 1
    fmt.Println(next()) // 2
    fmt.Println(next()) // 3
}

/*
Output
1
2
3
*/

/*
What is a Closure Again?
A closure is a function that captures variables from its surrounding lexical scope, allowing it to “remember” 
those variables even after the outer function has returned.
The key idea: closures bind to variables, not their values.
So if the original variable changes, the closure reflects that change.

Deep Dive: How Closures Remember Variables
When a closure is created, Go's compiler generates a new function object that includes:

The code of the function itself

A reference to any variables in the outer scope that it uses

This is known as a closure environment, and Go stores it alongside the function.

Even though those variables would normally be “out of scope,” Go keeps them alive on the heap if the closure needs them.
Closures bind to the variable, not a snapshot of its value.
*/