package main

import "fmt"

// 12. Number of Good Pairs
// https://leetcode.com/problems/number-of-good-pairs/

func main() {
	fmt.Println(numIdenticalPairs([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}))
	fmt.Println(numIdenticalPairs2([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}))
}

func numIdenticalPairs(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]^nums[j] == 0 {
				res++
			}
		}
	}
	return res
}

// cool solution from Leetcode : )
func numIdenticalPairs2(nums []int) int {
	var ans int
	var cnt = make(map[int]int)
	for _, x := range nums {
		ans += cnt[x]
		cnt[x]++
	}
	return ans
}
