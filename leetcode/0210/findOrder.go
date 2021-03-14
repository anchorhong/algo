package _0210

func findOrderByKahn(numCourses int, prerequisites [][]int) []int {
	inDegress := make([]int, numCourses)
	adj := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		adj[i] = make([]int, 0)
		inDegress[i] = 0
	}

	for _, pre := range prerequisites {
		adj[pre[1]] = append(adj[pre[1]], pre[0])
		inDegress[pre[0]]++
	}
	var queue []int
	var res []int
	for i := 0; i < numCourses; i++ {
		if inDegress[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		res = append(res, cur)
		for j := 0; j < len(adj[cur]); j++ {
			w := adj[cur][j]
			inDegress[w]--
			if inDegress[w] == 0 {
				queue = append(queue, w)
			}
		}
	}
	if len(res) == numCourses {
		return res
	} else {
		return []int{}
	}
}
