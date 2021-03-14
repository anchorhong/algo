package _1631

import "testing"

func Test_minimumEffortPath(t *testing.T) {
	type args struct {
		heights [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testCase 1",
			args: args{
				heights: [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}},
			},
			want: 2,
		},
		{
			name: "testCase 2",
			args: args{
				heights: [][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}},
			},
			want: 1,
		},
		{
			name: "testCase 3",
			args: args{
				heights: [][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}},
			},
			want: 0,
		},
		{
			name: "testCase 4",
			args: args{
				heights: [][]int{{3}},
			},
			want: 0,
		},
		{
			name: "testCase 5",
			args: args{
				heights: [][]int{{8, 3, 2, 5, 2, 10, 7, 1, 8, 9}, {1, 4, 9, 1, 10, 2, 4, 10, 3, 5}, {4, 10, 10, 3, 6, 1, 3, 9, 8, 8}, {4, 4, 6, 10, 10, 10, 2, 10, 8, 8}, {9, 10, 2, 4, 1, 2, 2, 6, 5, 7}, {2, 9, 2, 6, 1, 4, 7, 6, 10, 9}, {8, 8, 2, 10, 8, 2, 3, 9, 5, 3}, {2, 10, 9, 3, 5, 1, 7, 4, 5, 6}, {2, 3, 9, 2, 5, 10, 2, 7, 1, 8}, {9, 10, 4, 10, 7, 4, 9, 3, 1, 6}},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumEffortPath(tt.args.heights); got != tt.want {
				t.Errorf("minimumEffortPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
