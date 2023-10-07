package main

import (
	"fmt"

	array_stack "github.com/nixpig/datastructures-the-fun-way/go/array-stack"
)

func main() {
	ask := array_stack.ArrayStack[int]{}

	ask.Push(23)
	ask.Push(69)
	ask.Push(42)
	ask.Push(13)

	var item int
	var err error

	item, err = ask.Pop()
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("value:", item)
	}

	item, err = ask.Pop()
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("value:", item)
	}

	item, err = ask.Pop()
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("value:", item)
	}

	item, err = ask.Pop()
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("value:", item)
	}

	item, err = ask.Pop()
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("value:", item)
	}
}
