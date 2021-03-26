package _061

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	n := 0
	cur := dummy.Next
	tail := cur
	for cur != nil {
		n++
		tail = cur
		cur = cur.Next
	}
	k = k % n
	if k == 0 {
		return head
	}
	cur = dummy
	for i := 0; i < n-k; i++ {
		cur = cur.Next
	}
	tmpNext := dummy.Next
	dummy.Next = cur.Next
	cur.Next = nil
	tail.Next = tmpNext
	return dummy.Next
}
