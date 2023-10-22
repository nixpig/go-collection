package ring_buffer_test

import (
	"testing"

	ring_buffer "github.com/nixpig/go-collection/data-structures-algorithms/ring-buffer"
)

func dequeueAndCheckValue[T comparable](queue *ring_buffer.RingBuffer[T], expectedValue T, t *testing.T) {
	item, err := queue.Dequeue()
	if err != nil {
		t.Errorf("expected not to error, but received: %v", err)
	} else {
		if item != expectedValue {
			t.Errorf("expected dequeued value to be '%v', but received '%v'", expectedValue, item)
		}
	}
}

func TestRingBuffer(t *testing.T) {
	rb := ring_buffer.NewRingBuffer[int](4)

	rb.Enqueue(42)
	rb.Enqueue(23)

	dequeueAndCheckValue(rb, 42, t)

	rb.Enqueue(69)
	rb.Enqueue(666)
	rb.Enqueue(777)

	dequeueAndCheckValue(rb, 23, t)

	rb.Enqueue(145)
	// Buffer is full, this should not get added!!
	rb.Enqueue(999)
	rb.Enqueue(000)

	dequeueAndCheckValue(rb, 69, t)
	dequeueAndCheckValue(rb, 666, t)
	dequeueAndCheckValue(rb, 777, t)

	rb.Enqueue(343)
	rb.Enqueue(986)
	rb.Enqueue(254)

	dequeueAndCheckValue(rb, 145, t)
	dequeueAndCheckValue(rb, 343, t)
	dequeueAndCheckValue(rb, 986, t)
	dequeueAndCheckValue(rb, 254, t)

	// Queue should now be empty and return error if dequeue attempted
	val, err := rb.Dequeue()
	if err != nil {
	} else {
		t.Errorf("This should return an error, since the buffer is empty. Instead, got: '%v'", val)
	}
}
