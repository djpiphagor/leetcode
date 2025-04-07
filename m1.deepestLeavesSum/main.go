package main

// 1. Deepest Leaves Sum (bi-tree)
// https://leetcode.com/problems/deepest-leaves-sum/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS, breadth-first search, with no level saved
func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	levelArr := []*TreeNode{root}
	for {
		startLevelArrLen := len(levelArr)
		sum := 0
		for _, n := range levelArr {
			sum += n.Val
			if n.Left != nil {
				levelArr = append(levelArr, n.Left)
			}
			if n.Right != nil {
				levelArr = append(levelArr, n.Right)
			}
		}
		if startLevelArrLen == len(levelArr) {
			return sum
		} else {
			levelArr = levelArr[startLevelArrLen:]
		}
	}
}

func main() {
}
