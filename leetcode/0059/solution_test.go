package _059

import (
	"reflect"
	"testing"
)

func Test_generateMatrix(t *testing.T) {

	tests := []struct {
		name string
		n    int
		want [][]int
	}{
		{
			name: "testCase 1",
			n:    3,
			want: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
		},
		{
			name: "testCase 2",
			n:    1,
			want: [][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMatrix(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
