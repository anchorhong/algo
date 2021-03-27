package _173

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	var stack []*TreeNode
	cur := root
	for cur != nil {
		stack = append(stack, cur)
		cur = cur.Left
	}
	return BSTIterator{stack: stack}
}

func (b *BSTIterator) Next() int {
	top := b.stack[len(b.stack)-1]
	b.stack = b.stack[:len(b.stack)-1]
	cur := top.Right
	for cur != nil {
		b.stack = append(b.stack, cur)
		cur = cur.Left
	}
	return top.Val
}

func (b *BSTIterator) HasNext() bool {
	return len(b.stack) > 0
}
