package max_heap

import (
	"math"
)

type HeapNode struct {
	priority int
	aux      string
}

func NewHeapNode(priority int, aux string) *HeapNode {
	return &HeapNode{
		priority: priority,
		aux:      aux,
	}
}

type MaxHeap struct {
	data []*HeapNode
}

func (heap *MaxHeap) GetLength() int {
	return len(heap.data) - 1
}

func (*MaxHeap) getParentIndex(index int) int {
	return int(math.Floor(float64(index) / 2))
}

func (*MaxHeap) getLeftChildIndex(index int) int {
	return index * 2
}

func (*MaxHeap) getRightChildIndex(index int) int {
	return index*2 + 1
}

func (heap *MaxHeap) Insert(value HeapNode) {
	if len(heap.data) == 0 {
		heap.data = append(heap.data, nil)
	}

	heap.data = append(heap.data, &value)

	currentIndex := len(heap.data) - 1
	parentIndex := heap.getParentIndex(currentIndex)

	for parentIndex > 0 && heap.data[parentIndex].priority < heap.data[currentIndex].priority {
		heap.data[currentIndex], heap.data[parentIndex] = heap.data[parentIndex], heap.data[currentIndex]
		currentIndex, parentIndex = parentIndex, heap.getParentIndex(parentIndex)
	}
}

func (heap *MaxHeap) RemoveMax() *HeapNode {
	if heap.GetLength() == 0 {
		return nil
	}

	heap.data[1], heap.data[len(heap.data)-1] = heap.data[len(heap.data)-1], heap.data[1]

	max := heap.data[len(heap.data)-1]
	heap.data = heap.data[:len(heap.data)-1]

	for i := 1; i < len(heap.data); i++ {

		leftChildIndex := heap.getLeftChildIndex(i)
		rightChildIndex := heap.getRightChildIndex(i)

		if leftChildIndex < len(heap.data) {
			if heap.data[leftChildIndex].priority > heap.data[i].priority {
				heap.data[leftChildIndex], heap.data[i] = heap.data[i], heap.data[leftChildIndex]
			}
		}

		if rightChildIndex < len(heap.data) {
			if heap.data[rightChildIndex].priority > heap.data[i].priority {
				heap.data[rightChildIndex], heap.data[i] = heap.data[i], heap.data[rightChildIndex]
			}
		}
	}

	return max
}
