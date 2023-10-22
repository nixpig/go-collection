package main

import (
	"fmt"

	heap "github.com/nixpig/go-collection/data-structures-algorithms/heap"
)

func main() {
	fmt.Println("Max heap of strings sorted alphabetically")
	stringHeap := heap.Heap[string]{}
	stringHeap.Sort = func(a, b string) bool {
		return a > b
	}

	values := []heap.HeapNode[string]{
		*heap.NewHeapNode("foo"),
		*heap.NewHeapNode("bar"),
		*heap.NewHeapNode("baz"),
		*heap.NewHeapNode("qux"),
		*heap.NewHeapNode("rot"),
		*heap.NewHeapNode("apple"),
		*heap.NewHeapNode("chimp"),
		*heap.NewHeapNode("cup"),
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

	fmt.Println("Max heap of task structs sorted by descending priority")

	mapHeap := heap.Heap[Task]{}
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

	minHeap := heap.Heap[Task]{}
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

	fmt.Println("Heap sort an unsorted array")

	unsorted := []int{23, 72, 17, 13, 69, 42, 100}

	sorted := heap.HeapSort(unsorted, func(a, b int) bool { return a < b })

	fmt.Println(sorted)
}
