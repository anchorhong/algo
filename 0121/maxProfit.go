package _121

import "math"

func maxProfit(prices []int) int {
	var maxProfit int
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}

func maxProfitByScan(prices []int) int {
	minPrice := math.MaxInt64
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}

func maxProfitByDynamic(prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp0, dp1 := 0, math.MinInt32
	for i := 0; i < len(prices); i++ {
		dp0 = max(dp0, dp1+prices[i])
		dp1 = max(dp1, -prices[i])
	}
	return dp0
}
