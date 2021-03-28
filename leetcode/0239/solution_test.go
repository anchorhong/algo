package _239

import (
	"reflect"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "testCase 1",
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			want: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name: "testCase 2",
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
		{
			name: "testCase 3",
			nums: []int{1, -1},
			k:    1,
			want: []int{1, -1},
		},
		{
			name: "testCase 4",
			nums: []int{9, 11},
			k:    2,
			want: []int{11},
		},
		{
			name: "testCase 5",
			nums: []int{4, -2},
			k:    2,
			want: []int{4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow(tt.nums, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSlidingWindow1(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "testCase 1",
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			want: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name: "testCase 2",
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
		{
			name: "testCase 3",
			nums: []int{1, -1},
			k:    1,
			want: []int{1, -1},
		},
		{
			name: "testCase 4",
			nums: []int{9, 11},
			k:    2,
			want: []int{11},
		},
		{
			name: "testCase 5",
			nums: []int{4, -2},
			k:    2,
			want: []int{4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow1(tt.nums, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow1() = %v, want %v", got, tt.want)
			}
		})
	}
}