package _1006

import "testing"

func Test_clumsy(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{
			name: "testCase 1",
			N:    4,
			want: 7,
		},
		{
			name: "testCase 2",
			N:    10,
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clumsy(tt.N); got != tt.want {
				t.Errorf("clumsy() = %v, want %v", got, tt.want)
			}
		})
	}
}
