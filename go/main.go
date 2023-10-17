package main

import (
	"fmt"

	"github.com/nixpig/datastructures-the-fun-way/go/max_heap"
)

func main() {
	heap := max_heap.MaxHeap{}

	values := []int{23, 42, 69, 13, 7, 132, 666, 500}

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
