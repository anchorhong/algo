package _424

func characterReplacement(s string, k int) int {
	n := len(s)
	if n < 2 {
		return n
	}
	freq := make([]int, 26)
	var left, right, maxCount int
	for right < n {
		freq[s[right]-'A']++
		maxCount = max(maxCount, freq[s[right]-'A'])
		right++
		if right-left > maxCount+k {
			freq[s[left]-'A']--
			left++
		}
	}
	return right - left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
