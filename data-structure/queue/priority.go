package queue

import "iter"

type PriorityQueue[T any] struct{}

func (q *PriorityQueue[T]) Enqueue(val T) {

}

func (q *PriorityQueue[T]) Dequeue() (T, error) {
	return *new(T), nil
}

func (q *PriorityQueue[T]) Head() (T, error) {
	return *new(T), nil
}

func (q *PriorityQueue[T]) All() iter.Seq[T] {
	return nil
}

func (q *PriorityQueue[T]) Size() int {
	return 0
}
