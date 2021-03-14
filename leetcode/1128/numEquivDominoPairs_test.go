package _1128

import "testing"

func Test_numEquivDominoPairs(t *testing.T) {
	type args struct {
		dominoes [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testCase 1",
			args: args{
				dominoes: [][]int{{2, 1}, {5, 4}, {3, 7}, {6, 2}, {4, 4}, {1, 8}, {9, 6}, {5, 3}, {1, 9}, {1, 1}, {6, 6}, {9, 6}, {1, 3}, {9, 7}, {4, 7}, {5, 1}},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numEquivDominoPairs(tt.args.dominoes); got != tt.want {
				t.Errorf("numEquivDominoPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
