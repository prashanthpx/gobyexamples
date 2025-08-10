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

/*
Output (sample; values vary each run)
1350 9395 7829 8742 7720 4565 900 5761 211 7072
*/

/*
Code Explanation:
- Purpose: Generate pseudorandom names using a time-seeded RNG
- For i=0..9, create a new rand.Source with seed now+i to produce a value and print it
- Sleeps 2s between iterations; output values vary by time and iteration
*/
