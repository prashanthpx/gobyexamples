package main

import "fmt"

// Run with: go run types/004_untyped_iota.go

// Untyped const: adopts type at use site
const Pi = 3.14159

// Enums with iota

type Status int

const (
	Unknown Status = iota
	Ready
	Running
	Done
)

func (s Status) String() string {
	return [...]string{"Unknown","Ready","Running","Done"}[s]
}

func main() {
	var x = Pi // becomes typed float64 here
	fmt.Println(x, Ready, Running, Done)
}

