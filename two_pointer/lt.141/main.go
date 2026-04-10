package main

import "fmt"

/*
Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can
be reached again by continuously following the next pointer. Internally, pos is used to denote the
index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	lt1 := &ListNode{Val: 3}
	lt2 := &ListNode{Val: 3}
	lt1.Next = lt2

	lt3 := &ListNode{Val: 0}
	lt2.Next = lt3
	lt4 := &ListNode{Val: -4}
	lt3.Next = lt4

	lt4.Next = lt2

	result := hasCycle(lt1)

	fmt.Println(result)

	fmt.Println(hasCycle(nil))
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next

		if fast.Next != nil {
			fast = fast.Next
		}

		if fast.Next == slow {
			return true
		}
	}

	return false
}
