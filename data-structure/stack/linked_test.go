package stack_test

import (
	"testing"

	"github.com/aleogr/playground/data-structure/stack"
)

func TestLinkedStack(t *testing.T) {
	s := stack.LinkedStack[int]{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
}
