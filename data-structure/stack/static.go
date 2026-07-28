package stack

import (
	"iter"
)

type StaticStack[T any] struct {
	stack []T
}

func (s *StaticStack[T]) Push(val T) {
	s.stack = append(s.stack, val)
}

func (s *StaticStack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return *new(T), ErrEmpty
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
		return *new(T), ErrEmpty
	}
	return s.stack[s.Size()-1], nil
}

func (s *StaticStack[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := len(s.stack) - 1; i >= 0; i-- {
			if !yield(s.stack[i]) {
				return
			}
		}
	}
}

func (s *StaticStack[T]) Size() int {
	return len(s.stack)
}

func (s *StaticStack[T]) IsEmpty() bool {
	return s.Size() == 0
}
