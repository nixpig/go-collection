package heap_test

import (
	"testing"

	heap "github.com/nixpig/datastructures-the-fun-way/go/heap"
)

func TestMaxHeap(t *testing.T) {
	type Task struct {
		priority     int
		taskName     string
		instructions string
		assignee     string
		completed    bool
	}

	items := []heap.HeapNode[Task]{
		*heap.NewHeapNode(Task{
			priority:     1,
			taskName:     "P1 1",
			instructions: "",
			assignee:     "James",
			completed:    false,
		}),
		*heap.NewHeapNode(Task{
			priority:     3,
			taskName:     "P3 1",
			instructions: "",
			assignee:     "James",
			completed:    false,
		}),
		*heap.NewHeapNode(Task{
			priority:     3,
			taskName:     "P3 2",
			instructions: "",
			assignee:     "James",
			completed:    false,
		}),
		*heap.NewHeapNode(Task{
			priority:     2,
			taskName:     "P2 1",
			instructions: "",
			assignee:     "James",
			completed:    false,
		}),
		*heap.NewHeapNode(Task{
			priority:     1,
			taskName:     "P1 2",
			instructions: "",
			assignee:     "James",
			completed:    false,
		}),
		*heap.NewHeapNode(Task{
			priority:     4,
			taskName:     "P4 1",
			instructions: "",
			assignee:     "James",
			completed:    false,
		}),
	}

	taskMaxHeap := heap.Heap[Task]{}
	taskMaxHeap.Sort = func(gt Task, lt Task) bool {
		return gt.priority > lt.priority
	}

	for _, item := range items {
		taskMaxHeap.Insert(item)
	}

	var v int

	v = taskMaxHeap.RemoveMax().GetValue().priority
	if v != 1 {
		t.Errorf("expected 1 but got %v", v)
	}

	v = taskMaxHeap.RemoveMax().GetValue().priority
	if v != 1 {
		t.Errorf("expected 1 but got %v", v)
	}

	v = taskMaxHeap.RemoveMax().GetValue().priority
	if v != 2 {
		t.Errorf("expected 2 but got %v", v)
	}

	v = taskMaxHeap.RemoveMax().GetValue().priority
	if v != 3 {
		t.Errorf("expected 3 but got %v", v)
	}

	v = taskMaxHeap.RemoveMax().GetValue().priority
	if v != 3 {
		t.Errorf("expected 3 but got %v", v)
	}

	v = taskMaxHeap.RemoveMax().GetValue().priority
	if v != 4 {
		t.Errorf("expected 4 but got %v", v)
	}
}
