package main

import "fmt"

func main() {
	trie := New()
	trie.Set("gopher")
	trie.Set("golang")

	fmt.Println("gopher", trie.Get("gopher"))
	fmt.Println("golang", trie.Get("golang"))
	fmt.Println("go", trie.Get("go"))
}

type TrieNode struct {
	Items [26]*TrieNode
	IsLast bool
}

type Trie struct {
	Head *TrieNode
}

func New() *Trie {
	return &Trie{
		Head: &TrieNode{},
	}
}

func (t *Trie) Set(word string) {
	length := len(word)
	current := t.Head

	for i := 0; i < length; i++ {
		index := word[i] - 'a'

		if current.Items[index] == nil {
			current.Items[index] = &TrieNode{}
		}

		current = current.Items[index]
	}

	current.IsLast = true
}

func (t *Trie) Get(word string) bool {
	length := len(word)
	current := t.Head

	for i := 0; i < length; i++ {
		index := word[i] - 'a'

		if current.Items[index] == nil {
			return false
		}

		current = current.Items[index]
	}

	return current.IsLast
}