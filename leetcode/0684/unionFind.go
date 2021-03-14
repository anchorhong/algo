package _684

//https://leetcode-cn.com/problems/redundant-connection/
func findRedundantConnectionByUF(edges [][]int) []int {
	uf := newGraph(len(edges) + 1)
	var res []int
	for _, edge := range edges {
		p := edge[0]
		q := edge[1]
		if uf.isConnected(p, q) {
			res = edge
			continue
		}
		uf.union(p, q)
	}
	return res
}

type graph struct {
	id    []int
	sz    []int
	count int
}

func newGraph(n int) graph {
	id := make([]int, n)
	var sz []int
	for i := 0; i < n; i++ {
		id[i] = i
		sz = append(sz, 1)
	}
	return graph{
		id:    id,
		sz:    sz,
		count: n - 1,
	}
}

func (g graph) find(p int) int {
	if g.id[p] != p {
		g.id[p] = g.find(g.id[p])
	}
	return g.id[p]
}

func (g graph) isConnected(p, q int) bool {
	return g.find(p) == g.find(q)
}

func (g graph) union(p, q int) {
	pRoot := g.find(p)
	qRoot := g.find(q)

	if pRoot == qRoot {
		return
	}

	if g.sz[pRoot] > g.sz[qRoot] {
		g.id[qRoot] = pRoot
		g.sz[qRoot] += g.sz[qRoot]
	} else {
		g.id[pRoot] = qRoot
		g.sz[qRoot] += g.sz[pRoot]
	}
	g.count--
}
