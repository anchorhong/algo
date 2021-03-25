package _083

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
			input: []int{1, 1, 2},
			want:  []int{1, 2},
		},
		{
			name:  "testCase 2",
			input: []int{1, 1, 2, 3, 3},
			want:  []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		head := &ListNode{Val: tt.input[0]}
		cur := head
		for i := 1; i < len(tt.input); i++ {
			tmp := &ListNode{Val: tt.input[i]}
			cur.Next = tmp
			cur = cur.Next
		}
		var res []int
		got := deleteDuplicates(head)
		for got != nil {
			res = append(res, got.Val)
			got = got.Next
		}
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", res, tt.want)
			}
		})
	}
}
