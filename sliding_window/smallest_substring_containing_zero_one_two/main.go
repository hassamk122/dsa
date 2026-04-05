package main

import (
	"fmt"
	"math"
)

// Smallest window containing 0, 1 and 2
func main() {

	s := "12121"
	fmt.Println(smallestSubString(s))
}

func smallestSubString(s string) int {
	result := math.MaxInt

	var contains_0, contains_1, contains_2 bool
	var zeroIdx, oneIdx, twoIdx int

	for i := range len(s) {
		switch s[i] {
		case '0':
			contains_0 = true
			zeroIdx = i
		case '1':
			contains_1 = true
			oneIdx = i
		case '2':
			contains_2 = true
			twoIdx = i
		}

		if contains_0 && contains_1 && contains_2 {
			current := max(zeroIdx, oneIdx, twoIdx) - min(zeroIdx, oneIdx, twoIdx)
			result = min(result, current)
		}
	}

	if result == math.MaxInt {
		return -1
	}

	return result + 1
}
