package _309

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
			name: "testCase 1",
			args: args{
				prices: []int{1, 2, 3, 0, 2},
			},
			want: 3,
		},
		{
			name: "testCase 2",
			args: args{
				prices: []int{2, 1, 4},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfitOpt(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "testCase 1",
			args: args{
				prices: []int{1, 2, 3, 0, 2},
			},
			want: 3,
		},
		{
			name: "testCase 2",
			args: args{
				prices: []int{2, 1, 4},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitOpt(tt.args.prices); got != tt.want {
				t.Errorf("maxProfitOpt() = %v, want %v", got, tt.want)
			}
		})
	}
}