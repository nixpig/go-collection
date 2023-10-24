package main

import (
	"fmt"
	"os"

	nearest_neighbor "github.com/nixpig/go-collection/data-structures-algorithms/nearest-neighbor"
)

func main() {
	data := []int{23, 13, 7, 14, 666, 42, 69}

	comparison := func(test, target int) int {
		d := test - target
		if d < 0 {
			return d * -1
		} else {
			return d
		}
	}

	nearest, err := nearest_neighbor.LinearScan(data, 18, comparison)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("nearest:", nearest)
}
