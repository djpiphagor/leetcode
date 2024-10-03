package main

import (
	"fmt"
	"strings"
)

// 6. Unique Morse Code Words
// https://leetcode.com/problems/unique-morse-code-words/

func main() {
	fmt.Println(toMorse("sos"))
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
	fmt.Println(uniqueMorseRepresentations([]string{"a"}))
}

func uniqueMorseRepresentations(words []string) int {
	alphabetMorse := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	alphabetEnglish := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var mapEng = make(map[rune]int, 26)
	for i, r := range alphabetEnglish {
		mapEng[r] = i
	}
	var mapRes = make(map[string]int)
	for _, w := range words {
		var strRepr string
		for _, r := range strings.ToUpper(w) {
			strRepr += alphabetMorse[mapEng[r]]
		}
		mapRes[strRepr]++
	}
	return len(mapRes)
}

func toMorse(str string) string {
	alphabetMorse := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	alphabetEnglish := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var mapEng = make(map[rune]int, 26)
	for i, r := range alphabetEnglish {
		mapEng[r] = i
	}
	var res string
	for _, r := range strings.ToUpper(str) {
		res += alphabetMorse[mapEng[r]]
	}
	return res
}
