package _132

import "testing"

func Test_minCut(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testCase 1",
			args: args{
				s: "aab",
			},
			want: 1,
		},
		{
			name: "testCase 2",
			args: args{
				s: "a",
			},
			want: 0,
		},
		{
			name: "testCase 3",
			args: args{
				s: "ab",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCut(tt.args.s); got != tt.want {
				t.Errorf("minCut() = %v, want %v", got, tt.want)
			}
		})
	}
}
