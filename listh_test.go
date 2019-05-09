package algorithm

import (
	"fmt"
	"sync"
	"testing"
)

func TestRandFunnel(t *testing.T) {

	input1 := []int{}
	for i := -10; i < 10; i++ {
		input1 = append(input1, i)
	}

	input2 := []int{1, 2, 4, 5}

	wg := &sync.WaitGroup{}
	wg.Add(2)

	t.Run("continues number slice", func(t *testing.T) {
		go verify(input1, wg)
	})
	t.Run("discrete number slice", func(t *testing.T) {
		go verify(input2, wg)
	})
	wg.Wait()
}

func verify(input []int, wg *sync.WaitGroup) bool {
	set := make(map[int]bool)
	// every number should appear and appear only once
	for e := range RandFunnel(input) {
		if _, ok := set[e]; ok {
			fmt.Printf("found repeated number <%d>", e)
			return false
		} else {
			set[e] = true
		}
	}
	fmt.Println()
	for _, e := range input {
		if _, ok := set[e]; !ok {
			fmt.Printf("missing number <%d>", e)
			return false
		}
	}
	wg.Done()
	return true
}
