package _1052

func maxSatisfied(customers []int, grumpy []int, X int) int {
	n := len(grumpy)
	preSum := make([]int, n+1)
	suffixSum := make([]int, n+1)
	sumx := make([]int, n+1)
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			preSum[i+1] = preSum[i] + customers[i]
		} else {
			preSum[i+1] = preSum[i]
		}
		sumx[i+1] = sumx[i] + customers[i]
		if grumpy[n-i-1] == 0 {
			suffixSum[i+1] = suffixSum[i] + customers[n-i-1]
		} else {
			suffixSum[i+1] = suffixSum[i]
		}
	}
	maxSum := 0
	for i := 0; i < n-X+1; i++ {
		x := sumx[i+X] - sumx[i]
		pre := preSum[i]
		suf := suffixSum[n-i-X]
		maxSum = max(maxSum, x+pre+suf)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
