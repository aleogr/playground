package queue

import "iter"

type SimpleQueue[T any] struct{}

func (q *SimpleQueue[T]) Enqueue(val T) {

}

func (q *SimpleQueue[T]) Dequeue() (T, error) {
	return *new(T), nil
}

func (q *SimpleQueue[T]) Head() (T, error) {
	return *new(T), nil
}

func (q *SimpleQueue[T]) All() iter.Seq[T] {
	return nil
}

func (q *SimpleQueue[T]) Size() int {
	return 0
}
