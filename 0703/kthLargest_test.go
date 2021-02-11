package _703

import "testing"

func TestKthLargest_Add(t *testing.T) {
	tests := []struct {
		name string
		kTh  int
		nums []int
		adds []int
		want []int
	}{
		{
			name: "testCase 1",
			kTh:  3,
			nums: []int{4, 5, 8, 2},
			adds: []int{3, 5, 10, 9, 4},
			want: []int{4, 5, 5, 8, 8},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			kthLargest := Constructor(tt.kTh, tt.nums)
			for i, v := range tt.adds {
				if got := kthLargest.Add(v); got != tt.want[i] {
					t.Errorf("Add() = %v, want %v", got, tt.want[i])
				}
			}
		})
	}
}
