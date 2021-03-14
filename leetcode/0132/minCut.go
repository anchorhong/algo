package _132

import "math"

func minCut(s string) int {
	n := len(s)
	valid := make([][]bool, n)
	for i := range valid {
		valid[i] = make([]bool, n)
		valid[i][i] = true
	}
	for l := 2; l <= n; l++ {
		for i := 0; i < n-l+1; i++ {
			j := i + l - 1
			valid[i][j] = s[i] == s[j] && (l == 2 || valid[i+1][j-1])
		}
	}
	count := make([]int, n)
	for i := 0; i < n; i++ {
		if valid[0][i] {
			continue
		}
		count[i] = math.MaxInt64
		for j := 0; j < i; j++ {
			if valid[j+1][i] {
				count[i] = min(count[i], count[j]+1)
			}
		}

	}
	return count[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
