package _1438

import "testing"

func Test_longestSubarray(t *testing.T) {

	tests := []struct {
		name  string
		nums  []int
		limit int
		want  int
	}{
		{
			name:  "testCase 1",
			nums:  []int{8, 2, 4, 7},
			limit: 4,
			want:  2,
		},
		{
			name:  "testCase 2",
			nums:  []int{10, 1, 2, 4, 7, 2},
			limit: 5,
			want:  4,
		},
		{
			name:  "testCase 3",
			nums:  []int{4, 2, 2, 2, 4, 4, 2, 2},
			limit: 0,
			want:  3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.nums, tt.limit); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
