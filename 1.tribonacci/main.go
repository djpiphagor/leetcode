package main

import "fmt"

// 1. N-th Tribonacci Number
// https://leetcode.com/problems/n-th-tribonacci-number/

func main() {
	fmt.Println(tribonacci(4))
	fmt.Println(tribonacci(25))
	fmt.Println(fibonacci(4))
	fmt.Println(fibonacci(25))
}

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	a, b, c := 0, 1, 1
	for i := 0; i < n-2; i++ {
		a, b, c = b, c, a+b+c
	}
	return c
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 0; i < n-1; i++ {
		a, b = b, a+b
	}
	return b
}
