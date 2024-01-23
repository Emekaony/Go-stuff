package datastructures

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	elements []T
}

// push adds a new element to the top of the stack
func (s *Stack[T]) Push(element T) {
	// simply just add an element to the slice
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) IsEmpty() bool {
	// implementation
	return len(s.elements) == 0
}

func (s *Stack[T]) Pop() (*T, error) {
	// check if the stack is empty
	if s.IsEmpty() {
		return nil, errors.New("Cannot pop from an empty stack")
	}
	popped := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return &popped, nil
}

func (s *Stack[T]) String() string {
	return fmt.Sprintf("Stack contains: %v", s.elements)
}
