package queue

import (
	"errors"
	"iter"
)

var ErrEmpty = errors.New("queue: empty queue")

type Queue[T any] interface {
	Enqueue(val T)
	Dequeue() (T, error)
	Head() (T, error)
	All() iter.Seq[T]
	Size() int
}
