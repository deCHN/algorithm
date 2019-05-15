// Package algorithm contains implementations of classic algorithm solutions.
package algorithm

// Add return the addition of two integers.
// This is not even a algorithm.
func Add(a, b int) int {
	return (a + b)
}

// Identifier identify itself
type Identifier interface {
	getID() string
}

// IDImpl implements Identifier
type IDImpl struct {
	ID string
}

func (id *IDImpl) getID() string {
	return id.ID
}
