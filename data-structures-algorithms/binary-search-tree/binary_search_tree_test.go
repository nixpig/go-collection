package binary_search_tree_test

import (
	"testing"

	binary_search_tree "github.com/nixpig/go-collection/data-structures-algorithms/binary-search-tree"
)

func createTree() binary_search_tree.BinarySearchTree {
	tree := binary_search_tree.BinarySearchTree{}

	values := []binary_search_tree.BinarySearchInt{50, 23, 67, 14, 38, 60, 81, 6, 17, 27, 42, 59, 63, 78, 92, 1, 7, 21, 29, 58, 91, 95}

	for _, v := range values {
		tree.Insert(v)
	}

	return tree
}

func TestFindRootNode(t *testing.T) {
	tree := createTree()

	node := tree.Find(50)
	if node == nil {
		t.Error("expected to find 50, instead got nil")
	}
}

func TestFindInternalNode(t *testing.T) {
	tree := createTree()

	node := tree.Find(60)
	if node == nil {
		t.Error("expected to find 60, instead got nil")
	}
}

func TestFindLeafNode(t *testing.T) {
	tree := createTree()

	node := tree.Find(95)
	if node == nil {
		t.Error("expected to find 95, instead got nil")
	}
}

func TestInsertRootNode(t *testing.T) {
	tree := binary_search_tree.BinarySearchTree{}

	tree.Insert(23)

	node := tree.Find(23)
	if node == nil {
		t.Error("expected to find 95, instead got nil")
	}
}

func TestInsertInternalNode(t *testing.T) {
	tree := createTree()

	tree.Insert(18)
	node := tree.Find(18)
	if node == nil {
		t.Error("expected to find 18, instead got nil")
	}
}

func TestInsertLeafNode(t *testing.T) {
	tree := createTree()

	tree.Insert(100)
	node := tree.Find(100)
	if node == nil {
		t.Error("expected to find 100, instead got nil")
	}
}

func TestRemoveRootNode(t *testing.T) {
	tree := createTree()

	firstFind := tree.Find(50)
	tree.Remove(firstFind)

	secondFind := tree.Find(50)
	if secondFind != nil {
		t.Errorf("expected not to find, instead got %v", secondFind)
	}
}

func TestRemoveLeafNode(t *testing.T) {
	tree := createTree()

	firstFind := tree.Find(58)
	tree.Remove(firstFind)

	secondFind := tree.Find(58)
	if secondFind != nil {
		t.Errorf("expected not to find, instead got %v", secondFind)
	}
}

func TestRemoveInternalNodeWithOneChild(t *testing.T) {
	tree := createTree()

	firstFind := tree.Find(27)
	tree.Remove(firstFind)

	secondFind := tree.Find(27)
	if secondFind != nil {
		t.Errorf("expected not to find, instead got %v", secondFind)
	}
}

func TestRemoveInternalNodeWithTwoChildren(t *testing.T) {
	tree := createTree()

	firstFind := tree.Find(38)
	tree.Remove(firstFind)

	secondFind := tree.Find(38)
	if secondFind != nil {
		t.Errorf("expected not to find, instead got %v", secondFind)
	}
}
