package _131

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want [][]string
	}{
		{
			name: "testCase 1",
			s:    "aab",
			want: [][]string{{"aa", "b"}, {"a", "a", "b"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partition(tt.s)
			gotMap := make(map[string]bool)
			for _, g := range got {
				gotMap[createKey(g)] = true
			}
			wantMap := make(map[string]bool)
			for _, w := range tt.want {
				wantMap[createKey(w)] = true
			}
			if got := partition(tt.s); !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createKey(s []string) string {
	return fmt.Sprintf("%q", s)
}
