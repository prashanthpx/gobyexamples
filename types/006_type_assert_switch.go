package main

import "fmt"

// Run with: go run types/006_type_assert_switch.go

func describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	case fmt.Stringer:
		fmt.Println("stringer", v.String())
	default:
		fmt.Printf("unknown %T\n", v)
	}
}

func main() {
	describe(42)
	describe("x")
	describe(struct{}{})
}

