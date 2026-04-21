package main

import "fmt"

/*
	First Unique Character in a String

Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.
*/
func main() {
	fmt.Println(firstUniqChar("leetcode"))

	fmt.Println(firstUniqChar("loveleetcode"))

	fmt.Println(firstUniqChar("aabb"))
}

func firstUniqChar(s string) int {
	set := map[rune]int{}

	for _, v := range s {
		set[v]++
	}

	for i, v := range s {
		if set[v] == 1 {
			return i
		}
	}

	return -1
}
