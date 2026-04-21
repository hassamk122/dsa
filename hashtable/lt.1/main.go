package main

import "fmt"

/*
Two Sum
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
*/
func main() {
	arr := []int{2, 7, 11, 15}

	fmt.Println(twoSum(arr, 9))
}

func twoSum(nums []int, target int) []int {

	complimentsMap := map[int]int{}

	for i := range nums {
		compliment, ok := complimentsMap[nums[i]]
		if ok {
			return []int{compliment, i}
		}
		complimentsMap[target-nums[i]] = i
	}

	return nums
}
