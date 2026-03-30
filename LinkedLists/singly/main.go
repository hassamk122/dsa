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
	if l.head == nil {
		l.head = NewNode(data)
	} else {
		newNode := NewNode(data)
		newNode.next = l.head
		l.head = newNode
	}
}

func (l *SinglyLinkedList) Display() {
	curr := l.head
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
}

func main() {
	sl := NewSinglyLinkedList()
	sl.AppendToHead(3)
	sl.AppendToHead(2)
	sl.AppendToHead(1)
	sl.AppendToHead(0)

	sl.Display()
}
