package _503

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElements(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "testCase 1",
			nums: []int{1, 2, 1},
			want: []int{2, -1, 2},
		},
		{
			name: "testCase 2",
			nums: []int{1, 1, 5, 1, 4},
			want: []int{5, 5, -1, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElements(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
