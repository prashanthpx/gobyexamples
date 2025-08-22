package main

import (
	"fmt"
)

type engine struct {
	mileage int
	hp      int
}

func (e *engine) display() {
	fmt.Printf(" mileage: %d, hp: %d\n", e.mileage, e.hp)
}

func (e *engine) modify() {
	e.mileage = 10000
	e.hp = 1000
}

func (e engine) mod() {
	e.mileage = 11111
	e.hp = 11
}

func main() {
	var ice *engine = &engine{
		mileage: 100,
		hp:      100,
	}
	ice.display()
	ice.modify()
	fmt.Printf("After modification, value is changed !!! ")
	ice.display()

	var eletric engine = engine{
		mileage: 100,
		hp:      100,
	}
	eletric.display()
	eletric.mod()
	fmt.Printf("After modification, value doesn't change !!!")
	eletric.display()
}