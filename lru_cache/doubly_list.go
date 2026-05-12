package main

type DoublyLinkedList[T any] struct {
	head      *Node[T]
	tail      *Node[T]
	dummyNode *Node[T]
	size      int
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		dummyNode: &Node[T]{},
	}
}

func (dl *DoublyLinkedList[T]) Add(element T) *Node[T] {
	newNode := NewNode(element)

	if dl.head == nil {
		dl.head = newNode
		dl.tail = newNode
	} else {
		newNode.Next = dl.head
		dl.head.Prev = newNode
		dl.head = newNode
	}

	dl.size++

	return newNode
}

func (dl *DoublyLinkedList[T]) Detach(node *Node[T]) {

	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		dl.head = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		dl.tail = node.Prev
	}

	node.Next = nil
	node.Prev = nil

	dl.size--
}

func (dl *DoublyLinkedList[T]) UpdateAndMoveToFirst(node *Node[T], newElement T) *Node[T] {
	if node == nil {
		return dl.dummyNode
	}

	node.Element = newElement

	return dl.MoveToFront(node)
}

func (dl *DoublyLinkedList[T]) MoveToFront(node *Node[T]) *Node[T] {
	if node == nil || node == dl.head {
		return node
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	if node == dl.tail {
		dl.tail = node.Prev
	}

	node.Prev = nil

	node.Next = dl.head

	if dl.head != nil {
		dl.head.Prev = node
	}

	dl.head = node

	return node
}

func (dl *DoublyLinkedList[T]) RemoveLast() *Node[T] {
	if dl.tail == nil {
		return dl.dummyNode
	}

	slNode := dl.tail
	dl.Detach(dl.tail)
	return slNode
}
