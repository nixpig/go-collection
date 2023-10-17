package max_heap

import (
	"math"
)

type HeapNode struct {
	priority int
}

type MaxHeap struct {
	Data []*HeapNode
}

func (heap *MaxHeap) GetLength() int {
	return len(heap.Data) - 1
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

func (heap *MaxHeap) Insert(value int) {
	if len(heap.Data) == 0 {
		heap.Data = append(heap.Data, nil)
	}

	heap.Data = append(heap.Data, &HeapNode{priority: value})

	currentIndex := len(heap.Data) - 1
	parentIndex := heap.getParentIndex(currentIndex)

	for parentIndex > 0 && heap.Data[parentIndex].priority < heap.Data[currentIndex].priority {
		heap.Data[currentIndex], heap.Data[parentIndex] = heap.Data[parentIndex], heap.Data[currentIndex]
		currentIndex, parentIndex = parentIndex, heap.getParentIndex(parentIndex)
	}
}

func (heap *MaxHeap) RemoveMax() *HeapNode {
	if heap.GetLength() == 0 {
		return nil
	}

	heap.Data[1], heap.Data[len(heap.Data)-1] = heap.Data[len(heap.Data)-1], heap.Data[1]

	max := heap.Data[len(heap.Data)-1]
	heap.Data = heap.Data[:len(heap.Data)-1]

	for i := 1; i < len(heap.Data); i++ {

		leftChildIndex := heap.getLeftChildIndex(i)
		rightChildIndex := heap.getRightChildIndex(i)

		if leftChildIndex < len(heap.Data) {
			if heap.Data[leftChildIndex].priority > heap.Data[i].priority {
				heap.Data[leftChildIndex], heap.Data[i] = heap.Data[i], heap.Data[leftChildIndex]
			}
		}

		if rightChildIndex < len(heap.Data) {
			if heap.Data[rightChildIndex].priority > heap.Data[i].priority {
				heap.Data[rightChildIndex], heap.Data[i] = heap.Data[i], heap.Data[rightChildIndex]
			}
		}
	}

	return max
}
