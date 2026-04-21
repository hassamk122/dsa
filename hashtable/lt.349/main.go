package main

import "fmt"

/*

Intersection of Two Arrays


Given two integer arrays nums1 and nums2, return an array of their
. Each element in the result must be unique and you may return the result in any order.
*/

func main() {
	num1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	fmt.Println(intersection(num1, nums2))
}

func intersection(nums1 []int, nums2 []int) []int {
	set := map[int]bool{}

	for i := range nums1 {
		if !set[nums1[i]] {
			set[nums1[i]] = true
		}
	}

	var result []int

	for j := range nums2 {
		if set[nums2[j]] {
			set[nums2[j]] = false
			result = append(result, nums2[j])
		}
	}

	return result
}
