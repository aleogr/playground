package queue

import "iter"

type DoubleEndedQueue[T any] struct{}

func (q *DoubleEndedQueue[T]) Enqueue(val T) {

}

func (q *DoubleEndedQueue[T]) Dequeue() (T, error) {
	return *new(T), nil
}

func (q *DoubleEndedQueue[T]) Head() (T, error) {
	return *new(T), nil
}

func (q *DoubleEndedQueue[T]) All() iter.Seq[T] {
	return nil
}

func (q *DoubleEndedQueue[T]) Size() int {
	return 0
}
