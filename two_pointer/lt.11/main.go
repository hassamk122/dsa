package main

import "fmt"

/*
Container With Most Water

Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented

by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
*/
func main() {
	height := []int{1, 1}

	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	max_area := 0

	for left < right {
		width := right - left

		hei_ght := min(height[right], height[left])

		area := width * hei_ght

		max_area = max(max_area, area)

		if height[right] < height[left] {
			right--
		} else {
			left++
		}
	}

	return max_area
}
