package main

import (
	"fmt"
	"strings"
)

func main() {
	cipherText := `MLD BLF PMLD GSV TZNV SLD GL KOZB RG. TVG 
IVZWB ULI HRNROZI KILYOVNH RM BLFI VCZNH. 
XSVVIH. `

	freqMap := countFrequencyOfEachLetter(cipherText)

	fmt.Println(freqMap)

	nS := replaceLetter(cipherText, "L", "T")

	fmt.Println(replaceLetter(nS, "V", "E"))
}

func replaceLetter(cipherText string, ltrToReplace string, ltrToReplaceWith string) string {
	return strings.ReplaceAll(cipherText, ltrToReplace, ltrToReplaceWith)
}

func countFrequencyOfEachLetter(cipherText string) map[string]int {
	freqMap := map[string]int{}

	rmSpaces := strings.ReplaceAll(cipherText, " ", "")
	rmLines := strings.ReplaceAll(rmSpaces, "\n", "")
	rmDots := strings.ReplaceAll(rmLines, ".", "")

	ct := strings.ToLower(rmDots)
	for i := 0; i < len(ct); i++ {
		ch := string(ct[i])
		freqMap[ch]++
	}
	return freqMap
}
