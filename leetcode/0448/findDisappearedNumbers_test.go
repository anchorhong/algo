package _448

import (
	"reflect"
	"testing"
)

func Test_findDisappearedNumbers(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "testCase 1",
			nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			want: []int{5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDisappearedNumbers(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDisappearedNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
