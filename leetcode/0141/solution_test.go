package _141

import (
	"encoding/json"
	"testing"
)

func Test_hasCycle(t *testing.T) {
	type args struct {
		input string
		pos   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testCase 1",
			args: args{
				input: "[3, 2, 0, -4]",
				pos:   1,
			},
			want: true,
		},
		{
			name: "testCase 2",
			args: args{
				input: "[1,2]",
				pos:   0,
			},
			want: true,
		},
		{
			name: "testCase 3",
			args: args{
				input: "[1]",
				pos:   -1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		var in []int
		_ = json.Unmarshal([]byte(tt.args.input), &in)
		var head *ListNode
		var cur *ListNode
		var p *ListNode
		for idx, i := range in {
			node := &ListNode{Val: i}
			if idx == tt.args.pos {
				p = node
			}
			if head == nil {
				head = node
				cur = node
			} else {
				cur.Next = node
				cur = cur.Next
			}
		}
		cur.Next = p
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
