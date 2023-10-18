package max_heap

import (
	"math"
)

type HeapNode[T any] struct {
	value T
}

func NewHeapNode[T any](value T) *HeapNode[T] {
	return &HeapNode[T]{
		value: value,
	}
}

type Heap[T any] struct {
	data []*HeapNode[T]
	Sort func(a T, b T) bool
}

func (heap *Heap[T]) GetLength() int {
	return len(heap.data) - 1
}

func (*Heap[T]) getParentIndex(index int) int {
	return int(math.Floor(float64(index) / 2))
}

func (*Heap[T]) getLeftChildIndex(index int) int {
	return index * 2
}

func (*Heap[T]) getRightChildIndex(index int) int {
	return index*2 + 1
}

func (heap *Heap[T]) Insert(value HeapNode[T]) {
	if len(heap.data) == 0 {
		heap.data = append(heap.data, nil)
	}

	heap.data = append(heap.data, &value)

	currentIndex := len(heap.data) - 1
	parentIndex := heap.getParentIndex(currentIndex)

	for parentIndex > 0 && heap.Sort(heap.data[parentIndex].value, heap.data[currentIndex].value) {
		heap.data[currentIndex], heap.data[parentIndex] = heap.data[parentIndex], heap.data[currentIndex]
		currentIndex, parentIndex = parentIndex, heap.getParentIndex(parentIndex)
	}
}

func (heap *Heap[T]) RemoveMax() *HeapNode[T] {
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
			if heap.Sort(heap.data[i].value, heap.data[leftChildIndex].value) {
				heap.data[leftChildIndex], heap.data[i] = heap.data[i], heap.data[leftChildIndex]
			}
		}

		if rightChildIndex < len(heap.data) {
			if heap.Sort(heap.data[i].value, heap.data[rightChildIndex].value) {
				heap.data[rightChildIndex], heap.data[i] = heap.data[i], heap.data[rightChildIndex]
			}
		}
	}

	return max
}
