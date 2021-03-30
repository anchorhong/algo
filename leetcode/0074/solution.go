package _074

import "sort"

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	lowR, highR := 0, m-1
	for lowR <= highR {
		midRow := (lowR + highR) / 2
		if matrix[midRow][n-1] == target {
			return true
		} else if matrix[midRow][n-1] > target {
			if matrix[midRow][0] <= target {
				left, right := 0, n-1
				for left <= right {
					mid := (left + right) / 2
					if matrix[midRow][mid] == target {
						return true
					} else if matrix[midRow][mid] > target {
						right = mid - 1
					} else {
						left = mid + 1
					}
				}
			}
			highR = midRow - 1
		} else if matrix[midRow][n-1] < target {
			lowR = midRow + 1
		}
	}
	return false
}

func searchMatrix1(matrix [][]int, target int) bool {
	row := sort.Search(len(matrix), func(i int) bool {
		return matrix[i][0] > target
	}) - 1
	if row < 0 {
		return false
	}
	col := sort.SearchInts(matrix[row], target)
	return col < len(matrix[row]) && matrix[row][col] == target
}

func searchMatrix2(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i := sort.Search(m*n, func(i int) bool { return matrix[i/n][i%n] >= target })
	return i < m*n && matrix[i/n][i%n] == target
}
