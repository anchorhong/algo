package _309

func maxProfit(prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(prices)
	if n == 0 {
		return 0
	}
	buy := make([]int, n)
	sell := make([]int, n)
	buy[0] = -prices[0]
	for i := 1; i < n; i++ {
		if i > 1 {
			buy[i] = max(buy[i-1], sell[i-2]-prices[i])
		} else {
			buy[i] = max(buy[i-1], sell[i-1]-prices[i])
		}
		sell[i] = max(sell[i-1], buy[i-1]+prices[i])
	}
	return sell[n-1]
}

func maxProfitOpt(prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(prices)
	if n == 0 {
		return 0
	}
	buy, sell1, sell2 := -prices[0], 0, 0
	for i := 1; i < n; i++ {
		if i > 1 {
			buy = max(buy, sell2-prices[i])
		} else {
			buy = max(buy, sell1-prices[i])
		}
		sell2 = sell1
		sell1 = max(sell1, buy+prices[i])
	}
	return sell1
}
