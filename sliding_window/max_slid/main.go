package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(nums, 3))

}

type Deque struct {
	arr []int
}

func (d *Deque) Append(i int) {
	d.arr = append(d.arr, i)
}
func (d *Deque) RemoveFirst() {
	d.arr = d.arr[1:]
}

func (d *Deque) FirstElement() int {
	return d.arr[0]
}

func (d *Deque) LastElement() int {
	return d.arr[d.Len()-1]
}

func (d *Deque) RemoveLast() {
	d.arr = d.arr[:d.Len()-1]
}

func (d *Deque) Len() int {
	return len(d.arr)
}

func maxSlidingWindow(nums []int, k int) []int {
	var result []int
	var deque Deque

	for i := range nums {

		if deque.Len() > 0 && deque.FirstElement() <= i-k {
			deque.RemoveFirst()
		}

		for deque.Len() > 0 && nums[deque.LastElement()] < nums[i] {
			deque.RemoveLast()
		}

		deque.Append(i)

		if i >= k-1 {
			result = append(result, nums[deque.FirstElement()])
		}
	}

	return result
}
