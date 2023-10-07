package linked_list

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	lst := LinkedList[int]{}

	fmt.Println("insert at head")
	lst.InsertAtHead(23)
	lst.InsertAtHead(42)
	lst.InsertAtHead(69)

	fmt.Println("lookup")

	lu, err := lst.Lookup(0)
	if err != nil {
		t.Errorf("expected not to error, but got %v", lu)
	}
	if lu != 69 {
		t.Errorf("expected 69, but got %v", lu)
	}

	lu, err = lst.Lookup(1)
	if err != nil {
		t.Errorf("expected not to error, but got %v", lu)
	}
	if lu != 42 {
		t.Errorf("expected 42, but got %v", lu)
	}

	lu, err = lst.Lookup(2)
	if err != nil {
		t.Errorf("expected not to error, but got %v", lu)
	}
	if lu != 23 {
		t.Errorf("expected 23, but got %v", lu)
	}

	lu, err = lst.Lookup(3)
	if err != nil {
	} else {
		t.Errorf("expected to error, but got %v", lu)
	}

	fmt.Println("insert at tail")

	lst.InsertAtTail(420)

	lu, err = lst.Lookup(3)
	if err != nil {
		t.Errorf("expected not to error, but got %v", lu)
	}
	if lu != 420 {
		t.Errorf("expected 420, but got %v", lu)
	}

	fmt.Println("insert at position")

	lst.InsertAtPosition(2, 13)
	lu, err = lst.Lookup(2)
	if err != nil {
		t.Errorf("expected not to error, but got %v", lu)
	}
	if lu != 13 {
		t.Errorf("expected 13, but got %v", lu)
	}

	lu, err = lst.Lookup(3)
	if err != nil {
		t.Errorf("expected not to error, but got %v", lu)
	}
	if lu != 23 {
		t.Errorf("expected 23, but got %v", lu)
	}

	fmt.Println("delete at position")

	lst.DeleteAtPosition(3)
	lu, err = lst.Lookup(3)
	if err != nil {
		t.Errorf("expected not to error, but got %v", lu)
	}
	if lu != 420 {
		t.Errorf("expected 420, but got %v", lu)
	}

}
