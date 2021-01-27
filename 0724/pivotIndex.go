package _724

func pivotIndex(nums []int) int {
	n := len(nums)
	preSum := make([]int, n)
	for i := 0; i < n; i++ {
		if i-1 >= 0 {
			preSum[i] = nums[i] + preSum[i-1]
		} else {
			preSum[i] = nums[i]
		}
	}
	for i := 0; i < n; i++ {
		var leftSum int
		if i > 0 {
			leftSum = preSum[i-1]
		}
		rightSum := preSum[n-1] - preSum[i]
		if leftSum == rightSum {
			return i
		}
	}
	return -1
}
