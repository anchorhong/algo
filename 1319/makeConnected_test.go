package _1319

import "testing"

func Test_makeConnected(t *testing.T) {
	type args struct {
		n           int
		connections [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testCase 1",
			args: args{
				n:           4,
				connections: [][]int{{0, 1}, {0, 2}, {1, 2}},
			},
			want: 1,
		},
		{
			name: "testCase 2",
			args: args{
				n:           6,
				connections: [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}},
			},
			want: 2,
		},
		{
			name: "testCase 3",
			args: args{
				n:           6,
				connections: [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}},
			},
			want: -1,
		},
		{
			name: "testCase 4",
			args: args{
				n:           5,
				connections: [][]int{{0, 1}, {0, 2}, {3, 4}, {2, 3}},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeConnected(tt.args.n, tt.args.connections); got != tt.want {
				t.Errorf("makeConnected() = %v, want %v", got, tt.want)
			}
		})
	}
}
