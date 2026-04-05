package main

import (
	"fmt"
)

func main() {
	s := "aaa"
	fmt.Println(longestUniqueSubString(s))

}

func longestUniqueSubString(s string) int {
	lenStr := len(s)
	if lenStr == 0 || lenStr == 1 {
		return lenStr
	}

	result := 0

	visited := [26]bool{}

	left, right := 0, 0

	for right < lenStr {

		for visited[s[right]-'a'] {
			visited[s[left]-'a'] = false
			left++
		}

		visited[s[right]-'a'] = true

		result = max(result, (right - left + 1))
		right++
	}

	fmt.Println(visited)
	return result
}
