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

func cleanNodeReferences(node *TreeNode) {
	node.parent = nil
	node.left = nil
	node.right = nil
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
				currentNode.left = &TreeNode{
					value:  newValue,
					parent: currentNode,
					count:  1,
				}
			}

			return
		}

		if newValue > currentNode.value {
			if currentNode.right != nil {
				insert(currentNode.right, newValue)
			} else {
				currentNode.right = &TreeNode{
					value:  newValue,
					parent: currentNode,
					count:  1,
				}
			}

			return
		}
	}

	insert(tree.root, newValue)
}

func (tree *BinarySearchTree) Find(targetValue BinarySearchInt) *TreeNode {
	var find func(currentNode *TreeNode, targetValue BinarySearchInt) *TreeNode

	find = func(currentNode *TreeNode, targetValue BinarySearchInt) *TreeNode {

		fmt.Println("current:", currentNode.value)

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

func (tree *BinarySearchTree) Remove(node *TreeNode) {
	if tree.root == nil || node == nil {
		return
	}

	// decrement count, rather than deleting
	if node.count > 1 {
		fmt.Println("node count greater than 1; decrementing")
		node.count--
		return
	}

	// delete a leaf node
	if node.left == nil && node.right == nil {
		fmt.Println("deleting a leaf node")
		if node.parent == nil {
			tree.root = nil
			return
		}

		if node.parent.left == node {
			node.parent.left = nil
			return
		}

		if node.parent.right == node {
			node.parent.right = nil
			return
		}

		return
	}

	// delete a node with one child
	if node.left == nil || node.right == nil {
		fmt.Println("deleting a node with one child")
		var child *TreeNode

		if node.left == nil {
			fmt.Println("left node is nil")
			child = node.right
		}

		if node.right == nil {
			fmt.Println("right node is nil")
			child = node.left
		}

		fmt.Println("update child.parent to be node.parent")
		child.parent = node.parent

		if node.parent == nil {
			fmt.Println("node.parent is nil")
			tree.root = child
		} else if node.parent.left == node {
			fmt.Println("node.parent.left is node")
			node.parent.left = child
		} else if node.parent.right == node {
			fmt.Println("node.parent.right is node")
			node.parent.right = child
		}

		cleanNodeReferences(node)

		return
	}

	// delete a node with two children
	fmt.Println("deleting a node with two children")

	successor := node.right
	for successor.left != nil {
		successor = successor.left
	}

	tree.Remove(successor)

	if node.parent == nil {
		tree.root = successor
	} else if node.parent.left == node {
		node.parent.left = successor
	} else {
		node.parent.right = successor
	}

	successor.parent = node.parent

	successor.left = node.left
	successor.right = node.right

	if node.left != nil {
		node.left.parent = successor
	}

	if node.right != nil {
		node.right.parent = successor
	}

	cleanNodeReferences(node)
}
