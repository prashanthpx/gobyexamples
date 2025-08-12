package main

import "fmt"

// Run with: go run functions/mistakes/loop_var_capture.go
// Demonstrates loop variable capture bug and the fix.

func bad() {
	fmt.Println("bad:")
	var prints []func()
	for i := 0; i < 3; i++ {
		prints = append(prints, func() { fmt.Print(i, " ") })
	}
	for _, f := range prints { f() }
	fmt.Println()
}

func good() {
	fmt.Println("good:")
	var prints []func()
	for i := 0; i < 3; i++ {
		i := i // capture value per-iteration
		prints = append(prints, func() { fmt.Print(i, " ") })
	}
	for _, f := range prints { f() }
	fmt.Println()
}

func main() { bad(); good() }

