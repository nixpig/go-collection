package linked_list_queue_test

import (
	"testing"

	linked_list_queue "github.com/nixpig/go-collection/data-structures-algorithms/linked-list-queue"
)

func TestLinkedListQueue(t *testing.T) {
	llq := &linked_list_queue.LinkedListQueue[int]{}

	llq.Enqueue(42)
	llq.Enqueue(23)

	item, err := llq.Dequeue()
	if err != nil {
		t.Errorf("errored out: %v", err)
	} else {
		if item != 42 {
			t.Errorf("expected 42, but received %v", item)
		}
	}

	item, err = llq.Dequeue()
	if err != nil {
		t.Errorf("errored out: %v", err)
	} else {
		if item != 23 {
			t.Errorf("expected 23, but received %v", item)
		}
	}

	item, err = llq.Dequeue()
	if err != nil {
		// pass, since no items left in queue and should return error
	} else {
		t.Errorf("expected to error, since nothing in queue, instead received %v", item)
	}

	llq.Enqueue(69)

	item, err = llq.Dequeue()
	if err != nil {
		t.Errorf("errored out: %v", err)
	} else {
		if item != 69 {
			t.Errorf("expected 69, but received %v", item)
		}
	}

}
