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