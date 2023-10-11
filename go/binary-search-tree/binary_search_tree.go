package binary_search_tree

import "fmt"

type BinarySearchInt int

type TreeNode struct {
	value  BinarySearchInt
	count  int
	parent *TreeNode
	left   *TreeNode
	right  *TreeNode
}

type BinarySearchTree struct {
	root *TreeNode
}

func (tree *BinarySearchTree) Insert(newValue BinarySearchInt) {
	var insert func(currentNode *TreeNode, newValue BinarySearchInt)

	insert = func(currentNode *TreeNode, newValue BinarySearchInt) {
		if currentNode == nil {
			tree.root = &TreeNode{value: newValue, count: 1}
			return
		}

		if newValue == currentNode.value {
			currentNode.count++
			return
		}

		if newValue < currentNode.value {
			if currentNode.left != nil {
				insert(currentNode.left, newValue)
			} else {
				currentNode.left = &TreeNode{value: newValue, parent: currentNode, count: 1}
			}

			return
		}

		if newValue > currentNode.value {
			if currentNode.right != nil {
				insert(currentNode.right, newValue)
			} else {
				currentNode.right = &TreeNode{value: newValue, parent: currentNode, count: 1}
			}

			return
		}
	}

	insert(tree.root, newValue)
}

func (tree *BinarySearchTree) Find(targetValue BinarySearchInt) *TreeNode {
	var find func(currentNode *TreeNode, targetValue BinarySearchInt) *TreeNode

	find = func(currentNode *TreeNode, targetValue BinarySearchInt) *TreeNode {

		if currentNode == nil {
			fmt.Println("   nil")
			return nil
		}

		if targetValue == currentNode.value {
			fmt.Println("   found")
			return currentNode
		}

		if targetValue < currentNode.value && currentNode.left != nil {
			fmt.Println("<- left")
			return find(currentNode.left, targetValue)
		}

		if targetValue > currentNode.value && currentNode.right != nil {
			fmt.Println("   right ->")
			return find(currentNode.right, targetValue)
		}

		return nil
	}

	return find(tree.root, targetValue)
}
