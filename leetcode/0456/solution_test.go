package _456

import "testing"

func Test_find132pattern(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "testCase 1",
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			name: "testCase 2",
			nums: []int{3, 1, 4, 2},
			want: true,
		},
		{
			name: "testCase 3",
			nums: []int{-1, 3, 2, 0},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find132pattern(tt.nums); got != tt.want {
				t.Errorf("find132pattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
