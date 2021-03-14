package _304

type NumMatrix struct {
	preSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 {
		return NumMatrix{preSum: make([][]int, 0)}
	}
	m, n := len(matrix), len(matrix[0])
	preSum := make([][]int, m+1)
	preSum[0] = make([]int, n+1)
	for i := 0; i < m; i++ {
		preSum[i+1] = make([]int, n+1)
		for j := 0; j < n; j++ {
			preSum[i+1][j+1] = matrix[i][j] + preSum[i+1][j] + preSum[i][j+1] - preSum[i][j]
		}
	}
	return NumMatrix{preSum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row2+1][col2+1] - this.preSum[row2+1][col1] - this.preSum[row1][col2+1] + this.preSum[row1][col1]
}
