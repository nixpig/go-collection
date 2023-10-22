package trie

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

func (trie *Trie) Search(value string) *TrieNode {
	var trieNodeSearch func(currentNode *TrieNode, value string, index int) *TrieNode

	trieNodeSearch = func(currentNode *TrieNode, value string, index int) *TrieNode {
		if index == len(value) {
			if currentNode.isEntry {
				return currentNode
			} else {
				return nil
			}
		}

		nextLetter := value[index]
		nextIndex := nextLetter - 'a'

		if currentNode.children[nextIndex] == nil {
			return nil
		}

		return trieNodeSearch(currentNode.children[nextIndex], value, index+1)
	}

	return trieNodeSearch(trie.root, value, 0)
}

func (trie *Trie) Delete(value string) {
	var trieNodeDelete func(currentNode *TrieNode, value string, index int)

	trieNodeDelete = func(currentNode *TrieNode, value string, index int) {
		if index == len(value) {
			if currentNode.isEntry {
				currentNode.isEntry = false
			}
		} else {
			nextLetter := value[index]
			nextIndex := nextLetter - 'a'

			if currentNode.children[nextIndex] != nil {
				trieNodeDelete(currentNode.children[nextIndex], value, index+1)
			}
		}
	}

	trieNodeDelete(trie.root, value, 0)
}
