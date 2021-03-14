package _867

func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	trans := make([][]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			trans[j] = append(trans[j], matrix[i][j])
		}
	}
	return trans
}
