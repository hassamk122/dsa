package main

import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
}

func NewNode[T any](data T) *Node[T] {
	return &Node[T]{
		data: data,
		next: nil,
	}
}

type SinglyLinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		head: nil,
		tail: nil,
	}
}

func (l *SinglyLinkedList[T]) AppendToHead(data T) {

	newNode := NewNode(data)

	if l.EmptyList() {
		l.head = newNode
		return
	}

	if l.OnlyOneNode() {
		l.tail = l.head
		l.head.next = l.tail
	}

	newNode.next = l.head
	l.head = newNode
}

func (l *SinglyLinkedList[T]) AppendToTail(data T) {

	newNode := NewNode(data)

	if l.EmptyList() {
		l.head = newNode
		return
	}

	if l.OnlyOneNode() {
		l.tail = newNode
		l.head.next = l.tail
		return
	}

	l.tail.next = newNode
	l.tail = newNode
}

func (l *SinglyLinkedList[T]) Display() {
	curr := l.head
	for curr != nil {
		fmt.Printf("%v->", curr.data)
		curr = curr.next
	}
	fmt.Println()
}

func (l *SinglyLinkedList[T]) EmptyList() bool {
	return l.head == nil
}

func (l *SinglyLinkedList[T]) OnlyOneNode() bool {
	return l.head != nil && l.head.next == nil
}

func main() {
	sl := NewSinglyLinkedList[string]()
	sl.AppendToHead("1")
	sl.AppendToHead("2")
	sl.AppendToHead("3")
	sl.AppendToHead("5")
	sl.AppendToTail("4")

	sl.Display()
}
