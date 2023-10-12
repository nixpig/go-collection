package main

import (
	"fmt"

	binary_search_tree "github.com/nixpig/datastructures-the-fun-way/go/binary-search-tree"
)

func main() {
	tree := binary_search_tree.BinarySearchTree{}

	values := []binary_search_tree.BinarySearchInt{50, 23, 67, 14, 38, 60, 81, 6, 17, 27, 42, 59, 63, 78, 92, 1, 7, 21, 29, 58, 91, 95}

	for _, v := range values {
		tree.Insert(v)
	}

	fmt.Println("---[[ Should find ]]------------------")

	var removal binary_search_tree.BinarySearchInt = 81

	node := tree.Find(removal)

	fmt.Println("node:", node)

	tree.Remove(node)

	fmt.Println("")

	fmt.Println("---[[ Should not find ]]--------------")
	fmt.Println("node:", tree.Find(removal))

	// fmt.Println("find 111:", tree.Find(111))
	// fmt.Println("find 0:", tree.Find(0))
	// fmt.Println("find -5:", tree.Find(-5))

	// fmt.Println("find 7:", tree.Find(7))
	// fmt.Println("find 13:", tree.Find(13))
	// fmt.Println("find 9:", tree.Find(9))
	// fmt.Println("find 1:", tree.Find(1))

}
