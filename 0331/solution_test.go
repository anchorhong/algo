package _331

import "testing"

func Test_isValidSerialization(t *testing.T) {

	tests := []struct {
		name     string
		preorder string
		want     bool
	}{
		{
			name:     "testCase 1",
			preorder: "9,3,4,#,#,1,#,#,2,#,6,#,#",
			want:     true,
		},
		{
			name:     "testCase 2",
			preorder: "1,#",
			want:     false,
		},
		{
			name:     "testCase 3",
			preorder: "9,#,#,1",
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSerialization(tt.preorder); got != tt.want {
				t.Errorf("isValidSerialization() = %v, want %v", got, tt.want)
			}
		})
	}
}
