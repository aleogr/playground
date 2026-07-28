package stack

import (
	"errors"
	"iter"
)

var ErrEmpty = errors.New("stack: empty stack")

type Stack[T any] interface {
	Push(val T)
	Pop() (T, error)
	Peek() (T, error)
	All() iter.Seq[T]
	Size() int
}
