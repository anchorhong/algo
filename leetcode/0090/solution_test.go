package _090

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func Test_subsetsWithDup(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want string
	}{
		{
			name: "testCase 1",
			nums: []int{1, 2, 2},
			want: "[[],[1],[1,2],[1,2,2],[2],[2,2]]",
		},
		{
			name: "testCase 2",
			nums: []int{0},
			want: "[[],[0]]",
		},
	}
	for _, tt := range tests {
		var wantSlice [][]int
		_ = json.Unmarshal([]byte(tt.want), &wantSlice)
		wantMap := make(map[string]bool)
		for _, x := range wantSlice {
			k := fmt.Sprintf("%v", x)
			wantMap[k] = true
		}

		t.Run(tt.name, func(t *testing.T) {
			got := subsetsWithDup(tt.nums)
			gotMap := make(map[string]bool)
			for _, x := range got {
				k := fmt.Sprintf("%v", x)
				gotMap[k] = true
			}

			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("subsetsWithDup() = %v, want %v", got, tt.want)
			}
		})
	}
}
