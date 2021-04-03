package _143

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var isEq int
			if text1[i] == text2[j] {
				isEq = 1
			}
			if i > 0 && j > 0 {
				dp[i][j] = max(dp[i-1][j-1]+isEq, dp[i-1][j])
				dp[i][j] = max(dp[i][j], dp[i][j-1])
			} else if i > 0 {
				dp[i][j] = max(dp[i-1][j], isEq)
			} else if j > 0 {
				dp[i][j] = max(dp[i][j-1], isEq)
			} else {
				dp[i][j] = isEq
			}
			res = max(res, dp[i][j])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
