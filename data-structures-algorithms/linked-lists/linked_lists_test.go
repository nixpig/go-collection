package linked_list

import (
	"fmt"
	"testing"
)

func lookupAndExpect[T comparable](lst LinkedList[T], position int, expectedValue T, t *testing.T) {
	lu, err := lst.Lookup(position)
	if err != nil {
		t.Errorf("expected not to error, but got: %v", err)
	} else if lu != expectedValue {
		t.Errorf("expected %v, but got %v", expectedValue, lu)
	}
}

func TestLinkedList(t *testing.T) {
	lst := LinkedList[int]{}

	fmt.Println("insert at head")
	lst.InsertAtHead(23)
	lst.InsertAtHead(42)
	lst.InsertAtHead(69)

	fmt.Println("lookup within bounds")

	lookupAndExpect[int](lst, 0, 69, t)
	lookupAndExpect[int](lst, 1, 42, t)
	lookupAndExpect[int](lst, 2, 23, t)

	fmt.Println("lookup out of bounds")
	lu, err := lst.Lookup(3)
	if err != nil {
	} else {
		t.Errorf("expected to error, but got %v", lu)
	}

	fmt.Println("insert at tail")
	lst.InsertAtTail(420)
	lookupAndExpect[int](lst, 3, 420, t)

	fmt.Println("insert at position")
	lst.InsertAtPosition(2, 13)
	lookupAndExpect[int](lst, 2, 13, t)
	lookupAndExpect[int](lst, 3, 23, t)

	fmt.Println("delete at position")

	lst.DeleteAtPosition(3)
	lookupAndExpect[int](lst, 3, 420, t)
}
