package array_queue

import "fmt"

type ArrayQueue[T any] struct {
	queue []item[T]
}

type item[T any] any

func (q *ArrayQueue[T]) Enqueue(item T) {
	q.queue = append(q.queue, item)
}

func (q *ArrayQueue[T]) Dequeue() (item[T], error) {
	if len(q.queue) > 0 {
		value := q.queue[0]
		q.queue = q.queue[1:]

		return value, nil
	}

	return nil, fmt.Errorf("No items in the queue to dequeue")
}
