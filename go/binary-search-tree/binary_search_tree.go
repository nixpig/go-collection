package binary_search_tree

type binaryTreeValue int

type TreeNode struct {
	value  binaryTreeValue
	count  int
	parent *TreeNode
	left   *TreeNode
	right  *TreeNode
}

type BinarySearchTree struct {
	root *TreeNode
}

func (tree *BinarySearchTree) Insert(newValue binaryTreeValue) {
	var insert func(currentNode *TreeNode, newValue binaryTreeValue)

	insert = func(currentNode *TreeNode, newValue binaryTreeValue) {
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

func (tree *BinarySearchTree) Find(targetValue binaryTreeValue) *TreeNode {
	var find func(currentNode *TreeNode, targetValue binaryTreeValue) *TreeNode

	find = func(currentNode *TreeNode, targetValue binaryTreeValue) *TreeNode {

		if currentNode == nil {
			return nil
		}

		if targetValue == currentNode.value {
			return currentNode
		}

		if targetValue < currentNode.value && currentNode.left != nil {
			return find(currentNode.left, targetValue)
		}

		if targetValue > currentNode.value && currentNode.right != nil {
			return find(currentNode.right, targetValue)
		}

		return nil
	}

	return find(tree.root, targetValue)
}
