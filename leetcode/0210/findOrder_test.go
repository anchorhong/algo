package _0210

import (
	"reflect"
	"testing"
)

func Test_findOrderByKahn(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "testCase 1",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
			},
			want: [][]int{{0, 1}},
		},
		{
			name: "testCase 2",
			args: args{
				numCourses:    4,
				prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			},
			want: [][]int{{0, 1, 2, 3}, {0, 2, 1, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findOrderByKahn(tt.args.numCourses, tt.args.prerequisites)
			for _, w := range tt.want{
				if reflect.DeepEqual(got, w){
					return
				}
			}
			t.Errorf("findOrderByKahn() = %v, want %v", got, tt.want)
		})
	}
}
