package _778

import (
	"sort"
)

func swimInWater(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var edges []edge
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			id := i*n + j
			if i > 0 {
				edges = append(edges, edge{id, id - n, max(grid[i][j], grid[i-1][j])})
			}
			if j > 0 {
				edges = append(edges, edge{id, id - 1, max(grid[i][j], grid[i][j-1])})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].time < edges[j].time
	})
	uf := newUnionFind(n * m)
	for _, edge := range edges {
		uf.union(edge.s, edge.d)
		if uf.isConnected(0, m*n-1) {
			return edge.time
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type edge struct {
	s, d, time int
}

type unionFind struct {
	parent, rank []int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
		rank[i] = 1
	}
	return &unionFind{parent, rank}
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
	pRoot, qRoot := u.find(p), u.find(q)
	if pRoot == qRoot {
		return
	}

	if u.rank[pRoot] < u.rank[qRoot] {
		pRoot, qRoot = qRoot, pRoot
	}
	u.parent[qRoot] = pRoot
	u.rank[pRoot] += u.rank[qRoot]
}
