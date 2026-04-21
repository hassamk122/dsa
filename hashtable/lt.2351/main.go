package main

import "fmt"

/*
First Letter to Appear Twice
Given a string s consisting of lowercase English letters, return the first letter to appear twice.

Note:

	A letter a appears twice before another letter b if the second occurrence of a is before the second occurrence of b.
	s will contain at least one letter that appears twice.
*/
func main() {
	Input := "abccbaacz"

	fmt.Println(repeatedCharacter(Input))
}

func repeatedCharacter(s string) byte {
	occurrenceMap := make([]bool, 26)

	for i := range s {
		if occurrenceMap[int(s[i])-97] {
			return s[i]
		}
		occurrenceMap[int(s[i])-97] = true
	}

	return byte(' ')
}
