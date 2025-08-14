package main

import (
	"fmt"
	"sync"
)

// Run with: go run mutex/003_deadlock_lock_order.go

var muA, muB sync.Mutex

func AthenB() {
	muA.Lock(); defer muA.Unlock()
	muB.Lock(); defer muB.Unlock()
	fmt.Println("A then B")
}

func BthenA() {
	muB.Lock(); defer muB.Unlock()
	muA.Lock(); defer muA.Unlock()
	fmt.Println("B then A")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func(){ defer wg.Done(); AthenB() }()
	go func(){ defer wg.Done(); BthenA() }()
	wg.Wait()
	fmt.Println("If you see this hang, it indicates potential deadlock; enforce lock ordering")
}

