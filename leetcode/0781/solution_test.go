package _781

import "testing"

func Test_numRabbits(t *testing.T) {
	type args struct {
		answers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testCase 1",
			args: args{
				answers: []int{1, 1, 2},
			},
			want: 5,
		},
		{
			name: "testCase 2",
			args: args{
				answers: []int{10, 10, 10},
			},
			want: 11,
		},
		{
			name: "testCase 3",
			args: args{
				answers: []int{},
			},
			want: 0,
		},
		{
			name: "testCase 4",
			args: args{
				answers: []int{1,0,1,0,0},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRabbits(tt.args.answers); got != tt.want {
				t.Errorf("numRabbits() = %v, want %v", got, tt.want)
			}
		})
	}
}
