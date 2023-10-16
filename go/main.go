package main

import (
	// "fmt"

	"fmt"

	"github.com/nixpig/datastructures-the-fun-way/go/trie"
)

func main() {
	tr := trie.Trie{}

	words := []string{"foo", "bar", "fool", "barista", "foolish"}

	for _, word := range words {
		tr.Insert(word)
	}

	fmt.Println("---")

	search := []string{"foo", "bar", "apple", "fool", "barista", "chimp", "foolish"}

	for _, word := range search {
		found := tr.Search(word)

		if found != nil {
			fmt.Println("found word in trie:", word)
		} else {
			fmt.Println("did not find word in trie:", word)
		}
	}

	fmt.Println("---")

	remove := []string{"fool", "orangutan", "bar"}

	for _, word := range remove {
		found := tr.Search(word)

		if found != nil {
			fmt.Println("found word in trie:", word)
		} else {
			fmt.Println("did not find word in trie:", word)
		}
	}

	for _, word := range remove {
		tr.Delete(word)
	}

	for _, word := range remove {
		found := tr.Search(word)

		if found != nil {
			fmt.Println("found word in trie:", word)
		} else {
			fmt.Println("did not find word in trie:", word)
		}
	}
}
