package main

import (
	"fmt"
)

type AllowedTypes interface {
	int | string | rune
}

type Node[T AllowedTypes] struct {
	data T
	next *Node[T]
}

func NewNode[T AllowedTypes](data T) *Node[T] {
	return &Node[T]{
		data: data,
		next: nil,
	}
}

type SinglyLinkedList[T AllowedTypes] struct {
	head *Node[T]
	tail *Node[T]
}

func NewSinglyLinkedList[T AllowedTypes]() *SinglyLinkedList[T] {
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

func (l *SinglyLinkedList[T]) RemoveFromHead() {
	if l.EmptyList() {
		return
	}

	if l.OnlyOneNode() {
		l.head = nil
		l.tail = nil
		return
	}

	l.head = l.head.next
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

func (l *SinglyLinkedList[T]) RemoveFromTail() {
	if l.EmptyList() {
		return
	}

	if l.OnlyOneNode() {
		l.head = nil
		l.tail = nil
		return
	}

	prev := l.head
	for prev.next.next != nil {
		prev = prev.next
	}

	l.tail = prev
	l.tail.next = nil
}

func (l *SinglyLinkedList[T]) Delete(data T) {
	temp := l.head

	if temp.data == data {
		l.head = l.head.next
		return
	}

	for temp.next != nil {
		if temp.next.data == data {
			temp.next = temp.next.next
			return
		}
		temp = temp.next
	}

}

func (l *SinglyLinkedList[T]) Search(key T) {
	curr := l.head
	for curr != nil {
		if curr.data == key {
			fmt.Println("Found")
			return
		}
		curr = curr.next
	}
	fmt.Println("Not found")
}

func (l *SinglyLinkedList[T]) RemoveAll() {
	l.head = nil
	l.tail = nil
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

	sl.RemoveFromHead()

	sl.Display()

	sl.RemoveFromTail()

	sl.Delete("3")
	sl.Display()

	sl.Search("22")
}
