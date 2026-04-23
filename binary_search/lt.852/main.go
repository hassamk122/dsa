package main

import "fmt"

/*
852. Peak Index in a Mountain Array
You are given an integer mountain array arr of length n where the values increase to a peak element and then decrease.

Return the index of the peak element.

Your task is to solve it in O(log(n)) time complexity.

*/

func main() {
	arr := []int{0, 1, 0}

	arr2 := []int{0, 2, 1, 0}

	arr3 := []int{0, 10, 5, 2}

	fmt.Println(peakIndexInMountainArray(arr))
	fmt.Println(peakIndexInMountainArray(arr2))
	fmt.Println(peakIndexInMountainArray(arr3))
}

func peakIndexInMountainArray(arr []int) int {
	left := 0
	right := len(arr) - 1

	for left < right {
		mid := left + (right-left)/2

		if arr[mid] > arr[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
