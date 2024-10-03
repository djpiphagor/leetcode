package main

import "fmt"

// 11. Running Sum of 1d Array
// https://leetcode.com/problems/running-sum-of-1d-array/

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4}))
}

func runningSum(nums []int) []int {
	var res = make([]int, len(nums))
	res[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] + nums[i]
	}
	return res
}
