package _190

import "testing"

func Test_reverseBits(t *testing.T) {
	tests := []struct {
		name string
		num uint32
		want uint32
	}{
		{
			name: "testCase 1",
			num:  43261596,
			want: 964176192,
		},
		{
			name: "testCase 3",
			num:  4294967293,
			want: 3221225471,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits(tt.num); got != tt.want {
				t.Errorf("reverseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
