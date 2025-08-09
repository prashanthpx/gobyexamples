package main 
  
import (
	"fmt"
	"time"
)

func myFunc(ch chan int) {
	fmt.Printf("\n running myFunc")

	<-ch
	fmt.Printf("\n read channel in myFunc")
}

func main() { 
    fmt.Println("start Main method") 
	// Defining a channel
	ch := make(chan int)
	go myFunc(ch)
	time.Sleep(2 * time.Second)
	// Once written to a channel, current program execution gets blocked
	// One has to read it to unblock it
	ch<-23
	fmt.Println("exiting main\n") 
} 