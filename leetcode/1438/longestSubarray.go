package _1438

import (
	"container/heap"
)

func longestSubarray(nums []int, limit int) int {
	n := len(nums)
	minHeap := &minIntHeap{}
	maxHeap := &maxIntHeap{}
	windows := make(map[int]int)
	ans := 0
	for l, r := 0, 0; r < n; r++ {
		if _, ok := windows[nums[r]]; ok {
			windows[nums[r]] += 1
		} else {
			windows[nums[r]] = 1
		}
		heap.Push(minHeap, nums[r])
		heap.Push(maxHeap, nums[r])
		for absSub((*minHeap)[0], (*maxHeap)[0]) > limit {
			windows[nums[l]] -= 1
			if windows[nums[l]] == 0 {
				delete(windows, nums[l])
			}
			l++
			for _, ok := windows[(*minHeap)[0]]; !ok; {
				heap.Pop(minHeap)
				_, ok = windows[(*minHeap)[0]]
			}
			for _, ok := windows[(*maxHeap)[0]]; !ok; {
				heap.Pop(maxHeap)
				_, ok = windows[(*maxHeap)[0]]
			}
		}
		ans = max(ans, r-l+1)
	}
	return ans
}

type minIntHeap []int

func (m minIntHeap) Len() int {
	return len(m)
}

func (m minIntHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m minIntHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minIntHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *minIntHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}

type maxIntHeap []int

func (m maxIntHeap) Len() int {
	return len(m)
}

func (m maxIntHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m maxIntHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *maxIntHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *maxIntHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}

func absSub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
