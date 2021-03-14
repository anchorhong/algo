package _1319

func makeConnected(n int, connections [][]int) int {
	uf := newUnionFind(n)
	var replicated int
	for _, connection := range connections {
		v1 := connection[0]
		v2 := connection[1]
		if uf.isConnected(v1, v2) {
			replicated++
		} else {
			uf.union(v1, v2)
		}
	}
	if uf.count == 1 {
		return 0
	} else if uf.count-1 > replicated {
		return -1
	} else {
		return uf.count - 1
	}
}

type unionFind struct {
	parent []int
	rank   []int
	count  int
}

func newUnionFind(n int) *unionFind {
	var parent []int
	var rank []int
	for i := 0; i < n; i++ {
		parent = append(parent, i)
		rank = append(rank, 1)
	}
	return &unionFind{parent, rank, n}
}

func (u *unionFind) find(q int) int {
	if u.parent[q] != q {
		u.parent[q] = u.find(u.parent[q])
	}
	return u.parent[q]
}
func (u *unionFind) isConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *unionFind) union(p, q int) {
	pRoot := u.find(p)
	qRoot := u.find(q)
	if pRoot == qRoot {
		return
	}

	if u.rank[pRoot] < u.rank[qRoot] {
		u.rank[qRoot] += u.rank[pRoot]
		u.parent[pRoot] = qRoot
	} else {
		u.rank[pRoot] += u.rank[qRoot]
		u.parent[qRoot] = pRoot
	}
	u.count--
}
