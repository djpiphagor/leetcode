package main

import "fmt"

// 3. Convert the Temperature
// https://leetcode.com/problems/convert-the-temperature/

func main() {
	fmt.Println(convertTemperature(122.11))
}

func convertTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, celsius*1.80 + 32.00}
}
