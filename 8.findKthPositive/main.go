package main

import "fmt"

// 8. Kth Missing Positive Number
// https://leetcode.com/problems/kth-missing-positive-number/

func main() {
	fmt.Println(findKthPositive([]int{2, 3, 4, 7, 11}, 5))
	fmt.Println(findKthPositive([]int{1, 2, 3, 4}, 2))
}

func findKthPositive(arr []int, k int) int {
	i := 1
	diff := arr[0] - 1
	for diff <= k && i < len(arr) {
		diff += arr[i] - arr[i-1] - 1
		i++
	}
	if diff > k {
		diff++
	}
	return arr[i-1] + k - diff
}
