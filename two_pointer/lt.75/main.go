package main

import "fmt"

/*
Sort Colors

Example 1:

Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]

*/

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}

	sortColors(nums)

	fmt.Println(nums)
}

func sortColors(nums []int) {
	var newArr []int

	for _, v := range nums {
		if v == 0 {
			newArr = append(newArr, v)
		}
	}

	for _, v := range nums {
		if v == 1 {
			newArr = append(newArr, v)
		}
	}

	for _, v := range nums {
		if v == 2 {
			newArr = append(newArr, v)
		}
	}

	for i := range nums {
		nums[i] = newArr[i]
	}
}
