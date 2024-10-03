package main

import "fmt"

// 13. Jewels and Stones
// https://leetcode.com/problems/jewels-and-stones/

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
}

func numJewelsInStones(jewels string, stones string) int {
	var ans int
	var jMap = make(map[rune]int)
	for _, ch := range jewels {
		jMap[ch]++
	}
	for _, ch := range stones {
		ans += jMap[ch]
	}
	return ans
}
