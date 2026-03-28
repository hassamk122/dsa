package main

import (
	"fmt"
	"slices"
)

// Give 2 strings, write a method to decide if one is permutation
// of other.

// 2 strings are permutation if they have same elements
// but in different order

func main() {

	a, b := "abba", "baba"

	a2, b2 := "aba", "baa"
	a3, b3 := "ab", "ba"
	a4, b4 := "bba", "baba"
	a5, b5 := "ab", "bb"
	a6, b6 := "aa", "ba"

	fmt.Println(a, b, isPermutation(a, b))
	fmt.Println(a2, b2, isPermutation(a2, b2))
	fmt.Println(a3, b3, isPermutation(a3, b3))
	fmt.Println(a4, b4, isPermutation(a4, b4))
	fmt.Println(a5, b5, isPermutation(a5, b5))
	fmt.Println(a6, b6, isPermutation(a6, b6))

}

func isPermutation(a, b string) bool {
	strLen := len(a)

	if strLen != len(b) {
		return false
	}

	sa := sortString(a)
	sb := sortString(b)

	if sa == sb {
		return true
	}

	return false
}

func sortString(str string) string {
	sr := []rune(str)

	slices.Sort(sr)

	return string(sr)
}
