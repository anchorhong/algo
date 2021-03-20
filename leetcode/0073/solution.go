package _073

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	isCol := false
	for _, r := range matrix {
		if r[0] == 0 {
			isCol = true
		}
		for j := 1; j < n; j++ {
			if r[j] == 0 {
				matrix[0][j] = 0
				r[0] = 0
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
		if isCol {
			matrix[i][0] = 0
		}
	}

}
