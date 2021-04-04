package _781

func numRabbits(answers []int) int {
	var ans int
	counter := make(map[int]int)
	for _, a := range answers {
		if _, ok := counter[a]; !ok {
			counter[a] = 1
		} else {
			counter[a] += 1
		}
	}
	for k, v := range counter {
		num := v / (k + 1) * (k + 1)
		mod := v % (k + 1)
		if mod != 0 {
			num += k + 1
		}
		ans += num
	}
	return ans
}
