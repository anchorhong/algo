package _082

type ListNode struct {
	Val  int
	Next *ListNode
}


func deleteDuplicates(head *ListNode) *ListNode {
	flag := &ListNode{Next: head}
	cur := flag
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			r := cur.Next.Val
			t := cur.Next
			for t != nil && t.Val == r {
				t = t.Next
			}
			cur.Next = t
		} else {
			cur = cur.Next
		}
	}
	return flag.Next
}
