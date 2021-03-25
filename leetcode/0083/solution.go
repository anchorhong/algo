package _083

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		for cur.Next != nil && cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return head
}
