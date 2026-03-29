package main

import "fmt"

// One Away
// There are three tyoes of edits that can be performed on strings
// insert,remove, replace. Given 2 strings write a function to check
// if they are one edit away

func main() {
	a, b := "pale", "ple"

	a2, b2 := "pales", "pale"
	a3, b3 := "pale", "bale"
	a4, b4 := "pale", "bae"

	fmt.Println(a, b, OneAway(a, b))
	fmt.Println(a2, b2, OneAway(a2, b2))
	fmt.Println(a3, b3, OneAway(a3, b3))
	fmt.Println(a4, b4, OneAway(a4, b4))
}

func OneAway(a, b string) bool {
	strAlen := len(a)
	strBlen := len(b)

	if strAlen == strBlen {
		return oneEditReplace(a, b)
	} else if strAlen+1 == strBlen {
		return oneEditInsert(a, b)
	} else if strAlen-1 == strBlen {
		return oneEditInsert(b, a)
	}
	return false
}

func oneEditReplace(a, b string) bool {
	foundDiff := false
	for i := range len(b) {
		if a[i] != b[i] {
			if foundDiff {
				return false
			}

			foundDiff = true
		}
	}

	return true
}
func oneEditInsert(a, b string) bool {
	idxI := 0
	idxJ := 0

	for idxI < len(a) && idxJ < len(b) {
		if a[idxI] != b[idxJ] {
			if idxI != idxJ {
				return false
			}
			idxJ++
		} else {
			idxI++
			idxJ++
		}
	}
	return true
}
