package _329

func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	edges := make([][]int, n*m)
	inDegrees := make([]int, n*m)
	node2Value := make([]int, n*m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i-1 >= 0 && matrix[i-1][j] < matrix[i][j] {
				edges[(i-1)*n+j] = append(edges[(i-1)*n+j], i*n+j)
				inDegrees[i*n+j]++
			}
			if i+1 < m && matrix[i+1][j] < matrix[i][j] {
				edges[(i+1)*n+j] = append(edges[(i+1)*n+j], i*n+j)
				inDegrees[i*n+j]++
			}
			if j-1 >= 0 && matrix[i][j-1] < matrix[i][j] {
				edges[i*n+j-1] = append(edges[i*n+j-1], i*n+j)
				inDegrees[i*n+j]++
			}
			if j+1 < n && matrix[i][j+1] < matrix[i][j] {
				edges[i*n+j+1] = append(edges[i*n+j+1], i*n+j)
				inDegrees[i*n+j]++
			}
			node2Value[i*n+j] = matrix[i][j]
		}
	}
	maxPath := 1
	memo := make(map[int]int)
	var dfs func(int) int
	dfs = func(cur int) int {
		if v, ok := memo[cur]; ok {
			return v
		}
		max := 1
		for _, edge := range edges[cur] {
			if node2Value[edge] > node2Value[cur] {
				res := dfs(edge) + 1
				if res > max {
					max = res
				}
			}
		}
		memo[cur] = max
		return max
	}

	for i, inDegree := range inDegrees {
		if inDegree != 0 {
			continue
		}
		res := dfs(i)
		if res > maxPath {
			maxPath = res
		}
	}
	return maxPath
}
