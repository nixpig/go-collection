package linked_list_queue

import "fmt"

type LinkedListNode[T comparable] struct {
	value T
	next  *LinkedListNode[T]
}

type LinkedListQueue[T comparable] struct {
	head   *LinkedListNode[T]
	tail   *LinkedListNode[T]
	length int
}

func NewLinkedListQueue[T comparable]() *LinkedListQueue[T] {
	return &LinkedListQueue[T]{}
}

func (l *LinkedListQueue[T]) Enqueue(value T) {
	newNode := &LinkedListNode[T]{
		value: value,
	}

	if l.length == 0 {
		l.tail = newNode
		l.head = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}

	l.length++
}

func (l *LinkedListQueue[T]) Dequeue() (T, error) {
	emptyNode := LinkedListNode[T]{}

	if l.length == 0 {
		return emptyNode.value, fmt.Errorf("Queue is empty; nothing to dequeue.")
	}

	dequeuedValue := l.head.value

	l.head = l.head.next

	l.length--

	return dequeuedValue, nil
}
