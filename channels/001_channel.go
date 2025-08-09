package main

import (
	"fmt"
	"time"
)

func main() {
	uptimeTicker := time.NewTicker(3 * time.Second)
	//dateTicker := time.NewTicker(10 * time.Second)
	for {
		<-uptimeTicker.C
		run("uptime")
		<-uptimeTicker.C
		fmt.Printf("\n moved out now")
	}

	/*
	for {
		select {
		case <-uptimeTicker.C:
			run("uptime")
		case <-dateTicker.C:
			run("date")
		}
	}*/
}

func run(s string) {
	fmt.Printf(" \n path %s", s)
}
