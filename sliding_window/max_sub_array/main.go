package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{2, 1, 5, 1, 3, 2}
	k := 3

	fmt.Println(maxSubArraySum(nums, k))
}

func maxSubArraySum(nums []int, k int) int {

	if k > len(nums) {
		return -1
	}

	maxSum := math.MinInt
	windowSum := 0

	for i := range k {
		windowSum += nums[i]
	}

	for j := k; j < len(nums); j++ {
		windowSum += nums[j] - nums[j-k]
		maxSum = maxBetweenThese(maxSum, windowSum)
	}

	return maxSum
}

func maxBetweenThese(a, b int) int {
	if a > b {
		return a
	}
	return b
}
