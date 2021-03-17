package _092

import (
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {

	tests := []struct {
		name  string
		nodes []int
		left  int
		right int
		want  []int
	}{
		{
			name:  "testCase 1",
			nodes: []int{1, 2, 3, 4, 5},
			left:  2,
			right: 4,
			want:  []int{1, 4, 3, 2, 5},
		},
		{
			name:  "testCase 2",
			nodes: []int{1, 2, 3, 4},
			left:  2,
			right: 4,
			want:  []int{1, 4, 3, 2},
		},
		{
			name:  "testCase 3",
			nodes: []int{1, 2, 3, 4},
			left:  1,
			right: 1,
			want:  []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var head *ListNode
			var pre *ListNode
			for i := 0; i < len(tt.nodes); i++ {
				if head == nil {
					head = &ListNode{Val: tt.nodes[i]}
					pre = head
				} else {
					pre.Next = &ListNode{Val: tt.nodes[i]}
					pre = pre.Next
				}
			}
			var got []int
			cur := reverseBetween(head, tt.left, tt.right)
			for cur != nil {
				got = append(got, cur.Val)
				cur = cur.Next
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
