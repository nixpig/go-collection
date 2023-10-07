package array_stack_test

import (
	"testing"

	array_stack "github.com/nixpig/datastructures-the-fun-way/go/array-stack"
)

func TestArrayStack(t *testing.T) {
	ask := array_stack.ArrayStack[int]{}

	ask.Push(23)
	ask.Push(69)
	ask.Push(42)
	ask.Push(13)

	var item int
	var err error

	item, err = ask.Pop()
	if err != nil {
		t.Errorf("expected not to error, instead received: %v", err)
	} else {
		if item != 13 {
			t.Errorf("expected to receive 13, instead received: %v", item)
		}
	}

	item, err = ask.Pop()
	if err != nil {
		t.Errorf("expected not to error, instead received: %v", err)
	} else {
		if item != 42 {
			t.Errorf("expected to receive 42, instead received: %v", item)
		}
	}

	item, err = ask.Pop()
	if err != nil {
		t.Errorf("expected not to error, instead received: %v", err)
	} else {
		if item != 69 {
			t.Errorf("expected to receive 69, instead received: %v", item)
		}
	}

	item, err = ask.Pop()
	if err != nil {
		t.Errorf("expected not to error, instead received: %v", err)
	} else {
		if item != 23 {
			t.Errorf("expected to receive 23, instead received: %v", item)
		}
	}

	item, err = ask.Pop()
	if err != nil {
		// passed
	} else {
		t.Errorf("expected to error due to empty stack, instead received: %v", item)
	}
}
