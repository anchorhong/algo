package _122

func maxProfitByDynamic(prices []int) int {
	n := len(prices)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func maxProfitByDynamicOpt(prices []int) int {
	n := len(prices)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		newDp0 := max(dp0, dp1+prices[i])
		newDp1 := max(dp1, dp0-prices[i])
		dp0, dp1 = newDp0, newDp1
	}
	return dp0
}

func maxProfitByGreedy(prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ans, n := 0, len(prices)
	for i := 1; i < n; i++ {
		ans += max(0, prices[i]-prices[i-1])
	}
	return ans
}
