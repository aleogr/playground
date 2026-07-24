package stack

import "iter"

type Stack[T any] interface {
	Push(val T)
	Pop() (T, error)
	Peek() (T, error)
	All() iter.Seq[T]
	Size() int
	IsEmpty() bool
}
