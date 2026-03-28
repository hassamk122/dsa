package main

import (
	"fmt"
	"strings"
)

// Given a string , write a function to check if it is a permutation of a palindrome

func main() {
	input := "Tact coa"

	fmt.Println(isPermutationOfPalindrome(input))
}

func isPermutationOfPalindrome(phrase string) bool {
	p := strings.ToLower(strings.ReplaceAll(phrase, " ", ""))
	table := buildCharFrequencyTable(p)

	return checkMaxOneOdd(table)
}

func checkMaxOneOdd(table []int) bool {
	foundOdd := false

	for _, v := range table {
		if v%2 != 0 {
			if foundOdd {
				return false
			}
			foundOdd = true
		}
	}

	return true
}

func getCharNumber(c rune) int {
	a := byte('a')
	z := byte('z')

	val := byte(c)

	if a <= val && val <= z {
		return int(val) - int(a)
	}

	return -1
}

func buildCharFrequencyTable(phrase string) []int {
	table := [26]int{}

	for _, v := range phrase {
		x := getCharNumber(v)
		if x != -1 {
			table[x]++
		}
	}

	return table[:]
}
