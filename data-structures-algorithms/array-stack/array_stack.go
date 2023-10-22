package array_stack

import "fmt"

type ArrayStack[T any] struct {
	stack []arrayStackItem[T]
	top   int
}

type arrayStackItem[T any] struct {
	value T
}

func (s *ArrayStack[T]) Push(item T) {
	s.stack = append(s.stack, arrayStackItem[T]{value: item})
	s.top++
}

func (s *ArrayStack[T]) Pop() (T, error) {
	if len(s.stack) > 0 && s.top >= 0 {
		topItem := s.stack[len(s.stack)-1]
		s.stack = s.stack[:len(s.stack)-1]

		return topItem.value, nil
	} else {
		r := arrayStackItem[T]{}
		return r.value, fmt.Errorf("nothing on stack to pop")
	}
}

func (s *ArrayStack[T]) Peek() (T, error) {
	if len(s.stack) > 0 {
		return s.stack[len(s.stack)-1].value, nil
	}

	return arrayStackItem[T]{}.value, fmt.Errorf("no items in the stack to peek at")
}
