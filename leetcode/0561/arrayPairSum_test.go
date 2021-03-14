package _561

import "testing"

func Test_arrayPairSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "testCase 1",
			nums: []int{1, 4, 3, 2},
			want: 4,
		},
		{
			name: "testCase 2",
			nums: []int{6, 2, 6, 5, 1, 2},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayPairSum(tt.nums); got != tt.want {
				t.Errorf("arrayPairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
