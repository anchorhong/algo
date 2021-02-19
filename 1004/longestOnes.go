package _1004

import "sort"

func longestOnes(A []int, K int) int {
	n := len(A)
	maxLen := 0
	var queue []int
	i, j := 0, 0
	for ; j < n; j++ {
		if A[j] == 1 {
			continue
		} else {
			queue = append(queue, j)
			if K > 0 {
				K--
			} else {
				maxLen = max(j-i, maxLen)
				if len(queue) > 0 {
					i = queue[0] + 1
					queue = queue[1:]
				} else {
					i = j + 1
				}
			}
		}
	}
	return max(maxLen, j-i)
}

func longestOnesBinary(A []int, K int) int {
	n := len(A)
	P := make([]int, n+1)
	ans := 0
	for i, v := range A {
		P[i+1] = P[i] + 1 - v
	}
	for right, v := range P {
		left := sort.SearchInts(P, v-K)
		ans = max(ans, right-left)
	}
	return ans
}

func longestOnesSlip(A []int, K int) int {
	var left, lsum, rsum int
	var ans int
	for right, v := range A {
		rsum += 1 - v
		if rsum-lsum > K {
			lsum += 1 - A[left]
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
