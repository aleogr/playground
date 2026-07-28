package queue

import "iter"

type CircularQueue[T any] struct{}

func (q *CircularQueue[T]) Enqueue(val T) {

}

func (q *CircularQueue[T]) Dequeue() (T, error) {
	return *new(T), nil
}

func (q *CircularQueue[T]) Head() (T, error) {
	return *new(T), nil
}

func (q *CircularQueue[T]) All() iter.Seq[T] {
	return nil
}

func (q *CircularQueue[T]) Size() int {
	return 0
}
