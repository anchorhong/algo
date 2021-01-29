package _1631

import "sort"

func minimumEffortPath(heights [][]int) int {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	m, n := len(heights), len(heights[0])
	var edges []edge
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			id := i*n + j
			if i > 0 {
				edges = append(edges, edge{id - n, id, abs(heights[i][j] - heights[i-1][j])})
			}
			if j > 0 {
				edges = append(edges, edge{id - 1, id, abs(heights[i][j] - heights[i][j-1])})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].diff < edges[j].diff
	})
	uf := newUnionFind(n * m)
	for _, e := range edges {
		uf.union(e.s, e.d)
		if uf.isConnected(0, m*n-1) {
			return e.diff
		}
	}
	return 0
}

type edge struct {
	s, d, diff int
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

func (u unionFind) union(p, q int) {
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
}
