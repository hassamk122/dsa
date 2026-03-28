package main

import (
	"fmt"
	"strings"
)

// Urlify
// write a method to replace all spaces in string withh "%20"
// you may assume that the string has sufficient space at the end
// to hold additional characters and that you are given the "true" length
// Example  input = "Mr John Smith " 15
// Output "Mr%20John%20Smith"

func main() {
	input := "Mr John Smith  "
	input2 := "Muhammad Hassam 																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																			"

	fmt.Println(input, Urlify(input))

	fmt.Println(input2, Urlify(input2))
}

func Urlify(str string) string {
	strLen := len(str)
	sr := []rune(str)

	spacesAtEnd := countSpacesAtEnd(sr)

	strWithoutSpacesAtEnd := sr[:strLen-spacesAtEnd]

	replaceSpaces := strings.ReplaceAll(string(strWithoutSpacesAtEnd), " ", "%20")

	return replaceSpaces
}

func countSpacesAtEnd(sr []rune) int {

	spaceCount := 0
	for i := len(sr) - 1; i > 0; i-- {
		if sr[i] == 32 {
			spaceCount++
		} else {
			break
		}
	}

	return spaceCount
}
