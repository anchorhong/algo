package _888

import (
	"reflect"
	"testing"
)

func Test_fairCandySwap(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "testCase 1",
			args: args{
				A: []int{1, 1},
				B: []int{2, 2},
			},
			want: []int{1, 2},
		},
		{
			name: "testCase 2",
			args: args{
				A: []int{1, 2},
				B: []int{2, 3},
			},
			want: []int{1, 2},
		},
		{
			name: "testCase 3",
			args: args{
				A: []int{2},
				B: []int{1, 3},
			},
			want: []int{2, 3},
		},
		{
			name: "testCase 4",
			args: args{
				A: []int{1, 2, 5},
				B: []int{2, 4},
			},
			want: []int{5, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fairCandySwap(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fairCandySwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
