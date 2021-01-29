package _123

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{prices: []int{3, 3, 5, 0, 0, 3, 1, 4}},
			want: 6,
		},
		{
			name: "case2",
			args: args{prices: []int{1, 2, 3, 4, 5}},
			want: 4,
		},
		{
			name: "case3",
			args: args{prices: []int{7, 6, 4, 3, 1}},
			want: 0,
		},
		{
			name: "case4",
			args: args{prices: []int{1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitByDynamic(tt.args.prices); got != tt.want {
				t.Errorf("maxProfitByDynamic() = %v, want %v", got, tt.want)
			}
		})
	}
}
