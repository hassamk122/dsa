package main

import "fmt"

/*
Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.
*/
func main() {
	arr := []int{2, 2}

	fmt.Println(findFirst(arr, 2))
	fmt.Println(findLast(arr, 2))
	fmt.Println(searchRange(arr, 2))

	fmt.Println(searchRange(arr, 7))

	fmt.Println(searchRange(arr, 6))
}

func searchRange(nums []int, target int) []int {

	if len(nums) == 1 && nums[0] == target {
		return []int{0, 0}
	}

	resultArr := make([]int, 2)

	resultArr[0] = findFirst(nums, target)
	resultArr[1] = findLast(nums, target)

	return resultArr
}

func findFirst(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] == target {
			if mid-1 < 0 {
				return mid
			}

			if nums[mid-1] < nums[mid] {
				return mid
			}

			high = mid - 1
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func findLast(nums []int, target int) int {

	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] == target {
			if mid+1 >= len(nums) {
				return mid
			}

			if nums[mid+1] > nums[mid] {
				return mid
			}

			low = mid + 1
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
