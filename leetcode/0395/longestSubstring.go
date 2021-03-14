package _395

func longestSubstring(s string, k int) int {
	counter := make(map[uint8]int)
	invalid := make(map[uint8]interface{})
	for i := 0; i < len(s); i++ {
		invalid[s[i]] = nil
		counter[s[i]] += 1
		if counter[s[i]] >= k {
			delete(invalid, s[i])
		}
	}

	var longest func(int, int) int
	longest = func(left int, right int) int {
		if len(invalid) == 0 {
			return right - left + 1
		}
		counter[s[left]] -= 1
		if counter[s[left]] < k && counter[s[left]] > 0{
			invalid[s[left]] = nil
		}else {
			delete(invalid, s[left])
		}
		m1 := longest(left+1, right)
		counter[s[left]] += 1
		if counter[s[left]] < k && counter[s[left]] > 0{
			invalid[s[left]] = nil
		}else {
			delete(invalid, s[left])
		}
		if counter[s[right]] < k && counter[s[right]] > 0{
			invalid[s[right]] = nil
		}else {
			delete(invalid, s[right])
		}
		m2 := longest(left, right-1)
		if counter[s[right]] < k && counter[s[right]] > 0{
			invalid[s[right]] = nil
		}else {
			delete(invalid, s[right])
		}
		return max(m1, m2)
	}
	return longest(0, len(s)-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
