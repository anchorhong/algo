package _239

import (
	"container/heap"
	"container/list"
	"sort"
)

func maxSlidingWindow(nums []int, k int) []int {
	queue := list.New()
	n := len(nums)
	var res []int
	for i := 0; i < n; i++ {
		if queue.Len() == 0 {
			queue.PushBack(i)
		} else {
			for queue.Len() > 0 && nums[queue.Back().Value.(int)] <= nums[i] {
				queue.Remove(queue.Back())
			}
			queue.PushBack(i)
		}
		if queue.Len() > 0 && queue.Front().Value.(int) <= i-k {
			queue.Remove(queue.Front())
		}
		if i >= k-1 {
			res = append(res, nums[queue.Front().Value.(int)])
		}
	}
	return res
}

type maxHeap struct {
	nums []int
	sort.IntSlice
}

func (h *maxHeap) Less(i, j int) bool {
	return h.nums[h.IntSlice[i]] > h.nums[h.IntSlice[j]]
}

func (h *maxHeap) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *maxHeap) Pop() interface{} {
	slc := h.IntSlice
	v := slc[len(slc)-1]
	h.IntSlice = slc[:len(slc)-1]
	return v
}

func maxSlidingWindow1(nums []int, k int) []int {
	hp := &maxHeap{nums, make([]int, k)}
	for i := 0; i < k; i++ {
		hp.IntSlice[i] = i
	}
	heap.Init(hp)
	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[hp.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(hp, i)
		for hp.IntSlice[0] <= i-k {
			heap.Pop(hp)
		}
		ans = append(ans, nums[hp.IntSlice[0]])
	}
	return ans
}
