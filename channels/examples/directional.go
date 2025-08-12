package main

import "fmt"

// Run with: go run channels/examples/directional.go

func producer(out chan<- int) { for i := 0; i < 3; i++ { out <- i }; close(out) }
func consumer(in <-chan int)  { for v := range in { fmt.Print(v, " ") } }

func main() {
	ch := make(chan int, 3)
	go producer(ch)
	consumer(ch)
}

