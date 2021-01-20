package _0628

import "testing"

func Test_maximumProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testCase 1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 6,
		},
		{
			name: "testCase 2",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 24,
		},
		{
			name: "testCase 3",
			args: args{
				nums: []int{-3, -5, 1, 2, 3, 4},
			},
			want: 60,
		},
		{
			name: "testCase 4",
			args: args{
				nums: []int{-3, -1, 1},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProduct(tt.args.nums); got != tt.want {
				t.Errorf("maximumProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
