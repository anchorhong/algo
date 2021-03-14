package _947

func removeStones(stones [][]int) int {
	uf := make([]int, len(stones))
	wt := make([]int, len(stones))
	for i := 0; i < len(stones); i++ {
		uf[i] = i
		wt[i] = 1
	}
	var find func(int) int
	find = func(q int) int {
		if uf[q] != q {
			uf[q] = find(uf[q])
		}
		return uf[q]
	}
	isConnected := func(p, q int) bool {
		return stones[p][0] == stones[q][0] || stones[p][1] == stones[q][1]
	}
	union := func(p, q int) {
		pRoot := find(p)
		qRoot := find(q)
		if pRoot == qRoot {
			return
		}
		if wt[pRoot] > wt[qRoot] {
			uf[qRoot] = pRoot
			wt[pRoot] += wt[qRoot]
		} else {
			uf[pRoot] = qRoot
			wt[qRoot] += wt[pRoot]
		}
	}

	for i := 0; i < len(stones); i++ {
		for j := i + 1; j < len(stones); j++ {
			if isConnected(i, j) {
				union(i, j)
			}
		}
	}
	res := make(map[int]interface{})
	for i, _ := range uf {
		parent := find(i)
		if _, ok := res[parent]; !ok {
			res[parent] = nil
		}
	}
	return len(stones) - len(res)
}
