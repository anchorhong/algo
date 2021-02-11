package _703

import "container/heap"

type KthLargest struct {
	kTh     int
	intHeap intHeap
}

type intHeap []int

func (in intHeap) Len() int {
	return len(in)
}

func (in intHeap) Less(i, j int) bool {
	return in[i] < in[j]
}
func (in intHeap) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

func (in *intHeap) Push(x interface{}) {
	*in = append(*in, x.(int))
}

func (in *intHeap) Pop() interface{} {
	old := *in
	n := len(old)
	x := old[n-1]
	*in = old[0 : n-1]
	return x
}

func Constructor(k int, nums []int) KthLargest {
	x := KthLargest{
		kTh:     k,
		intHeap: []int{},
	}
	for _, v := range nums {
		_ = x.Add(v)
	}
	return x
}

func (this *KthLargest) Add(val int) int {
	if len(this.intHeap) < this.kTh {
		heap.Push(&this.intHeap, val)
	} else if val > this.intHeap[0] {
		heap.Pop(&this.intHeap)
		heap.Push(&this.intHeap, val)
	}
	if len(this.intHeap) == this.kTh {
		return this.intHeap[0]
	} else {
		return 0
	}
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
