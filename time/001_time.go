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
