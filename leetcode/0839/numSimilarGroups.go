package _839

func numSimilarGroups(strs []string) int {
	uf := newUnionFind(len(strs))
	for i := 0; i < len(strs)-1; i++ {
		for j := i + 1; j < len(strs); j++ {
			if uf.isConnected(i, j) {
				continue
			}
			if isSimilar(strs[i], strs[j]) {
				uf.union(i, j)
			}
		}
	}
	return uf.count
}

func isSimilar(str1, str2 string) bool {
	if str1 == str2 {
		return true
	}
	var index []int
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			index = append(index, i)
		}
		if len(index) > 2 {
			return false
		}
	}
	if len(index) != 2 {
		return false
	}
	return str1[index[0]] == str2[index[1]] && str1[index[1]] == str2[index[0]]
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
	return &unionFind{parent, rank, n}
}

func (u *unionFind) isConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *unionFind) find(q int) int {
	if u.parent[q] != q {
		u.parent[q] = u.find(u.parent[q])
	}
	return u.parent[q]
}

func (u *unionFind) union(p, q int) {
	pRoot, qRoot := u.find(p), u.find(q)
	if pRoot == qRoot {
		return
	}

	if u.rank[pRoot] < u.rank[qRoot] {
		pRoot, qRoot = qRoot, pRoot
	}

	u.rank[pRoot] += u.rank[qRoot]
	u.parent[qRoot] = pRoot
	u.count--
}
