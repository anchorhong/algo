package _0803

func hitBricks(grid [][]int, hits [][]int) []int {
	m, n := len(grid), len(grid[0])
	gridCopy := make([][]int, m)
	parents := make([]int, m*n)
	weights := make([]int, m*n)
	var count int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			gridCopy[i] = append(gridCopy[i], grid[i][j])
			if grid[i][j] == 1 {
				parents[i*m+j] = i*m + j
				weights[i*m+j] = 1
				count++
			} else {
				parents[i*m+j] = -1
				weights[i*m+j] = 1
			}
		}
	}
	for _, hit := range hits {
		x, y := hit[0], hit[1]
		if grid[x][y] == 1 {
			count--
		}
		gridCopy[x][y] = 0
		parents[x*m+y] = -1
		weights[x*m+y] = -1
	}
	uf := unionFind{
		parents: parents,
		weights: weights,
		count:   count,
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if gridCopy[i][j] == 1 {
				if i > 0 && gridCopy[i-1][j] == 1 {
					uf.union(i*m+j, (i-1)*m+j)
				}
				if j > 0 && gridCopy[i][j-1] == 1 {
					uf.union(i*m+j, i*m+j-1)
				}
			}
		}
	}
	var res []int
	for i := len(hits) - 1; i >= 0; i-- {
		x, y := hits[i][0], hits[i][1]
		if grid[x][y] == 0 {
			res = append(res, 0)
		} else {
			uf.parents[x*m+y] = x*m + y
			uf.count++
			preCount := uf.count
			// upper
			if x > 0 && gridCopy[x-1][y] == 1 {
				uf.union(x*m+y, (x-1)*m+y)
			}
			// down
			if x < m-1 && gridCopy[x+1][y] == 1 {
				uf.union(x*m+y, (x+1)*m+y)
			}
			// left
			if y > 0 && gridCopy[x][y-1] == 1 {
				uf.union(x*m+y, x*m+y-1)
			}
			// right
			if y < n-1 && gridCopy[x][y+1] == 1 {
				uf.union(x*m+y, x*m+y+1)
			}
			postCount := uf.count
			res = append(res, postCount-preCount-1)
		}
	}
	var reversedRes []int
	for i := len(res) - 1; i >= 0; i-- {
		reversedRes = append(reversedRes, res[i])
	}
	return reversedRes
}

type unionFind struct {
	parents []int
	weights []int
	count   int
}

func (u unionFind) find(q int) int {
	if u.parents[q] != q {
		u.parents[q] = u.find(u.parents[q])
	}
	return u.parents[q]
}

func (u unionFind) union(p, q int) {
	pRoot := u.find(p)
	qRoot := u.find(q)

	if pRoot == qRoot {
		return
	}

	if u.weights[pRoot] < u.weights[qRoot] {
		u.parents[qRoot] = qRoot
		u.weights[pRoot] += u.weights[qRoot]
	} else {
		u.parents[pRoot] = qRoot
		u.weights[pRoot] += u.weights[qRoot]
	}
	u.count--
}
