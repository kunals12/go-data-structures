/*
Generics enable you to write functions, types, or methods that can operate on multiple types without sacrificing type safety.

Think of it like a template — instead of duplicating code for int, string, float64, etc.,
you define one version that works with all those types.
*/

/*
*************Basic Usage**************
package main

import (
	"fmt"
)

// Define a type constraint for comparable types (can use >, <, ==)
type Number interface {
	~int | ~int64 | ~float64 // ~ allows type aliases
}

// Generic function to get max of two values
func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("Max Int:", Max(10, 5))      // T inferred as int
	fmt.Println("Max Float:", Max(2.3, 4.5)) // T inferred as float64
}
*/

// Think of a generic warehouse shelf that can store boxes, bottles, or books.
// You don't build a new shelf each time — the shelf adjusts to what you need, while ensuring it doesn't break.
package main

import "fmt"

// Generic stack using a type parameter T
type Stack[T any] struct {
	items []T
}

// Push a value onto the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop a value from the stack
func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		var zero T // default zero value
		return zero
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last
}

func main() {
	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	fmt.Println("Popped from int stack:", intStack.Pop())

	strStack := Stack[string]{}
	strStack.Push("hello")
	strStack.Push("world")
	fmt.Println("Popped from string stack:", strStack.Pop())
}
