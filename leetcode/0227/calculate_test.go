package _227

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "testCase 1",
			args: args{
				s: "3+2*2",
			},
			wantAns: 7,
		},
		{
			name: "testCase 2",
			args: args{
				s: " 3/2 ",
			},
			wantAns: 1,
		},
		{
			name: "testCase 3",
			args: args{
				s: " 3+5 / 2 ",
			},
			wantAns: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := calculate(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("calculate() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
