// Package algorithm contains implementations of classic algorithm solutions.
package algorithm

// Add return the addition of n integers.
func Add(n... int) int {
   var sum int
   for _, i:= range n {
      sum += i
   }
   return sum
}
