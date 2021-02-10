package _567

func checkInclusion(s1 string, s2 string) bool {
	wLen, n := len(s1), len(s2)
	if n < wLen {
		return false
	}
	var counter [26]int
	for _, c := range s1 {
		counter[c-'a']++
	}
	left, right := 0, 0
	for right < n {
		for right-left < wLen && right < n {
			counter[s2[right]-'a']--
			right++
		}
		if right-left == wLen {
			match := true
			for _, c := range counter {
				if c != 0 {
					counter[s2[left]-'a']++
					left++
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}
