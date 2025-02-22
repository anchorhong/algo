package _224

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "testCase 1",
			s:    "1 + 1",
			want: 2,
		},
		{
			name: "testCase 2",
			s:    " 2-1 + 2 ",
			want: 3,
		},
		{
			name: "testCase 3",
			s:    "(1+(4+5+2)-3)+(6+8)",
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
