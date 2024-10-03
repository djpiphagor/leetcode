package main

import "fmt"

// 5. Count of Matches in Tournament
// https://leetcode.com/problems/count-of-matches-in-tournament/

func main() {
	fmt.Println(numberOfMatches(7))
	fmt.Println(numberOfMatches(14))
}

func numberOfMatches(n int) int {
	var res int
	for n > 3 {
		res += n/2 + n%2
		n = n/2 + n%2
	}
	return res
}
