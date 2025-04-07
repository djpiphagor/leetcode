package main

// Sort the Students by Their Kth Score
// https://leetcode.com/problems/sort-the-students-by-their-kth-score/

func sortTheStudents(score [][]int, k int) [][]int {
	if len(score) <= 1 {
		return score
	}
	mid := len(score) / 2
	up := sortTheStudents(score[:mid], k)
	down := sortTheStudents(score[mid:], k)
	return merge(up, down, k)
}

// using merge sort approach
func merge(up, down [][]int, k int) [][]int {
	res := make([][]int, 0)
	for len(up) > 0 && len(down) > 0 {
		if up[0][k] >= down[0][k] {
			res = append(res, up[0])
			up = up[1:]
		} else {
			res = append(res, down[0])
			down = down[1:]
		}
	}
	res = append(res, up...)
	res = append(res, down...)
	return res
}

func main() {

}
