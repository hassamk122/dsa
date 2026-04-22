package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(Search(arr, 4))

	fmt.Println(Search(arr, 3))

	fmt.Println(Search(arr, 2))

	fmt.Println(Search(arr, 1))

	fmt.Println(Search(arr, 5))
}

func Search(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if target > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
