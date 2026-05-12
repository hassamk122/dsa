package main

type Node[T any] struct {
	Element T
	Next    *Node[T]
	Prev    *Node[T]
}

func NewNode[T any](element T) *Node[T] {
	return &Node[T]{
		Element: element,
	}
}
