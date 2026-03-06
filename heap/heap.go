package main

import (
	"fmt"
)

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// compares with current with parent
// until curr is at suitable place
func (h *MaxHeap) maxHeapifyUp(index int) {
	for index > 0 && h.array[parent(index)] < h.array[index] {
		p := parent(index)
		h.swap(p, index)
		index = parent(index)
	}
}

// kinda like delete where you extract top
// we delete first by swapping first index
// with last index and then making array size
// smaller by index
// finally use heapifydown method to move
// it to suitable place
func (h *MaxHeap) Extract() int {
	if len(h.array) == 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}

	lastIdx := len(h.array) - 1
	firstIdx := 0

	extracted := h.array[firstIdx]

	h.array[firstIdx] = h.array[lastIdx]

	h.array = h.array[:lastIdx]

	h.maxHeapifyDown(firstIdx)

	return extracted
}

// find left and right of current
// loop while index has at least one child
// when left child is only child compare to left
// when left child is greater than right compare to left
// else when right child is larger compare to right
// compare array values of current index to larger child and swap if smaller
// return if not
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIdx := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIdx {
		if l == lastIdx {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}

}

func (h *MaxHeap) swap(a, b int) {
	h.array[a], h.array[b] = h.array[b], h.array[a]
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func main() {
	score := []int{5, 4, 3, 2, 1}

	retunredArr := findRelativeRanks(score)

	fmt.Println(retunredArr)

	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(nums, 4))
}

func findKthLargest(nums []int, k int) int {
	m := &MaxHeap{}
	for _, v := range nums {
		m.Insert(v)
	}

	for range k - 1 {
		m.Extract()
	}

	return m.Extract()

}

func findRelativeRanks(score []int) []string {
	m := &MaxHeap{}
	for _, v := range score {
		m.Insert(v)
	}

	tempMap := map[int]string{}

	for i := range score {
		extractedVal := m.Extract()
		switch i {
		case 0:
			tempMap[extractedVal] = "Gold Medal"
		case 1:
			tempMap[extractedVal] = "Silver Medal"
		case 2:
			tempMap[extractedVal] = "Bronze Medal"
		default:
			tempMap[extractedVal] = fmt.Sprintf("%d", i+1)
		}
	}

	newArr := []string{}
	for i := range score {

		newArr = append(newArr, tempMap[score[i]])
	}

	return newArr
}
