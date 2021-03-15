package _054

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	upper := 0
	down := m - 1
	left := 0
	right := n - 1
	var res []int
	for true {
		for i := left; i <= right; i++ {
			res = append(res, matrix[upper][i])
		}
		upper++
		if upper > down {
			break
		}
		for j := upper; j <= down; j++ {
			res = append(res, matrix[j][right])
		}
		right--
		if right < left {
			break
		}
		for j := right; j >= left; j-- {
			res = append(res, matrix[down][j])
		}
		down--
		if down < upper {
			break
		}
		for j := down; j >= upper; j-- {
			res = append(res, matrix[j][left])
		}
		left++
		if left > right {
			break
		}
	}
	return res
}
