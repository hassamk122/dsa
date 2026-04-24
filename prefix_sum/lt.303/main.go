package main

import "fmt"

/*Given an integer array nums, handle multiple queries of the following type:

 Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.

Implement the NumArray class:

    NumArray(int[] nums) Initializes the object with the integer array nums.
    int sumRange(int left, int right) Returns the sum of the elements of nums
	between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).
*/

type NumArray struct {
	prefixArr []int
}

func Constructor(nums []int) NumArray {
	prefixArr := make([]int, len(nums))

	sum := 0
	for i := range nums {
		sum += nums[i]
		prefixArr[i] = sum
	}

	return NumArray{
		prefixArr: prefixArr,
	}

}

func (this *NumArray) SumRange(left int, right int) int {
	if left < 0 || right > len(this.prefixArr) {
		return -1
	}

	if left == 0 {
		return this.prefixArr[right] + 0
	}

	return this.prefixArr[right] - this.prefixArr[left-1]
}

func main() {
	numArray := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(numArray.SumRange(0, 2))
	fmt.Println(numArray.SumRange(2, 5))
	fmt.Println(numArray.SumRange(0, 5))
}
