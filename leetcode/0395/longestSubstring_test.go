package _395

import "testing"

func Test_longestSubstring(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testCase 1",
			args: args{
				s: "aaabb",
				k: 3,
			},
			want: 3,
		},
		{
			name: "testCase 2",
			args: args{
				s: "ababbc",
				k: 2,
			},
			want: 5,
		},
		{
			name: "testCase 3",
			args: args{
				s: "ababacb",
				k: 3,
			},
			want: 0,
		},
		{
			name: "testCase 4",
			args: args{
				s: "bbaaacddcaabdbd",
				k: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubstring(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("longestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
