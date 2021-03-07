package _496

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	var stack []int
	mapping := make(map[int]int)
	for i := 0; i < n2; i++ {
		for len(stack) > 0 && stack[len(stack)-1] < nums2[i] {
			removed := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			mapping[removed] = nums2[i]
		}
		stack = append(stack, nums2[i])
	}
	for len(stack) != 0 {
		removed := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		mapping[removed] = -1
	}
	var ans []int
	for i := 0; i < n1; i++ {
		ans = append(ans, mapping[nums1[i]])
	}
	return ans
}
