package _0207

import "testing"

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testCase 1",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
			},
			want: true,
		},
		{
			name: "testCase 2",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}, {0, 1}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinishByKahn(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinishByKahn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canFinishByDFS(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testCase 1",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
			},
			want: true,
		},
		{
			name: "testCase 2",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}, {0, 1}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinishByDFS(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinishByDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
