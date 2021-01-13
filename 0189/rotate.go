package main

func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	n := len(nums)
	for i := 0; i < n; i++ {
		newNums[(i+k)%n] = nums[i]
	}
	copy(nums, newNums)
}

//func rotate_1(nums []int, k int) {
//
//}
