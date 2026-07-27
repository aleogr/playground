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
	e := s.stack[s.Size()-1]
	s.stack = s.stack[0 : s.Size()-1]
	return e, nil
}

func (s *StaticStack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		return *new(T), errors.New("Stack is empty")
	}
	return s.stack[s.Size()-1], nil
}

func (s *StaticStack[T]) All() iter.Seq[T] {
	return slices.Values(s.stack)
}

func (s *StaticStack[T]) Size() int {
	return len(s.stack)
}

func (s *StaticStack[T]) IsEmpty() bool {
	return s.Size() == 0
}
