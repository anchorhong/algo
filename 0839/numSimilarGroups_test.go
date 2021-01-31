package _839

import "testing"

func Test_numSimilarGroups(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testCase 1",
			args: args{
				strs: []string{"tars", "rats", "arts", "star"},
			},
			want: 2,
		},
		{
			name: "testCase 2",
			args: args{
				strs: []string{"omv", "ovm"},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSimilarGroups(tt.args.strs); got != tt.want {
				t.Errorf("numSimilarGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
