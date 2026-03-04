package main

import (
	"fmt"
	"sort"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	fmt.Println(len(strs[0]))
	var result strings.Builder

	for i := 0; i < len(strs[0]); i++ {
		ch := strs[0][i]
		for j := range strs {
			if strs[j][i] != ch {
				return result.String()
			}
		}
		result.WriteString(string(ch))
	}
	return result.String()
}

func main() {
	words := []string{"dog", "racecar", "car"}

	fmt.Println(longestCommonPrefix(words))
}
