package main

import "fmt"

// 4. https://leetcode.com/problems/build-array-from-permutation/

func main() {
	fmt.Println(buildArray([]int{5, 0, 1, 2, 3, 4}))
}

func buildArray(nums []int) []int {
	var res = make([]int, len(nums))
	for i, j := range nums {
		res[i] = nums[j]
	}
	return res
}
