package main

// Example showing channel buffering in go
func main() {
	chBuf := make(chan int, 2)
	// Even without a receiver we an send data to channel
	// and receiver it at later point

	// Here the sender blocks only when the buffer is full
	chBuf <- 1
	chBuf <- 2
	// Enabling the below statement makes the current context block
	//chBuf <- 3
}

/*
Output

*/

/*
Code Explanation:
- Purpose: Demonstrate buffered channel send without receiver
- chBuf has capacity 2; two sends succeed without blocking; a third would block
*/
