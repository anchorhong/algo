package _354

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n == 0 {
		return 0
	}
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] == envelopes[j][0] && envelopes[i][1] < envelopes[j][1])
	})

	f := make([]int, n)
	for i := range f {
		f[i] = 1
	}
	maxCount := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if envelopes[j][1] < envelopes[i][1] {
				f[i] = max(f[i], f[j]+1)
			}
			maxCount = max(f[i], maxCount)
		}
	}
	return  maxCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
