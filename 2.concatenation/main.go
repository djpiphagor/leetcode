package main

import "fmt"

// 2. Concatenation of Array
// https://leetcode.com/problems/concatenation-of-array/

func main() {
	fmt.Println(getConcatenation([]int{1, 3, 2, 1}))
	fmt.Println(getConcatenation2([]int{1, 3, 2, 1}))
	fmt.Println(getConcatenation3([]int{1, 3, 2, 1}))
}

func getConcatenation(nums []int) []int {
	var res = make([]int, 0, 2*len(nums))
	res = append(res, nums...)
	res = append(res, nums...)
	return res
}

func getConcatenation2(nums []int) []int {
	var res = make([]int, 2*len(nums))
	for i := 0; i < len(nums); i++ {
		res[i], res[i+len(nums)] = nums[i], nums[i]
	}
	return res
}

func getConcatenation3(nums []int) []int {
	return append(nums, nums...)
}
