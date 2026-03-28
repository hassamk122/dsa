package main

import "fmt"

// Implement an algorithm to determine if a string has all unique characters.
// what if you can't use any other data structures
func main() {

	word := "aba"
	word2 := "ab"
	word5 := "has"
	word3 := ""
	word4 := "a"
	word6 := "hassam"

	fmt.Println(word, isUnique(word))
	fmt.Println(word2, isUnique(word2))
	fmt.Println(word3, isUnique(word3))
	fmt.Println(word4, isUnique(word4))
	fmt.Println(word5, isUnique(word5))
	fmt.Println(word6, isUnique(word6))
}

func isUnique(str string) bool {
	strLen := len(str)

	if strLen == 0 {
		return true
	}

	var charSet = [255]bool{}

	charArray := []byte(str)

	for i := range strLen {
		if charSet[charArray[i]] {
			return false
		}
		charSet[charArray[i]] = true
	}

	return true
}
