package main

import "fmt"

/*
You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average
value and return this value. Any answer with a calculation error less than 10-5 will be accepted.
Example 1:

Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75000
Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75


*/

func main() {
	nums := []int{5}
	k := 1

	fmt.Println(findMaxAverage(nums, k))
}

func findMaxAverage(nums []int, k int) float64 {

	if len(nums) == 0 {
		return float64(0)
	}

	if len(nums) == 1 {
		return float64(nums[0])
	}

	windowSum := 0
	windowAvg := 0.0

	for i := range k {
		windowSum += nums[i]
	}

	windowAvg = float64(windowSum) / float64(k)

	maxWindowAvg := windowAvg

	for j := k; j < len(nums); j++ {
		windowSum += nums[j] - nums[j-k]
		windowAvg = float64(windowSum) / float64(k)

		maxWindowAvg = max(maxWindowAvg, windowAvg)
	}

	return maxWindowAvg
}
