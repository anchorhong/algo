package _1202

import "testing"

func Test_smallestStringWithSwaps(t *testing.T) {
	type args struct {
		s     string
		pairs [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testCase 1",
			args: args{
				s:     "dcab",
				pairs: [][]int{{0, 3}, {1, 2}},
			},
			want: "bacd",
		},
		{
			name: "testCase 2",
			args: args{
				s:     "dcab",
				pairs: [][]int{{0, 3}, {1, 2}, {0, 2}},
			},
			want: "abcd",
		},
		{
			name: "testCase 3",
			args: args{
				s:     "cba",
				pairs: [][]int{{0, 1}, {1, 2}},
			},
			want: "abc",
		},
		{
			name: "testCase 4",
			args: args{
				s:     "udyyek",
				pairs: [][]int{{3, 3}, {3, 0}, {5, 1}, {3, 1}, {3, 4}, {3, 5}},
			},
			want: "deykuy",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestStringWithSwaps(tt.args.s, tt.args.pairs); got != tt.want {
				t.Errorf("smallestStringWithSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
