package array_queue_test

import (
	"testing"

	array_queue "github.com/nixpig/datastructures-the-fun-way/go/array-queue"
)

func dequeueAndCheckValue[T comparable](queue *array_queue.ArrayQueue[T], expectedValue T, t *testing.T) {
	item, err := queue.Dequeue()
	if err != nil {
		t.Errorf("expected not to error, but received: %v", err)
	} else {
		if item != expectedValue {
			t.Errorf("expected dequeued value to be '%v', but received '%v'", expectedValue, item)
		}
	}

}

func TestArrayQueue(t *testing.T) {

	aq := array_queue.ArrayQueue[int]{}

	aq.Enqueue(23)
	aq.Enqueue(42)

	dequeueAndCheckValue[int](&aq, 23, t)
	dequeueAndCheckValue[int](&aq, 42, t)

	_, err := aq.Dequeue()
	if err == nil {
		t.Errorf("Expected to receive an error, but didn't")
	}

	aq.Enqueue(69)
	dequeueAndCheckValue[int](&aq, 69, t)

}
