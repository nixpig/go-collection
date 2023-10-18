package main

import (
	"fmt"

	"github.com/nixpig/datastructures-the-fun-way/go/max_heap"
)

func main() {
	fmt.Println("Max heap of strings sorted alphabetically")
	stringHeap := max_heap.MaxHeap[string]{}
	stringHeap.Sort = func(a, b string) bool {
		return a > b
	}

	values := []max_heap.HeapNode[string]{
		*max_heap.NewHeapNode("foo"),
		*max_heap.NewHeapNode("bar"),
		*max_heap.NewHeapNode("baz"),
		*max_heap.NewHeapNode("qux"),
		*max_heap.NewHeapNode("rot"),
		*max_heap.NewHeapNode("apple"),
		*max_heap.NewHeapNode("chimp"),
		*max_heap.NewHeapNode("cup"),
	}

	for _, value := range values {
		stringHeap.Insert(value)
	}

	for stringHeap.GetLength() > 0 {
		max := stringHeap.RemoveMax()
		if max != nil {
			fmt.Println("max:", *max)
		} else {
			fmt.Println("no items left in heap")
		}

	}

	fmt.Println("---")

	type Task struct {
		priority     int
		taskName     string
		instructions string
		assignee     string
		completed    bool
	}

	items := []max_heap.HeapNode[Task]{
		*max_heap.NewHeapNode(Task{
			priority:     1,
			taskName:     "P1 1",
			instructions: "",
			assignee:     "James",
			completed:    false,
		}),
		*max_heap.NewHeapNode(Task{
			priority:     3,
			taskName:     "P3 1",
			instructions: "",
			assignee:     "James",
			completed:    false,
		}),
		*max_heap.NewHeapNode(Task{
			priority:     3,
			taskName:     "P3 2",
			instructions: "",
			assignee:     "James",
			completed:    false,
		}),
		*max_heap.NewHeapNode(Task{
			priority:     2,
			taskName:     "P2 1",
			instructions: "",
			assignee:     "James",
			completed:    false,
		}),
		*max_heap.NewHeapNode(Task{
			priority:     1,
			taskName:     "P1 2",
			instructions: "",
			assignee:     "James",
			completed:    false,
		}),
		*max_heap.NewHeapNode(Task{
			priority:     4,
			taskName:     "P4 1",
			instructions: "",
			assignee:     "James",
			completed:    false,
		}),
	}

	fmt.Println("Max heap of task structs sorted by descending priority")

	mapHeap := max_heap.MaxHeap[Task]{}
	mapHeap.Sort = func(gt Task, lt Task) bool {
		return gt.priority > lt.priority
	}

	for _, item := range items {
		mapHeap.Insert(item)
	}

	for mapHeap.GetLength() > 0 {
		got := mapHeap.RemoveMax()
		if got != nil {
			fmt.Println("got:", got)
		} else {
			fmt.Println("no items left")
		}
	}

	fmt.Println("Min heap of tasks structs sorted by ascending priority")

	minHeap := max_heap.MaxHeap[Task]{}
	minHeap.Sort = func(gt Task, lt Task) bool {
		return gt.priority < lt.priority
	}

	for _, item := range items {
		minHeap.Insert(item)
	}

	for minHeap.GetLength() > 0 {
		got := minHeap.RemoveMax()
		if got != nil {
			fmt.Println("got:", got)
		} else {
			fmt.Println("no items left")
		}
	}
}
