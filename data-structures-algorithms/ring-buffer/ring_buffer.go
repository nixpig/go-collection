package ring_buffer

import "fmt"

type RingBuffer[T any] struct {
	queue []item[T]
	head  int
	tail  int
	size  int
	count int
}

type item[T any] any

func NewRingBuffer[T any](size int) *RingBuffer[T] {
	return &RingBuffer[T]{
		queue: make([]item[T], size),
		size:  size,
	}
}

func (rb *RingBuffer[T]) Enqueue(value T) (bool, error) {
	if rb.count == rb.size {
		return false, fmt.Errorf("Buffer is full!")
	}

	if rb.count == 0 {
		rb.queue[0] = value
		rb.tail = 0
	} else if rb.tail < rb.size-1 {
		rb.queue[rb.tail+1] = value
		rb.tail++
	} else {
		rb.queue[rb.size-rb.tail-1] = value
		rb.tail = (rb.tail - rb.size + 1)
	}

	rb.count++

	return true, nil
}

func (rb *RingBuffer[T]) Dequeue() (item[T], error) {
	if rb.count == 0 {
		return nil, fmt.Errorf("No items left")
	}

	value := rb.queue[rb.head]

	if rb.head < rb.size-1 {
		rb.head++
	} else {
		rb.head = 0
	}

	rb.count--

	return value, nil
}
