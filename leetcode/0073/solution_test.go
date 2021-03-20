package _073

import (
	"reflect"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name:   "testCase 1",
			matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			want:   [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.matrix)
			if !reflect.DeepEqual(tt.matrix, tt.want) {
				t.Errorf("result is:\n %v \n but want:\n %v", tt.matrix, tt.want)
			}
		})
	}
}
