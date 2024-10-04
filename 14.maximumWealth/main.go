package main

import "fmt"

// 14. Richest Customer Wealth
//https://leetcode.com/problems/richest-customer-wealth/

func main() {
	fmt.Println(maximumWealth([][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}))
}

func maximumWealth(accounts [][]int) int {
	var sMax int
	var s int
	for _, r := range accounts {
		s = 0
		for _, i := range r {
			s += i
		}
		if sMax < s {
			sMax = s
		}
	}
	return sMax
}
