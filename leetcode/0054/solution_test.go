package _054

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{
			name:   "testCase 1",
			matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			want:   []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name:   "testCase 2",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name:   "testCase 3",
			matrix: [][]int{{1}},
			want:   []int{1},
		},
		{
			name:   "testCase 4",
			matrix: [][]int{{1}, {2}, {3}, {4}},
			want:   []int{1, 2, 3, 4},
		},
		{
			name:   "testCase 4",
			matrix: [][]int{{2, 5, 8}, {4, 0, -1}},
			want:   []int{2, 5, 8, -1, 0, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrder(tt.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
