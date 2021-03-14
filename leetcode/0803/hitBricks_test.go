package _0803

import (
	"reflect"
	"testing"
)

func Test_hitBricks(t *testing.T) {
	type args struct {
		grid [][]int
		hits [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "testCase 1",
			args: args{
				grid: [][]int{{1, 0, 0, 0}, {1, 1, 1, 0}},
				hits: [][]int{{1, 0}},
			},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hitBricks(tt.args.grid, tt.args.hits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hitBricks() = %v, want %v", got, tt.want)
			}
		})
	}
}
