package main

import "fmt"

// Run with: go run types/005_methods_on_types.go

type Celsius float64

func (c Celsius) String() string { return fmt.Sprintf("%.1f°C", c) }

func (c *Celsius) Add(d Celsius) { *c += d }

func main() {
	t := Celsius(24)
	fmt.Println(t)
	t.Add(1)
	fmt.Println(t)
}

