package _1128

func numEquivDominoPairs(dominoes [][]int) int {
	var pairs int
	matches := make(map[int]int)
	for _, domino := range dominoes {
		if domino[0] > domino[1] {
			domino[0], domino[1] = domino[1], domino[0]
		}
		res := domino[0]*10 + domino[1]
		if _, ok := matches[res]; ok {
			matches[res] += 1
		} else {
			matches[res] = 1
		}
	}
	for _, v := range matches {
		pairs += v * (v - 1) / 2

	}
	return pairs
}
