package _143

import "testing"

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "textCase 1",
			args: args{
				text1: "abcde",
				text2: "ace",
			},
			want: 3,
		},
		{
			name: "textCase 2",
			args: args{
				text1: "abc",
				text2: "abc",
			},
			want: 3,
		},
		{
			name: "textCase 3",
			args: args{
				text1: "abc",
				text2: "def",
			},
			want: 0,
		},
		{
			name: "textCase 4",
			args: args{
				text1: "bsbininm",
				text2: "jmjkbkjkv",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
