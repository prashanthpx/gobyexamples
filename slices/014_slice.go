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
	now := time.Now().Unix()
	for i := 0; i < 10; i++ {
		var objs []*data
		s2 := rand.NewSource(now + int64(i))
		r2 := rand.New(s2)

		dt := &data{
			name: fmt.Sprintf("%v", r2.Intn(10000)),
		}
		objs = append(objs, dt)
		for _, val := range objs {
			fmt.Printf("%v ", val.name)
		}
		time.Sleep(2 * time.Second)
	}
}