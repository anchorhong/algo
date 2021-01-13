package _123

import "math"

func maxProfit(prices []int) int {
	max := func(x int, y int) int {
		if x >= y {
			return x
		}
		return y
	}
	dp00, dp10, dp01, dp11, dp02 := math.MinInt32, math.MinInt32, math.MinInt32, 0-prices[0], 0
	for i := 1; i < len(prices); i++ {
		tmp00, tmp10, tmp01, tmp11 := dp00, dp10, dp01, dp11
		dp00 = max(tmp00, tmp10+prices[i])
		dp10 = max(tmp10, tmp01-prices[i])
		dp01 = max(tmp01, tmp11+prices[i])
		dp11 = max(tmp11, dp02-prices[i])
	}
	return max(0, max(dp00, dp01))
}
