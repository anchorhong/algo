package _1202

import (
	"container/heap"
	"strings"
)

//https://leetcode-cn.com/problems/smallest-string-with-swaps/
func smallestStringWithSwaps(s string, pairs [][]int) string {
	sMapping := make(map[int]string)
	var id []int
	var weights []int
	for i := 0; i < len(s); i++ {
		sMapping[i] = string(s[i])
		id = append(id, i)
		weights = append(weights, 1)
	}
	uf := unionFind{
		id:      id,
		weights: weights,
		count:   len(s),
	}
	for _, pair := range pairs {
		uf.union(pair[0], pair[1])
	}
	group := make(map[int]*priorityQueue)
	for i, _ := range uf.id {
		root := uf.find(i)
		if pq, ok := group[root]; ok {
			heap.Push(pq, sMapping[i])
		} else {
			pq := make(priorityQueue, 0)
			heap.Push(&pq, sMapping[i])
			group[root] = &pq
		}
	}
	var res []string
	for i := 0; i < len(s); i++ {
		root := uf.find(uf.id[i])
		pq := group[root]
		if pq.Len() > 0 {
			item := heap.Pop(pq).(string)
			res = append(res, item)
		}
	}
	return strings.Join(res, "")
}

type priorityQueue []string

func (p priorityQueue) Len() int {
	return len(p)
}

func (p priorityQueue) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p priorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *priorityQueue) Push(c interface{}) {
	*p = append(*p, c.(string))
}

func (p *priorityQueue) Pop() interface{} {
	old := *p
	n := len(old)
	removed := old[n-1]
	*p = old[0 : n-1]
	return removed
}

type unionFind struct {
	id      []int
	weights []int
	count   int
}

func (u *unionFind) find(q int) int {
	if u.id[q] != q {
		u.id[q] = u.find(u.id[q])
	}
	return u.id[q]
}

func (u *unionFind) union(p, q int) {
	pRoot := u.find(p)
	qRoot := u.find(q)
	if pRoot == qRoot {
		return
	}

	if u.weights[pRoot] < u.weights[qRoot] {
		u.id[pRoot] = qRoot
		u.weights[qRoot] += u.weights[pRoot]
	} else {
		u.id[qRoot] = pRoot
		u.weights[pRoot] += u.weights[qRoot]
	}
	u.count--
}
