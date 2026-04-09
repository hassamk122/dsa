package main

import "fmt"

/*
Merge Two Sorted Lists

You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := &ListNode{Val: 1}
	list1.Next = &ListNode{Val: 2}
	list1.Next.Next = &ListNode{Val: 4}

	list2 := &ListNode{Val: 1}
	list2.Next = &ListNode{Val: 3}
	list2.Next.Next = &ListNode{Val: 4}

	mergedtList := mergeTwoLists(list1, list2)

	curr := mergedtList
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	mergedList := &ListNode{}
	curr := mergedList

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}

		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}

	return mergedList.Next
}
