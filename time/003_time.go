package main

import (
	"fmt"
	"math/rand"
	"time"
)

// shuffle shuffles the elements of a slice in-place using the Fisher-Yates algorithm.
func shuffle(slice []int) {
	sd := time.Now().Unix()
	rand.Seed(sd) // Seed the random number generator.
	fmt.Printf("seed value %v", sd)
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)                   // Generate a random index between 0 and i.
		slice[i], slice[j] = slice[j], slice[i] // Swap elements.
	}
}

func main() {
	slice := []int{10, 20, 30, 40, 50}

	// Shuffle the slice
	shuffle(slice)

	// Iterate over the shuffled slice
	for index, value := range slice {
		fmt.Printf(" Index: %d, Value: %d\n", index, value)
	}
}