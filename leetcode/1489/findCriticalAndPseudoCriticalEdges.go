package _1489

import "sort"

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})
	uf := newUnionFind(n)
	var weight int
	for _, edge := range edges {
		p1, p2, w := edge[0], edge[1], edge[2]
		if !uf.isConnected(p1, p2) {
			uf.union(p1, p2)
			weight += w
		}
	}



}

func newUnionFind(n int) *unionFind {
	var vertex []int
	var sz []int
	for i := 0; i < n; i++ {
		vertex = append(vertex, i)
		sz = append(sz, 1)
	}
	return &unionFind{vertex, sz, n}
}

type unionFind struct {
	vertex []int
	sz     []int
	count  int
}

func (u *unionFind) find(q int) int {
	if q != u.vertex[q] {
		u.vertex[q] = u.find(u.vertex[q])
	}
	return u.vertex[q]
}

func (u *unionFind) isConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u unionFind) union(p, q int) {
	pRoot := u.find(p)
	qRoot := u.find(q)
	if pRoot == qRoot {
		return
	}

	if u.sz[pRoot] < u.sz[qRoot] {
		u.vertex[pRoot] = qRoot
		u.sz[qRoot] += u.sz[pRoot]
	} else {
		u.vertex[qRoot] = pRoot
		u.sz[pRoot] += u.sz[qRoot]
	}
	u.count--
}
