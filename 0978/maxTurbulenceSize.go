package _978

func maxTurbulenceSize(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 1
	maxSize := 1
	if n >= 2 && arr[0] != arr[1] {
		dp[1] = 2
		maxSize = 2
	}
	for i := 2; i < n; i++ {
		if arr[i] == arr[i-1] {
			dp[i] = 1
		} else if arr[i-1] > arr[i] && arr[i-1] > arr[i-2] {
			dp[i] = dp[i-1] + 1
		} else if arr[i-1] < arr[i] && arr[i-1] < arr[i-2] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 2
		}
		if dp[i] > maxSize {
			maxSize = dp[i]
		}
	}

	return maxSize
}
