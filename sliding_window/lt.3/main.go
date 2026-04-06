package main

import "fmt"

// Given a string s, find the length of the longest without duplicate characters.
/*
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct answers.
*/
func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	set := [255]bool{}

	left := 0
	right := 0
	longest := 0

	for right = range s {
		for set[s[right]] {
			fmt.Println(string(s[left]))
			set[s[left]] = false
			left++
		}

		window := (right - left) + 1
		longest = max(longest, window)

		set[s[right]] = true
	}

	return longest
}
