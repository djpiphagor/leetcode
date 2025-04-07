package main

// 15. Design Parking System
// https://leetcode.com/problems/design-parking-system/

import "fmt"

type ParkSystem struct {
	schema map[int]int
}

func Constructor(b, m, s int) *ParkSystem {

	return &ParkSystem{
		schema: map[int]int{1: b, 2: m, 3: s},
	}
}

func (p *ParkSystem) addCar(carType int) bool {
	if p.schema[carType] > 0 {
		p.schema[carType] -= 1
		return true
	}
	return false
}

func main() {
	input := [][]int{{1, 1, 0}, {1}, {2}, {3}, {1}}
	p := Constructor(input[0][0], input[0][1], input[0][2])
	for _, add := range input[1:] {
		fmt.Println(p.addCar(add[0]))
	}
}
