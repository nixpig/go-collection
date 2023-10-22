package array_stack_test

import (
	"testing"

	array_stack "github.com/nixpig/go-collection/data-structures-algorithms/array-stack"
)

func popAndAssert[T comparable](stack *array_stack.ArrayStack[T], expectedValue T, t *testing.T) {
	item, err := stack.Pop()
	if err != nil {
		t.Errorf("expected not to error, instead received: %v", err)
	} else {
		if item != expectedValue {
			t.Errorf("expected to receive %v, instead received: %v", expectedValue, item)
		}
	}
}

func peekAndAssert[T comparable](stack *array_stack.ArrayStack[T], expectedValue T, t *testing.T) {
	item, err := stack.Peek()
	if err != nil {
		t.Errorf("expected not to error, instead received: %v", err)
	} else {
		if item != expectedValue {
			t.Errorf("expected to receive %v, instead received: %v", expectedValue, item)
		}
	}
}

func TestArrayStack(t *testing.T) {
	ask := array_stack.ArrayStack[int]{}

	ask.Push(23)
	ask.Push(69)
	ask.Push(42)
	ask.Push(13)

	peekAndAssert[int](&ask, 13, t)

	popAndAssert[int](&ask, 13, t)
	popAndAssert[int](&ask, 42, t)

	ask.Push(666)

	peekAndAssert[int](&ask, 666, t)

	popAndAssert[int](&ask, 666, t)
	popAndAssert[int](&ask, 69, t)
	popAndAssert[int](&ask, 23, t)

	var item int
	var err error

	item, err = ask.Pop()
	if err != nil {
		// passed
	} else {
		t.Errorf("expected to error due to empty stack, instead received: %v", item)
	}

	item, err = ask.Peek()
	if err != nil {
		// passed
	} else {
		t.Errorf("expected to error due to empty stack, instead received: %v", item)
	}
}
