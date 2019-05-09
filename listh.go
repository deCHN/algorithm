package algorithm

import (
	"math/rand"
	"time"
)

//RandFunnel return all the element of a number list in pyseudo random pick
func RandFunnel(src []int) chan int {

	rand.Seed(time.Now().UnixNano())

	rt := make(chan int)

	go func() {
		for i := len(src); i > 0; i-- {
			next := rand.Intn(i)
			rt <- src[next]
			if i == 1 {
				break
			}
			last := len(src) - 1
			src[next] = src[last] // switch with the last one
			src = src[:last]      // chop off the last one
		}
		defer close(rt)
	}()

	return rt
}
