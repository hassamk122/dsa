package main

import (
	"fmt"
)

/*
Middle of Linked List

Given the head of a singly linked list, return the middle node of the linked list.

If there are two middle nodes, return the second middle node.

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := &ListNode{Val: 1}
	list1.Next = &ListNode{Val: 2}
	list1.Next.Next = &ListNode{Val: 3}
	list1.Next.Next.Next = &ListNode{Val: 4}
	list1.Next.Next.Next.Next = &ListNode{Val: 5}

	result := middleNode(list1)
	fmt.Println(result.Val)
}

func middleNode(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	slow := head
	fast := head

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next

		if fast.Next != nil {
			fast = fast.Next
		}
	}

	return slow
}
