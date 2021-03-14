package _665

func checkPossibility(nums []int) bool {
	canChange := true
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] && canChange {
			canChange = false
			if i > 0 && nums[i+1] < nums[i-1] {
				nums[i+1] = nums[i]
			}
		} else if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}
