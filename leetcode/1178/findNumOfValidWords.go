package _1178

import "math/bits"

func findNumOfValidWords(words []string, puzzles []string) []int {
	word2Num := make(map[int]int)
	for _, w := range words {
		bitNum := toBits(w)
		if bits.OnesCount(uint(bitNum)) > 7 {
			continue
		}
		word2Num[bitNum]++
	}
	ans := make([]int, len(puzzles))
	for i, p := range puzzles {
		h := 1 << (p[0] - 'a')
		mask := toBits(p[1:])
		subset := mask
		for {
			ans[i] += word2Num[subset|h]
			subset = (subset - 1) & mask
			if subset == mask {
				break
			}
		}
	}
	return ans
}

func toBits(word string) int {
	bits := 0
	for i := 0; i < len(word); i++ {
		bits |= 1 << (word[i] - 'a')
	}
	return bits
}
