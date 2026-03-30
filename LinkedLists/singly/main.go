package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func NewNode(data int) *Node {
	return &Node{
		data: data,
		next: nil,
	}
}

type SinglyLinkedList struct {
	head *Node
	tail *Node
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		head: nil,
		tail: nil,
	}
}

func (l *SinglyLinkedList) AppendToHead(data int) {

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

func (l *SinglyLinkedList) AppendToTail(data int) {

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

func (l *SinglyLinkedList) Display() {
	curr := l.head
	for curr != nil {
		fmt.Printf("%d->", curr.data)
		curr = curr.next
	}
	fmt.Println()
}

func (l *SinglyLinkedList) EmptyList() bool {
	return l.head == nil
}

func (l *SinglyLinkedList) OnlyOneNode() bool {
	return l.head != nil && l.head.next == nil
}

func main() {
	sl := NewSinglyLinkedList()
	sl.AppendToHead(3)
	sl.AppendToHead(2)
	sl.AppendToHead(1)
	sl.AppendToHead(0)
	sl.AppendToTail(4)

	sl.Display()
}
