package _714

func maxProfit(prices []int, fee int) int {
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
		buy[i] = max(buy[i-1], sell[i-1]-prices[i])
		sell[i] = max(sell[i-1], buy[i-1]+prices[i]-fee)
	}
	return sell[n-1]
}

func maxProfitOpt(prices []int, fee int) int {
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
	buy, sell := -prices[0], 0
	for i := 1; i < n; i++ {
		newBuy, newSell := buy, sell
		buy = max(newBuy, newSell-prices[i])
		sell = max(newSell, newBuy+prices[i]-fee)
	}
	return sell
}

func maxProfitByGreedy(prices []int, fee int) int {
	n := len(prices)
	buy := prices[0] + fee
	profit := 0
	for i := 1; i < n; i++ {
		if prices[i]+fee < buy {
			buy = prices[i] + fee
		} else if prices[i] > buy {
			profit += prices[i] - buy
			buy = prices[i]
		}
	}
	return profit
}
