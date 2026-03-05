package main

import "fmt"

type MinHeap struct {
	array []int
}

func (h *MinHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.minheapifyUp(len(h.array) - 1)
}

func (h *MinHeap) minheapifyUp(index int) {
	if h.array[parent(index)] > h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MinHeap) Extract() int {
	lastIdx := len(h.array) - 1
	firstIdx := 0

	extracted := h.array[firstIdx]

	h.array[firstIdx] = h.array[lastIdx]

	h.array = h.array[:lastIdx]

	h.minheapifyDown(firstIdx)

	return extracted
}

func (h *MinHeap) minheapifyDown(index int) {
	l, r := left(index), right(index)
	lastIdx := len(h.array) - 1
	childToCompare := 0

	for l <= lastIdx {
		if l == lastIdx {
			childToCompare = l
		} else if h.array[l] < h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index] > h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func (h *MinHeap) swap(a, b int) {
	h.array[a], h.array[b] = h.array[b], h.array[a]
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return 2*index + 1
}

func right(index int) int {
	return 2*index + 2
}

func main() {
	m := &MinHeap{}
	buildHeap := []int{40, 10, 20, 30}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	m.Extract()

	fmt.Println(m)
	m.Extract()

	fmt.Println(m)
	m.Extract()

	fmt.Println(m)
	m.Extract()

	fmt.Println(m)
}
