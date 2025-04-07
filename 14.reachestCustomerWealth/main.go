package main

import "fmt"

// 14. Reachest Customer Wealth
// https://leetcode.com/problems/richest-customer-wealth/

func maximumWealth(accounts [][]int) int {
	var bigS int
	for _, c := range accounts {
		var s int
		for _, w := range c {
			s += w
		}
		if s > bigS {
			bigS = s
		}
	}
	return bigS
}

func main() {
	input := [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}
	fmt.Println(maximumWealth(input))
}
