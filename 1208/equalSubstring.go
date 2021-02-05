package _1208

func equalSubstring(s string, t string, maxCost int) int {
	n := len(s)
	sum := 0
	preSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum += absDiff(s[i], t[i])
		preSum[i+1] = sum
	}
	maxLen := 0
	for i, j := 0, 1; i < n+1-maxLen && j < n+1; {
		if preSum[j]-preSum[i] <= maxCost {
			maxLen = max(maxLen, j-i)
			j++
		} else {
			i++
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func absDiff(a, b uint8) int {
	if a > b {
		return int(a - b)
	}
	return int(b - a)
}
