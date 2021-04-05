package dualpoint

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

// HasCycle use the fast and slow point to determine whether a link has cycle
func HasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// DetectCycle use the fast and slow point to detect the entry of a cycle
func DetectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			break
		}
	}
	// can not find a cycle
	if fast == nil || fast.Next == nil {
		return nil
	}
	// let slow point repoint to the head
	slow = head
	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

// FindMiddleNode returns the middle node
func FindMiddleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	fast, slow := head, head
	for n != 0 {
		fast = fast.Next
		n--
	}
	if fast == nil {
		return head.Next
	}
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}
