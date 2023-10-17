package main

import (
	"fmt"

	"github.com/nixpig/datastructures-the-fun-way/go/max_heap"
)

func main() {
	heap := max_heap.MaxHeap[string]{}

	values := []max_heap.HeapNode[string]{
		*max_heap.NewHeapNode(23, "foo"),
		*max_heap.NewHeapNode(42, "bar"),
		*max_heap.NewHeapNode(69, "baz"),
		*max_heap.NewHeapNode(13, "qux"),
		*max_heap.NewHeapNode(7, "rot"),
		*max_heap.NewHeapNode(132, "apple"),
		*max_heap.NewHeapNode(666, "chimp"),
		*max_heap.NewHeapNode(500, "cup"),
	}

	for _, value := range values {
		heap.Insert(value)
	}

	for heap.GetLength() > 0 {
		max := heap.RemoveMax()
		if max != nil {
			fmt.Println("max:", *max)
		} else {
			fmt.Println("no items left in heap")
		}

	}
}
