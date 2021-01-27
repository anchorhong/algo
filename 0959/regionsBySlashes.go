package _959

func regionsBySlashes(grid []string) int {
	n := len(grid)
	uf := newUnionFind(4 * n * n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == ' ' {
				uf.union(4*n*i+4*j, 4*n*i+4*j+1)
				uf.union(4*n*i+4*j+1, 4*n*i+4*j+2)
				uf.union(4*n*i+4*j+2, 4*n*i+4*j+3)
			} else if grid[i][j] == '/' {
				uf.union(4*n*i+4*j, 4*n*i+4*j+1)
				uf.union(4*n*i+4*j+2, 4*n*i+4*j+3)
			} else {
				uf.union(4*n*i+4*j, 4*n*i+4*j+3)
				uf.union(4*n*i+4*j+1, 4*n*i+4*j+2)
			}
			if j < n-1 {
				uf.union(4*n*i+4*j+3, 4*n*i+4*j+5)
			}
			if i < n-1 {
				uf.union(4*n*i+4*j+2, 4*(i+1)*n+4*j)
			}
		}
	}
	return uf.count
}

type unionFind struct {
	parent []int
	rank   []int
	count  int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 1
	}
	return &unionFind{parent, rank, n}
}

func (u *unionFind) find(q int) int {
	if u.parent[q] != q {
		u.parent[q] = u.find(u.parent[q])
	}
	return u.parent[q]
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
