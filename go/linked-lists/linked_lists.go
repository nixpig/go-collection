package linked_list

import "fmt"

// import "fmt"

type LinkedListNode[T any] struct {
	value T
	next  *LinkedListNode[T]
}

type LinkedList[T any] struct {
	head *LinkedListNode[T]
	tail *LinkedListNode[T]
}

func (lst *LinkedList[T]) InsertAtTail(newValue T) {
	if lst.tail == nil {
		lst.head = &LinkedListNode[T]{
			value: newValue,
		}

		lst.tail = lst.head

	} else {
		lst.tail.next = &LinkedListNode[T]{
			value: newValue,
		}

		lst.tail = lst.tail.next
	}
}

func (lst *LinkedList[T]) InsertAtHead(newValue T) {
	if lst.head == nil {
		lst.head = &LinkedListNode[T]{
			value: newValue,
		}
	} else {
		newHead := &LinkedListNode[T]{
			value: newValue,
			next:  lst.head,
		}

		lst.head = newHead
	}
}

func (lst *LinkedList[T]) InsertAtPosition(position int, value T) error {
	if position == 0 {
		lst.InsertAtHead(value)
		return nil
	}

	current := lst.head

	for i := 0; i < position-1; i++ {
		if current.next != nil {
			current = current.next
		} else {
			return fmt.Errorf("Value '%d' provided for 'position' exceeds the length of the list", position)
		}
	}

	newNode := LinkedListNode[T]{
		value: value,
		next:  current.next,
	}

	current.next = &newNode

	return nil
}

func (lst *LinkedList[T]) Lookup(position int) (T, error) {
	current := lst.head
	for i := 0; i < position; i++ {
		if current.next != nil {

			current = current.next
		} else {
			empty := &LinkedListNode[T]{}
			return empty.value, fmt.Errorf("Value '%d' provided for 'position' exceeds the length of the list", position)
		}
	}

	return current.value, nil
}
