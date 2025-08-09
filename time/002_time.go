package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	loc, _ := time.LoadLocation("Local")
	
	fmt.Printf("Local Time: %s\n", now.In(loc))
	t := time.Now()
    zone_name, _ := t.Zone()
    fmt.Printf("Zone name is: %s\n", zone_name)
}

