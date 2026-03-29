package main

import (
	"fmt"
	"strings"
)

// string compression
// implement a method to perform basic string compression using the counts
// of repeated chars
// for example aabccccaa would be a2b1c5a3.
// if compressed is not smaller than original return original

func main() {
	input := "aabccccaa"

	fmt.Println(compressString(input))
}

func compressString(str string) string {
	lenStr := len(str)
	var compStr strings.Builder
	repeatCount := 0

	for i := range lenStr {
		repeatCount++
		if i+1 >= lenStr || str[i] != str[i+1] {
			compStr.WriteString(fmt.Sprintf("%s%d", string(str[i]), repeatCount))
			repeatCount = 0
		}
	}

	if compStr.Len() >= lenStr {
		return str
	}

	return compStr.String()
}
