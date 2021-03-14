package _0207

func canFinishByKahn(numCourses int, prerequisites [][]int) bool {
	inDegress := make([]int, numCourses)
	adjs := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		adjs[i] = make([]int, 0)
	}
	for _, pre := range prerequisites {
		s := pre[0]
		t := pre[1]
		inDegress[s]++
		adjs[t] = append(adjs[t], s)
	}
	var res []int
	var queue []int
	for i := 0; i < numCourses; i++ {
		if inDegress[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		cur := queue[0]
		res = append(res, cur)
		queue = queue[1:]
		for j := 0; j < len(adjs[cur]); j++ {
			w := adjs[cur][j]
			inDegress[w]--
			if inDegress[w] == 0 {
				queue = append(queue, w)
			}
		}
	}
	return len(res) == numCourses
}

func canFinishByDFS(numCourses int, prerequisites [][]int) bool {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  []int
		valid   = true
		dfs     func(int)
	)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}
