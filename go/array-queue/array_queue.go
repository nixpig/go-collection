package array_queue

import "fmt"

type ArrayQueue[T any] struct {
	queue []arrayQueueItem[T]
}

type arrayQueueItem[T any] struct {
	value T
}

func (q *ArrayQueue[T]) Enqueue(item T) {
	q.queue = append(q.queue, arrayQueueItem[T]{value: item})
}

func (q *ArrayQueue[T]) Dequeue() (T, error) {
	if len(q.queue) > 0 {
		item := q.queue[0]
		q.queue = q.queue[1:]

		return item.value, nil
	}

	return arrayQueueItem[T]{}.value, fmt.Errorf("No items in the queue to dequeue")
}

// also build a circular queue (ring buffer)
