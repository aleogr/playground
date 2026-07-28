package queue

import "iter"

type ClassicQueue[T any] struct{}

func (q *ClassicQueue[T]) Enqueue(val T) {

}

func (q *ClassicQueue[T]) Dequeue() (T, error) {
	return *new(T), nil
}

func (q *ClassicQueue[T]) Head() (T, error) {
	return *new(T), nil
}

func (q *ClassicQueue[T]) All() iter.Seq[T] {
	return nil
}

func (q *ClassicQueue[T]) Size() int {
	return 0
}
