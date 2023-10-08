package main

import (
	"fmt"

	ring_buffer "github.com/nixpig/datastructures-the-fun-way/go/ring-buffer"
)

func main() {
	rb := ring_buffer.NewRingBuffer[int](4)

	rb.Enqueue(23)

	rb.Enqueue(42)

	rb.Enqueue(69)

	rb.Enqueue(13)

	// isFull, err = rb.Enqueue(7)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(isFull)
	// }
	//
	// isFull, err = rb.Enqueue(666)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(isFull)
	// }

	value, err := rb.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	value, err = rb.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	rb.Enqueue(666)

	rb.Enqueue(777)

	value, err = rb.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	value, err = rb.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	value, err = rb.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

}
