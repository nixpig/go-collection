package main

import (
	"fmt"

	binary_search_tree "github.com/nixpig/datastructures-the-fun-way/go/binary-search-tree"
)

func main() {
	tree := binary_search_tree.BinarySearchTree{}

	tree.Insert(13)
	tree.Insert(7)
	tree.Insert(9)
	tree.Insert(23)
	tree.Insert(1)

	fmt.Println("---[[ Should find ]]------------------")
	fmt.Println("find 23:", tree.Find(23))
	fmt.Println("find 7:", tree.Find(7))
	fmt.Println("find 13:", tree.Find(13))
	fmt.Println("find 9:", tree.Find(9))
	fmt.Println("find 1:", tree.Find(1))

	fmt.Println("")

	fmt.Println("---[[ Should not find ]]--------------")
	fmt.Println("find 111:", tree.Find(111))
	fmt.Println("find 0:", tree.Find(0))
	fmt.Println("find -5:", tree.Find(-5))

}
