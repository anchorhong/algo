package _092

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	flag := &ListNode{Next: head}
	pre, cur := flag, head
	var tail *ListNode
	var post *ListNode
	for i := 1; i <= right; i++ {
		if i == left {
			tail = cur
			post = cur
			cur = cur.Next
			post.Next = nil
		} else if i > left {
			newCur := cur.Next
			cur.Next = post
			post = cur
			cur = newCur
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	pre.Next = post
	if cur != nil {
		tail.Next = cur
	}
	return flag.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
