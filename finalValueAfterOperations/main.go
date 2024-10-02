package main

import (
	"fmt"
	"regexp"
)

// 9. Final Value of Variable After Performing Operations
// https://leetcode.com/problems/final-value-of-variable-after-performing-operations/

func main() {
	fmt.Printf("%v -> %d\n", []string{"++X", "++X", "X++"}, finalValueAfterOperations([]string{"++X", "++X", "X++"}))
	fmt.Printf("%v -> %d\n", []string{"X++", "++X", "--X", "X--"}, finalValueAfterOperations([]string{"X++", "++X", "--X", "X--"}))
}

func finalValueAfterOperations(operations []string) int {
	rePlus := regexp.MustCompile(`(\+{2}X|X\+{2})`)
	reMinus := regexp.MustCompile(`(-{2}X|X-{2})`)
	var x int
	for _, el := range operations {
		if rePlus.Match([]byte(el)) {
			x++
			continue
		}
		if reMinus.Match([]byte(el)) {
			x--
			continue
		}
	}
	return x
}
