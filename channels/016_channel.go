package main

import (
	"fmt"
	"time"
)

func main() {

	// make the request chan chan that both go-routines will be given
	requestChan := make(chan chan string)
	fmt.Println(" line 10 requestChan: %v", requestChan)
	// start the goroutines
	go goroutineC(requestChan)
	go goroutineD(requestChan)

	// sleep for a second to let the goroutines complete
	time.Sleep(time.Second)

}

func goroutineC(requestChan chan chan string) {

	// make a new response chan
	responseChan := make(chan string)
	fmt.Println(" line 24 responseChan: %v", responseChan)
	// send the responseChan to goRoutineD
	requestChan <- responseChan

	// read the response
	response := <-responseChan

	fmt.Printf("Response: %v\n", response)

}

func goroutineD(requestChan chan chan string) {

	// read the responseChan from the requestChan
	dresponseChan := <-requestChan
	fmt.Println(" line 39 dresponseChan: %v", dresponseChan)
	// send a value down the responseChan
	dresponseChan <- "wassup!"

}

/*
Output
 line 10 requestChan: %v 0x<addr>
 line 24 responseChan: %v 0x<addr>
 line 39 dresponseChan: %v 0x<addr>
Response: wassup!
*/

/*
Code Explanation:
- Purpose: Channel-of-channels request/response pattern between goroutines
- goroutineC creates a response channel and sends it to goroutineD via requestChan
- goroutineD receives it and replies; C prints the response
*/
