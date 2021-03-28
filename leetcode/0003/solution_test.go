package _003

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "testCase 1",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "testCase 2",
			s:    "bbbbb",
			want: 1,
		},
		{
			name: "testCase 3",
			s:    "pwwkew",
			want: 3,
		},
		{
			name: "testCase 4",
			s:    "",
			want: 0,
		},
		{
			name: "testCase 5",
			s:    "qrsvbspk",
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
