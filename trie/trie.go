package main

import (
	"fmt"
	"strings"
)

const LEN_ALPHABET = 26

type Node struct {
	child [LEN_ALPHABET]*Node
	isEnd bool
}

func NewNode() *Node {
	node := &Node{}
	for i := range LEN_ALPHABET {
		node.child[i] = nil
	}
	return node
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: NewNode(),
	}
}

func (t *Trie) Insert(word string) {
	curr := t.root
	str := strings.ToLower(strings.ReplaceAll(word, " ", ""))
	for i := 0; i < len(word); i++ {
		idx := str[i] - 'a'
		if curr.child[idx] == nil {
			curr.child[idx] = NewNode()
		}

		curr = curr.child[idx]
	}
	curr.isEnd = true
}

func (t *Trie) Search(key string) bool {
	curr := t.root
	for i := 0; i < len(key); i++ {
		idx := key[i] - 'a'
		if curr.child[idx] == nil {
			return false
		}

		curr = curr.child[idx]
	}
	return curr != nil && curr.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	curr := t.root
	for i := 0; i < len(prefix); i++ {
		idx := prefix[i] - 'a'
		if curr.child[idx] == nil {
			return false
		}

		curr = curr.child[idx]
	}
	return true
}

func main() {
	trie := NewTrie()
	words := []string{"hassam", "kiani"}
	for _, word := range words {
		trie.Insert(word)
	}

	fmt.Println(trie.Search("hassam"))

	fmt.Println(trie.Search("hassa"))

	fmt.Println(trie.StartsWith("hassa"))
}
