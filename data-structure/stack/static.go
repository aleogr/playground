package stack

import (
	"errors"
	"iter"
	"slices"
)

type StaticStack[T any] struct {
	stack []T
}

func (s *StaticStack[T]) Push(val T) {
	s.stack = append(s.stack, val)
}

func (s *StaticStack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return *new(T), errors.New("Stack is empty")
	}
	n := len(s.stack) - 1
	val := s.stack[n]
	var zero T
	s.stack[n] = zero
	s.stack = s.stack[:n]
	return val, nil
}

func (s *StaticStack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		return *new(T), errors.New("Stack is empty")
	}
	return s.stack[s.Size()-1], nil
}

func (s *StaticStack[T]) All() iter.Seq[T] {
	stack := slices.Clone(s.stack)
	slices.Reverse(stack)
	return slices.Values(stack)
}

func (s *StaticStack[T]) Size() int {
	return len(s.stack)
}

func (s *StaticStack[T]) IsEmpty() bool {
	return s.Size() == 0
}
