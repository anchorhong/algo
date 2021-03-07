package _131

func partition(s string) [][]string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	for l := 2; l <= n; l++ {
		for i := 0; i < n-l+1; i++ {
			j := i + l - 1
			if s[i] == s[j] && (l == 2 || dp[i+1][j-1]) {
				dp[i][j] = true
			}
		}
	}
	var res [][]string
	var dfs func(int)
	var combine []string
	dfs = func(start int) {
		if start == n {
			res = append(res, append([]string{}, combine...))
		}

		for i := start; i < n; i++ {
			if dp[start][i] {
				combine = append(combine, s[start:i+1])
				dfs(i + 1)
				combine = combine[:len(combine)-1]
			}
		}
	}

	dfs(0)
	return res
}
