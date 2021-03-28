package _003

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	left, right := 0, 1
	maxLen := 1
	set := make(map[uint8]bool)
	set[s[0]] = true
	for right < n {
		v := s[right]
		for set[v] {
			delete(set, s[left])
			left++
		}
		set[v] = true
		maxLen = max(maxLen, len(set))
		right++
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
