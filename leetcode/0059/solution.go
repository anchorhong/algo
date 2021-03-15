package _059

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i, _ := range matrix {
		matrix[i] = make([]int, n)
	}
	l, r, u, d := 0, n-1, 0, n-1
	num := 1
	for true {
		for i := l; i <= r; i++ {
			matrix[u][i] = num
			num++
		}
		u++
		if u > d {
			break
		}
		for i := u; i <= d; i++ {
			matrix[i][r] = num
			num++
		}
		r--
		if r < l {
			break
		}
		for i := r; i >= l; i-- {
			matrix[d][i] = num
			num++
		}
		d--
		if d < u {
			break
		}
		for i := d; i >= u; i-- {
			matrix[i][l] = num
			num++
		}
		l++
		if l > r {
			break
		}
	}
	return matrix
}
