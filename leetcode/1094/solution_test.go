package _1094

import (
	"encoding/json"
	"testing"
)

func Test_carPooling(t *testing.T) {
	tests := []struct {
		name     string
		tripStr  string
		capacity int
		want     bool
	}{
		{
			name:     "testCase 1",
			tripStr:  "[[2,1,5],[3,3,7]]",
			capacity: 4,
			want:     false,
		},
		{
			name:     "testCase 2",
			tripStr:  "[[2,1,5],[3,3,7]]",
			capacity: 5,
			want:     true,
		},
		{
			name:     "testCase 3",
			tripStr:  "[[2,1,5],[3,5,7]]",
			capacity: 3,
			want:     true,
		},
		{
			name:     "testCase 4",
			tripStr:  "[[3,2,7],[3,7,9],[8,3,9]]",
			capacity: 11,
			want:     true,
		},
	}
	for _, tt := range tests {
		var trips [][]int
		json.Unmarshal([]byte(tt.tripStr), &trips)
		t.Run(tt.name, func(t *testing.T) {
			if got := carPooling(trips, tt.capacity); got != tt.want {
				t.Errorf("carPooling() = %v, want %v", got, tt.want)
			}
		})
	}
}
