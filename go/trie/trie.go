package trie

import (
// "fmt"
)

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	isEntry  bool
	children [26]*TrieNode
}

func (trie *Trie) Insert(value string) {
	if trie.root == nil {
		trie.root = &TrieNode{}
	}

	var trieNodeInsert func(currentNode *TrieNode, value string, index int)

	trieNodeInsert = func(currentNode *TrieNode, value string, index int) {
		if index == len(value) {
			currentNode.isEntry = true
		} else {
			nextLetter := value[index]
			nextIndex := nextLetter - 'a'

			if currentNode.children[nextIndex] == nil {
				currentNode.children[nextIndex] = &TrieNode{}
			}

			trieNodeInsert(currentNode.children[nextIndex], value, index+1)
		}
	}

	trieNodeInsert(trie.root, value, 0)
}

func (trie *Trie) Search(target string) *TrieNode {
	var trieNodeSearch func(currentNode *TrieNode, target string, index int) *TrieNode

	trieNodeSearch = func(currentNode *TrieNode, target string, index int) *TrieNode {
		if index == len(target) {
			if currentNode.isEntry == true {
				return currentNode
			} else {
				return nil
			}
		}

		nextLetter := target[index]
		nextIndex := nextLetter - 'a'

		if currentNode.children[nextIndex] == nil {
			return nil
		}

		return trieNodeSearch(currentNode.children[nextIndex], target, index+1)

	}

	return trieNodeSearch(trie.root, target, 0)
}
