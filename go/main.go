package main

import (
	"fmt"

	"github.com/nixpig/datastructures-the-fun-way/go/max_heap"
)

func main() {
	stringHeap := max_heap.MaxHeap[string]{}

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

	mapHeap := max_heap.MaxHeap[map[int]string]{}

	items := []max_heap.HeapNode[map[int]string]{
		*max_heap.NewHeapNode(14, map[int]string{23: "foo"}),
		*max_heap.NewHeapNode(34, map[int]string{42: "apple"}),
		*max_heap.NewHeapNode(73, map[int]string{12223: "banana"}),
		*max_heap.NewHeapNode(900, map[int]string{834: "dog"}),
		*max_heap.NewHeapNode(4, map[int]string{1: "cat"}),
		*max_heap.NewHeapNode(19, map[int]string{0: "tiger"}),
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
}
