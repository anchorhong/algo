package _896

import "sort"

func isMonotonic(A []int) bool {
	return sort.IntsAreSorted(A) || sort.IsSorted(sort.Reverse(sort.IntSlice(A)))
}
