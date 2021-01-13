package _399

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	nodeMapping := make(map[string][]string)
	valueMap := make(map[string]float64)
	for i, equation := range equations {
		node1 := equation[0]
		node2 := equation[1]
		nodeMapping[node1] = append(nodeMapping[node1], node2)
		nodeMapping[node2] = append(nodeMapping[node2], node1)
		valueMap[node1+node2] = values[i]
		valueMap[node2+node1] = 1.0 / values[i]
	}
	var result []float64
	for _, query := range queries {
		op1 := query[0]
		op2 := query[1]
		var path []string
		visisted := make(map[string]bool)
		visisted[op1] = true
		path = append(path, op1)
		found := dfs(op1, op2, nodeMapping, visisted, &path)
		res := -1.0
		if found {
			res = 1.0
			for i := 0; i < len(path)-1; i++ {
				relation := (path)[i] + (path)[i+1]
				if v, ok := valueMap[relation]; ok {
					res *= v
				} else {
					res = -1.0
					break
				}
			}
		}
		result = append(result, res)
	}
	return result
}

func dfs(op1, dst string, mapping map[string][]string, visisted map[string]bool, path *[]string) bool {
	if _, ok := mapping[op1]; !ok {
		return false
	}
	if op1 == dst {
		return true
	}
	if nodes, ok := mapping[op1]; ok {
		for _, node := range nodes {
			if _, ok := visisted[node]; !ok {
				*path = append(*path, node)
				visisted[node] = true
				found := dfs(node, dst, mapping, visisted, path)
				if found {
					return found
				}
				*path = (*path)[:len(*path)-1]
			}
		}
	}
	return false
}
