package _188

import "testing"

func Test_maxProfitByDynamic(t *testing.T) {
	type args struct {
		k      int
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
				k:      2,
				prices: []int{2, 4, 1},
			},
			want: 2,
		},
		{
			name: "testCase 2",
			args: args{
				k:      2,
				prices: []int{3, 2, 6, 5, 0, 3},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitByDynamic(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfitByDynamic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfitByDynamicOpt(t *testing.T) {
	type args struct {
		k      int
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
				k:      2,
				prices: []int{2, 4, 1},
			},
			want: 2,
		},
		{
			name: "testCase 2",
			args: args{
				k:      2,
				prices: []int{3, 2, 6, 5, 0, 3},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitByDynamicOpt(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfitByDynamicOpt() = %v, want %v", got, tt.want)
			}
		})
	}
}