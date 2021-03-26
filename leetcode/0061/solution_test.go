package _061

import (
	"reflect"
	"testing"
)

func Test_rotateRight(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		k     int
		want  []int
	}{
		{
			name:  "testCase 1",
			input: []int{1, 2, 3, 4, 5},
			k:     2,
			want:  []int{4, 5, 1, 2, 3},
		},
		{
			name:  "testCase 2",
			input: []int{0, 1, 2},
			k:     4,
			want:  []int{2, 0, 1},
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
		got := rotateRight(head, tt.k)
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
