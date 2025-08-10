package main

import (
	"fmt"
)

type Car struct {
	license string
}

func (c *Car) Name() string {
	return "car"
}
func (c *Car) License() string {
	return c.license
}

type MotorBike struct {
	license string
}

func (mb *MotorBike) Name() string {
	return "motor bike"
}
func (mb *MotorBike) License() string {
	return mb.license
}

type Vehicle interface {
	License() string
	Name() string
}

func PrintLicense(v Vehicle) {
	fmt.Println("I've seen a " + v.Name() + " with the license plate " + v.License())
}

func main() {
	car := Car{"LJ178FU"}
	bike := MotorBike{"LK6IDVR"}

	PrintLicense(&car)
	PrintLicense(&bike)
}

/*
Output
I've seen a car with the license plate LJ178FU
I've seen a motor bike with the license plate LK6IDVR
*/

/*
Code Explanation:
- Purpose: Implement an interface with multiple concrete types and pass them to a function
- Vehicle interface requires Name and License; Car and MotorBike implement both
- PrintLicense accepts any Vehicle and calls its methods via the interface
*/
