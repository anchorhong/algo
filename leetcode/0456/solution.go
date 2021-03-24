package _456

func find132pattern(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	var minHeap []int
	min := nums[0]
	for i := 0; i < n; i++ {
		if nums[i] < min {
			min = nums[i]
		}
		minHeap = append(minHeap, min)
	}
	stack := []int{}
	for i := n - 1; i >= 0; i-- {
		if nums[i] > minHeap[i] {
			for len(stack) > 0 && stack[len(stack)-1] <= minHeap[i] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 && stack[len(stack)-1] < nums[i] {
				return true
			}
			stack = append(stack, nums[i])
		}

	}
	return false
}
