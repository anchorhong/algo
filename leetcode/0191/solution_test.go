package _191

import "testing"

func Test_hammingWeight(t *testing.T) {
	tests := []struct {
		name string
		num  uint32
		want int
	}{
		{
			name: "testCase 1",
			num:  00000000000000000000000000001011,
			want: 3,
		},
		{
			name: "testCase 2",
			num:  00000000000000000000000010000000,
			want: 1,
		},
		{
			name: "testCase 3",
			num:  11111111111111111111111111111101,
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.num); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
