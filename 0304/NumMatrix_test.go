package _304

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		cases  [][]int
		want   []int
	}{
		//{
		//	name: "testCase 1",
		//	matrix: [][]int{{3, 0, 1, 4, 2},
		//		{5, 6, 3, 2, 1},
		//		{1, 2, 0, 1, 5},
		//		{4, 1, 0, 1, 7},
		//		{1, 0, 3, 0, 5},
		//	},
		//	cases: [][]int{
		//		{2, 1, 4, 3},
		//		{1, 1, 2, 2},
		//		{1, 2, 2, 4},
		//	},
		//	want: []int{8, 11, 12},
		//},
		{
			name: "testCase 2",

			matrix: [][]int{},
			cases:  [][]int{},

			want: []int{},
		},
	}
	for _, tt := range tests {
		numMatrix := Constructor(tt.matrix)
		t.Run(tt.name, func(t *testing.T) {
			for i, got := range tt.cases {
				if result := numMatrix.SumRegion(got[0], got[1], got[2], got[3]); result != tt.want[i] {
					t.Errorf("SumRegion() = %v, want %v", result, tt.want[i])
				}
			}
		})
	}
}
