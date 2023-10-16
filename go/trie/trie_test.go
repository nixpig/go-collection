package trie_test

import (
	"testing"

	"github.com/nixpig/datastructures-the-fun-way/go/trie"
)

func assertFound(t *testing.T, trie *trie.Trie, value string, expect bool) {
	found := trie.Search(value)

	// expect to find
	if expect && found != nil {
		return
	}

	// expect not to find
	if !expect && found == nil {
		return
	}

	// else, fail
	t.Errorf("incorrect return result for: %s", value)
}

var words = []string{"foo", "bar", "orangutan", "fool", "barista", "foolish"}

func TestTrie(t *testing.T) {

	tr := trie.Trie{}

	for _, word := range words {
		tr.Insert(word)
	}

	for _, word := range words {
		assertFound(t, &tr, word, true)
	}

	tr.Delete("bar")
	tr.Delete("foolish")

	for _, word := range words {
		if word == "bar" || word == "foolish" {
			assertFound(t, &tr, word, false)
		} else {
			assertFound(t, &tr, word, true)
		}
	}
}
