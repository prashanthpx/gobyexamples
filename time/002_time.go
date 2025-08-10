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

/*
Output (sample; time/location dependent)
Local Time: 2025-08-10 17:20:03.01176 +0530 IST
Zone name is: IST
*/

/*
Code Explanation:
- Purpose: Print local time and time zone name
- now.In(loc) converts to the Local location; t.Zone() returns zone name and offset
*/
