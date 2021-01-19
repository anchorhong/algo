package _1584

import "testing"

func Test_minCostConnectPoints(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testCase 1",
			args: args{
				points: [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}},
			},
			want: 20,
		},
		{
			name: "testCase 2",
			args: args{
				points: [][]int{{3, 12}, {-2, 5}, {-4, 1}},
			},
			want: 18,
		},
		{
			name: "testCase 3",
			args: args{
				points: [][]int{{0, 0}, {1, 1}, {1, 0}, {-1, 1}},
			},
			want: 4,
		},
		{
			name: "testCase 4",
			args: args{
				points: [][]int{{-1000000, -1000000}, {1000000, 1000000}},
			},
			want: 4000000,
		},
		{
			name: "testCase 5",
			args: args{
				points: [][]int{{0, 0}},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostConnectPoints(tt.args.points); got != tt.want {
				t.Errorf("minCostConnectPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
