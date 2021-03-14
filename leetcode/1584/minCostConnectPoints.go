package _1584

import "sort"

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	type edge struct {
		s, t, distance int
	}
	var edges []edge
	for i, p := range points {
		for j := i + 1; j < n; j++ {
			edges = append(edges, edge{i, j, dist(p, points[j])})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})
	uf := newUnionFind(n)
	var res int
	for _, e := range edges {
		if uf.connected(e.s, e.t) {
			continue
		}
		uf.union(e.s, e.t)
		res += e.distance
	}
	return res
}

type unionFind struct {
	id   []int
	rank []int
}

func newUnionFind(n int) *unionFind {
	id := make([]int, n)
	rank := make([]int, n)
	for i := range id {
		id[i] = i
		rank[i] = 1
	}

	return &unionFind{id, rank}
}

func (u *unionFind) find(q int) int {
	if u.id[q] != q {
		u.id[q] = u.find(u.id[q])
	}
	return u.id[q]
}

func (u *unionFind) connected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *unionFind) union(p, q int) {
	pRoot, qRoot := u.find(p), u.find(q)
	if pRoot == qRoot {
		return
	}
	if u.rank[pRoot] < u.rank[qRoot] {
		u.rank[qRoot] += u.rank[pRoot]
		u.id[pRoot] = qRoot
	} else {
		u.rank[pRoot] += u.rank[qRoot]
		u.id[qRoot] = pRoot
	}
}

func dist(p, q []int) int {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	return abs(p[0]-q[0]) + abs(p[1]-q[1])
}
