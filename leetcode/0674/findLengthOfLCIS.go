package _674

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := 1
	length := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			length++
		} else {
			if length > max {
				max = length
			}
			length = 1
		}
	}
	if length > max {
		return length
	}
	return max
}
