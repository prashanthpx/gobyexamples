package main

import (
	"fmt"
	"time"
)

func main() {
	var ball = make(chan string)
	kickBall := func(playerName string) {
		for {
			fmt.Println(<-ball, "kicked the ball.")
			time.Sleep(time.Second)
			ball <- playerName
		}
	}
	go kickBall("John")
	go kickBall("Alice")
	go kickBall("Bob")
	go kickBall("Emily")
	ball <- "referee" // kick off
	var c chan bool   // nil
	<-c               // blocking here for ever
}

/*
Output (non-terminating; sample)
referee kicked the ball.
John kicked the ball.
Alice kicked the ball.
Bob kicked the ball.
Emily kicked the ball.
... repeats as players pass the ball ...
*/

/*
Code Explanation:
- Purpose: Demonstrate fan-in/out with a shared channel among goroutines
- Four players loop: receive from ball, print, sleep, send their name back
- main kicks off with "referee" and then blocks forever on a nil channel
*/
