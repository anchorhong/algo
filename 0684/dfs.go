package _684

import "fmt"

func findRedundantConnectionByDFS(edges [][]int) []int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	adjs := make([][]int, len(edges)+1)
	for _, edge := range edges {
		adjs[edge[0]] = append(adjs[edge[0]], edge[1])
		adjs[edge[1]] = append(adjs[edge[1]], edge[0])
	}
	visited := make(map[int]int)
	var path []int
	var dfs func(int, int) bool
	dfs = func(id, parent int) bool {
		for _, adj := range adjs[id] {
			if adj != parent {
				if j, ok := visited[adj]; ok {
					path = path[j:]
					return true
				}
				path = append(path, adj)
				visited[adj] = len(path) - 1
				if dfs(adj, id) {
					return true
				}
				path = (path)[:len(path)-1]
			}
		}
		return false
	}

	visited[edges[0][0]] = 0
	path = append(path, edges[0][0])
	dfs(edges[0][0], 0)
	resMap := make(map[string]interface{})
	if len(path) > 1 {
		for i := 0; i < len(path); i++ {
			id1 := path[i]
			var id2 int
			if i < len(path)-1 {
				id2 = path[i+1]
			} else {
				id2 = path[0]
			}
			p := fmt.Sprintf("%v->%v", min(id1, id2), max(id1, id2))
			resMap[p] = nil
		}

	}
	for i := len(edges) - 1; i >= 0; i-- {
		edge := edges[i]
		p := fmt.Sprintf("%v->%v", edge[0], edge[1])
		if _, ok := resMap[p]; ok {
			return edge
		}
	}
	return nil
}
