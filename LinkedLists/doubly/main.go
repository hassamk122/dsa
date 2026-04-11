package main

import "fmt"

type Node struct {
	data int
	Next *Node
	Prev *Node
}

func NewNode(data int) *Node {
	return &Node{
		data: data,
		Next: nil,
		Prev: nil,
	}
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head: nil,
		tail: nil,
	}
}

func (d *DoublyLinkedList) Append(data int) {
	newNode := NewNode(data)

	if d.EmptyList() {
		d.head = newNode
		return
	}

	if d.OnlyOneNode() {
		d.head.Next = newNode
		newNode.Prev = d.head
		d.tail = newNode
		return
	}

	d.tail.Next = newNode
	newNode.Prev = d.tail
	d.tail = newNode
}

func (d *DoublyLinkedList) Prepend(data int) {
	newNode := NewNode(data)

	if d.EmptyList() {
		d.head = newNode
		return
	}

	if d.OnlyOneNode() {
		d.tail = d.head
		d.head = newNode
		d.head.Next = d.tail
		d.tail.Prev = d.head
	}

	d.head.Prev = newNode
	newNode.Next = d.head
	d.head = newNode
}

func (d *DoublyLinkedList) display() {
	curr := d.head

	for curr != nil {
		fmt.Print(curr.data, "->")
		curr = curr.Next
	}
	fmt.Println()
}

func (d *DoublyLinkedList) displayDesc() {
	curr := d.tail

	for curr != nil {
		fmt.Print("<-", curr.data)
		curr = curr.Prev
	}
	fmt.Println()
}

func (d *DoublyLinkedList) EmptyList() bool {
	return d.head == nil && d.tail == nil
}

func (d *DoublyLinkedList) OnlyOneNode() bool {
	return d.head != nil && d.head.Next == nil
}

func main() {
	ddl := NewDoublyLinkedList()

	ddl.Append(1)
	ddl.Append(2)
	ddl.Append(3)
	ddl.Prepend(0)
	ddl.Prepend(-1)
	ddl.Prepend(-2)

	ddl.display()

	ddl.displayDesc()

}
