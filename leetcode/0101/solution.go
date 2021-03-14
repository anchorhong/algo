package _101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}
func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

func isSymmetric1(root *TreeNode) bool {
	queue := []*TreeNode{root, root}
	for len(queue) > 0 {
		p, q := queue[0], queue[1]
		queue = queue[2:]
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		queue = append(queue, p.Left)
		queue = append(queue, q.Right)
		queue = append(queue, p.Right)
		queue = append(queue, q.Left)
	}
	return true
}
