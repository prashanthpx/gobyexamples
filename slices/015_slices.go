package main

import (
	"fmt"
	"math/rand"
	"time"
)

type data struct {
	name string
}

func main() {
	var orgObjs []*data
	var objs []*data
	now := time.Now().Unix()
	for i := 0; i < 10; i++ {
		
		s2 := rand.NewSource(now + int64(i))
		r2 := rand.New(s2)

		dt := &data{
			name: fmt.Sprintf("%v", r2.Intn(10000)),
		}
		objs = append(objs, dt)
		for _, val := range objs {
			fmt.Println("line 27 %v ", val.name)
		}
		
	}
	orgObjs = objs
	// objs = nil
	time.Sleep(2 * time.Second)
	for _, val := range orgObjs {
		fmt.Println("line 33 %v ", val.name)
	}
}