package _765

func minSwapsCouples(row []int) int {
	n := len(row)
	uf := newUnionFind(n / 2)
	for i := 0; i < n; i += 2 {
		uf.union(row[i]/2, row[i+1]/2)
	}
	return n/2 - uf.count
}

func minSwapCouplesDFS(row []int) int{
	n := len(row)
	edges := make([][]int, n)
	for i := 0; i < n; i += 2 {
		l, r := row[i]/2, row[i+1]/2
		if l != r {
			edges[l] = append(edges[l], r)
			edges[r] = append(edges[r], l)
		}
	}
	visited := make([]bool, n)
	var count int
	var dfs func(v int)
	dfs = func(v int) {
		if !visited[v] {
			visited[v] = true
			for _, u := range edges[v] {
				dfs(u)
			}
		}
	}
	for u := 0; u < n/2; u++ {
		if !visited[u] {
			dfs(u)
			count++
		}
	}
	return n/2 - count
}

func minSwapCouplesBFS(row []int) int {
	n := len(row)
	edges := make([][]int, n/2)
	for i := 0; i < n; i += 2 {
		l, r := row[i]/2, row[i+1]/2
		if l != r {
			edges[l] = append(edges[l], r)
			edges[r] = append(edges[r], r)
		}
	}
	ans := 0
	visited := make([]bool, n/2)
	for i, v := range visited {
		if !v {
			visited[i] = true
			cnt := 0
			q := []int{i}
			for len(q) != 0 {
				cnt++
				q = q[1:]
				for _, u := range edges[i] {
					if !visited[u] {
						q = append(q, u)
					}
				}
			}
			ans += cnt - 1
		}
	}
	return ans
}

type unionFind struct {
	parent, rank []int
	count        int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
		rank[i] = 1
	}
	return &unionFind{
		parent, rank, n,
	}
}

func (u *unionFind) find(p int) int {
	if u.parent[p] != p {
		u.parent[p] = u.find(u.parent[p])
	}
	return u.parent[p]
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
	u.count--
}
