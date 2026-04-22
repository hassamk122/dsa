package main

import (
	"fmt"
)

/*
Search Insert Position

Given a sorted array of distinct integers and a target value,
return the index if the target is found.
If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.
*/

func main() {
	arr := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(arr, 5))

	fmt.Println(searchInsert(arr, 7))
}

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] == target {
			return mid
		}

		if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
