package _331

func isValidSerialization(preorder string) bool {
	stack := []int{1}
	n := len(preorder)
	for i := 0; i < len(preorder); {
		if len(stack) == 0 {
			return false
		}
		if preorder[i] == ',' {
			i++
		} else if preorder[i] == '#' {
			stack[len(stack)-1]--
			if stack[len(stack)-1] == 0 {
				stack = stack[:len(stack)-1]
			}
			i++
		} else {
			for i < n && preorder[i] != ',' {
				i++
			}
			stack[len(stack)-1]--
			if stack[len(stack)-1] == 0 {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, 2)
		}
	}
	return len(stack) == 0
}