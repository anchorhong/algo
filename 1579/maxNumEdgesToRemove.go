package _1579

import "sort"

func maxNumEdgesToRemove(n int, edges [][]int) int {
	alice := newUnionFind(n)
	bob := newUnionFind(n)
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] > edges[j][0]
	})
	var deleted int
	for _, edge := range edges {
		edgeType := edge[0]
		p := edge[1] - 1
		q := edge[2] - 1
		if edgeType == 3 {
			if alice.isConnected(p, q) || bob.isConnected(p, q) {
				deleted++
			} else {
				alice.union(p, q)
				bob.union(p, q)
			}
		} else if edgeType == 1 {
			if alice.isConnected(p, q) {
				deleted++
			} else {
				alice.union(p, q)
			}
		} else {
			if bob.isConnected(p, q) {
				deleted++
			} else {
				bob.union(p, q)
			}
		}
	}
	if alice.count != 1 || bob.count != 1 {
		return -1
	}
	return deleted
}

type unionFind struct {
	parent []int
	wt     []int
	count  int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	var wt []int
	for i := 0; i < n; i++ {
		parent[i] = i
		wt = append(wt, 1)
	}
	return &unionFind{parent, wt, n}
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

	if u.wt[pRoot] < u.wt[qRoot] {
		u.wt[qRoot] += u.wt[pRoot]
		u.parent[pRoot] = qRoot
	} else {
		u.wt[pRoot] += u.wt[qRoot]
		u.parent[qRoot] = pRoot
	}
	u.count--
}