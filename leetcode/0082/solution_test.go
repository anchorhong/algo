package _082

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {

	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "testCase 1",
			input: []int{1, 2, 3, 3, 4, 4, 5},
			want:  []int{1, 2, 5},
		},
		{
			name:  "testCase 2",
			input: []int{1, 1, 1, 2, 3},
			want:  []int{2, 3},
		},
		{
			name:  "testCase 3",
			input: []int{1, 1, 1, 1, 1},
			want:  nil,
		},
		{
			name:  "testCase 4",
			input: []int{1, 2, 3, 4, 4},
			want:  []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		head := &ListNode{Val: tt.input[0]}
		cur := head
		for i := 1; i < len(tt.input); i++ {
			l := &ListNode{Val: tt.input[i]}
			cur.Next = l
			cur = cur.Next
		}
		t.Run(tt.name, func(t *testing.T) {
			var res []int
			cur := deleteDuplicates(head)
			for cur != nil {
				res = append(res, cur.Val)
				cur = cur.Next
			}
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", res, tt.want)
			}
		})
	}
}
