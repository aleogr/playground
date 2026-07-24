package stack

import (
	"errors"
	"iter"
)

type LinkedStack[T any] struct {
	top  *Element[T]
	size int
}

type Element[T any] struct {
	val  T
	next *Element[T]
}

func (s *LinkedStack[T]) Push(val T) {
	e := &Element[T]{val: val, next: s.top}
	s.top = e
	s.size++
}

func (s *LinkedStack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return *new(T), errors.New("Stack is empty")
	}
	val := s.top.val
	s.top = s.top.next
	s.size--
	return val, nil
}

func (s *LinkedStack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		return *new(T), errors.New("Stack is empty")
	}
	return s.top.val, nil
}

func (s *LinkedStack[T]) All() iter.Seq[T] {
	return nil
}

func (s *LinkedStack[T]) Size() int {
	return s.size
}

func (s *LinkedStack[T]) IsEmpty() bool {
	return s.top == nil
}
