package _115

import "testing"

func Test_numDistinct(t *testing.T) {

	tests := []struct {
		name string
		s    string
		t    string
		want int
	}{
		{
			name: "testCase 1",
			s:    "rabbbit",
			t:    "rabbit",
			want: 3,
		},
		{
			name: "testCase 2",
			s:    "babgbag",
			t:    "bag",
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDistinct(tt.s, tt.t); got != tt.want {
				t.Errorf("numDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
