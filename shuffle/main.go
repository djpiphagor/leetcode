package main

import "fmt"

// 10. Shuffle the Array
// https://leetcode.com/problems/shuffle-the-array/

func main() {
	fmt.Println(shuffle([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4))
}

func shuffle(nums []int, n int) []int {
	var r = make([]int, 2*n)
	var j int
	for i := 0; i < n; i++ {
		r[j], r[j+1] = nums[i], nums[i+n]
		j += 2
	}
	return r
}
