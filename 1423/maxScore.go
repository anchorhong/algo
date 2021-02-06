package _1423

func maxScore(cardPoints []int, k int) int {
	sum, n := 0, len(cardPoints)
	preSum := make([]int, n+1)
	for i := range cardPoints {
		sum += cardPoints[i]
		preSum[i+1] = sum
	}
	maxSum := max(preSum[k], preSum[n]-preSum[n-k])
	for i := 1; i < k; i++ {
		maxSum = max(maxSum, preSum[i] + preSum[n] - preSum[n-k+i])
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
