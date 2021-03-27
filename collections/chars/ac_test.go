package chars

import (
	"reflect"
	"testing"
)

func TestMatch(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		text  string
		want  []string
	}{
		{
			name:  "testCase 1",
			words: []string{"go", "she", "test"},
			text:  "hahshetester",
			want:  []string{"she", "test"},
		},
	}
	for _, tt := range tests {
		ac := NewAcNodeTrie(tt.words)
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := ac.Search(tt.text); !reflect.DeepEqual(gotRes, tt.want) {
				t.Errorf("Search() = %v, want %v", gotRes, tt.want)
			}
		})
	}
}
