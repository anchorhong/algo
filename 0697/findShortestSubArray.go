package _697

func findShortestSubArray(nums []int) (ans int) {
	var maxDegree int
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v] += 1
		if freq[v] > maxDegree {
			maxDegree = freq[v]
		}
	}
	counter := make(map[int]int)
	ans = len(nums)
	for left, right := 0, 0; right < len(nums); right++ {
		counter[nums[right]] += 1
		if counter[nums[right]] == maxDegree {
			for left < right && nums[left] != nums[right] {
				counter[nums[left]] -= 1
				left++
			}
			ans = min(ans, right-left+1)
			counter[nums[right]] -= 1
		}
	}
	return
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
