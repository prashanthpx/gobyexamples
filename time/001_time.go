package main

import (
	"fmt"
	"time"
)

func main() {
	tNow := time.Now().UTC()
	fmt.Printf("tNow: %v", tNow)
	location, err := time.LoadLocation("Asia/Calcutta")
	if err != nil {
		fmt.Printf("line 359 err: %v", err)
		//	return time.Time{}, nil, err
	}
	//logrus.Debugf("line 362 loc", location)
	tNow = tNow.In(location)
	startDay := tNow.Day()
	startTime := time.Date(tNow.Year(), tNow.Month(), startDay, 0, 0, 0, 0, location)
	fmt.Printf("time: %v", startTime)
	//startTime = time.Date(tNow.Year(), tNow.Month(), startDay + 40, 0, 0, 0, 0, location)
	//fmt.Printf("\n line 21 time: %v", startTime)
	//startTime = time.Date(tNow.Year(), tNow.Month(), startDay, 48, 0, 0, 0, location)
	//fmt.Printf("\n line 23 time: %v", startTime)
	startTime = time.Date(tNow.Year(), tNow.Month(), startDay, 0, 59, 59, 59, location)
	fmt.Printf("\n line 23 time: %v", startTime)
}

/*
Output (sample; time/location dependent)
tNow: 2025-08-10 11:48:50.471473 +0000 UTCtime: 2025-08-10 00:00:00 +0530 IST
 line 23 time: 2025-08-10 00:59:59.000000059 +0530 IST
*/

/*
Code Explanation:
- Purpose: Work with time zones and construct specific times in a location
- tNow := time.Now().UTC() then converted to Asia/Calcutta location
- Construct start of day and a specific minute/second in that location
*/
