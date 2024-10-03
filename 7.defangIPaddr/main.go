package main

import (
	"fmt"
)

// 7. Defanging an IP Address
// https://leetcode.com/problems/defanging-an-ip-address/

func defangIPaddr(address string) string {
	//return strings.ReplaceAll(address, ".", "[.]")
	var res string
	for _, ch := range address {
		if string(ch) == "." {
			res += "[.]"
			continue
		}
		res += string(ch)
	}
	return res
}

func main() {
	fmt.Println(defangIPaddr("0.0.0.0"))
}
