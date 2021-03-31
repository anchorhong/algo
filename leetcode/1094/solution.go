package _1094

func carPooling(trips [][]int, capacity int) bool {
	n := 0
	for _, t := range trips {
		n = max(n, t[2])
	}
	diff := make([]int, n+1)
	for _, trip := range trips {
		num, start, end := trip[0], trip[1], trip[2]
		if num > capacity {
			return false
		}
		diff[start] += num
		diff[end] -= num
	}
	for i := 1; i <= n; i++ {
		diff[i] += diff[i-1]
		if diff[i] > capacity {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
