package _989

import (
	"reflect"
	"testing"
)

func Test_addToArrayForm(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "testCase 1",
			args: args{
				A: []int{1, 2, 0, 0},
				K: 34,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "testCase 2",
			args: args{
				A: []int{2, 7, 4},
				K: 181,
			},
			want: []int{4, 5, 5},
		},
		{
			name: "testCase 3",
			args: args{
				A: []int{2, 1, 5},
				K: 806,
			},
			want: []int{1, 0, 2, 1},
		},
		{
			name: "testCase 4",
			args: args{
				A: []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
				K: 1,
			},
			want: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addToArrayForm(tt.args.A, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addToArrayForm() = %v, want %v", got, tt.want)
			}
		})
	}
}
