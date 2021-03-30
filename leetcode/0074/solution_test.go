package _074

import "testing"

func Test_searchMatrix(t *testing.T) {

	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{
			name:   "testCase 1",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}},
			target: 3,
			want:   true,
		},
		{
			name:   "testCase 2",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}},
			target: 13,
			want:   false,
		},
		{
			name:   "testCase 3",
			matrix: [][]int{{1}},
			target: 2,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.matrix, tt.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchMatrix1(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{
			name:   "testCase 1",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}},
			target: 3,
			want:   true,
		},
		{
			name:   "testCase 2",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}},
			target: 13,
			want:   false,
		},
		{
			name:   "testCase 3",
			matrix: [][]int{{1}},
			target: 2,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix1(tt.matrix, tt.target); got != tt.want {
				t.Errorf("searchMatrix1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchMatrix2(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{
			name:   "testCase 1",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}},
			target: 3,
			want:   true,
		},
		{
			name:   "testCase 2",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}},
			target: 13,
			want:   false,
		},
		{
			name:   "testCase 3",
			matrix: [][]int{{1}},
			target: 2,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix2(tt.matrix, tt.target); got != tt.want {
				t.Errorf("searchMatrix2() = %v, want %v", got, tt.want)
			}
		})
	}
}