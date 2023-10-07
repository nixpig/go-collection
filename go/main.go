package main

import (
	"fmt"

	linked_list "github.com/nixpig/datastructures-the-fun-way/go/linked-lists"
)

func main() {
	lst := linked_list.LinkedList[int]{}

	lst.InsertAtTail(42)
	lst.InsertAtTail(69)
	lst.InsertAtTail(420)

	val, err := lst.Lookup(1)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(val)

	lst.InsertAtHead(23)
	val, err = lst.Lookup(0)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(val)

	err = lst.InsertAtPosition(2, 111)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	val, err = lst.Lookup(1)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("1:", val)

	val, err = lst.Lookup(2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("2:", val)

	val, err = lst.Lookup(3)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("3:", val)

	err = lst.DeleteAtPosition(2)
	val, err = lst.Lookup(2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("2:", val)

	err = lst.DeleteAtPosition(0)
	val, err = lst.Lookup(2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("2:", val)

}
