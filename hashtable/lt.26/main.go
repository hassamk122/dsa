package main

import "fmt"

/*
Remove Duplicates from Sorted Array
*/
func main() {
	nums := []int{1, 1, 2}

	k := removeDuplicates(nums)

	fmt.Println(k)
}

func removeDuplicates(nums []int) int {
	seen := make(map[int]bool, len(nums))

	var newArr []int

	for i := range nums {
		if !seen[nums[i]] {
			seen[nums[i]] = true
			newArr = append(newArr, nums[i])
		}
	}

	copy(nums, newArr)

	return len(newArr)
}
