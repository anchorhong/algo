package _1603

import "testing"

func TestAddCar_Case(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want bool
	}{
		{
			name: "testCase 1",
			in:   1,
			want: true,
		},
		{
			name: "testCase 1",
			in:   2,
			want: true,
		},
		{
			name: "testCase 1",
			in:   3,
			want: false,
		},
		{
			name: "testCase 1",
			in:   1,
			want: false,
		},
	}
	p := Constructor(1, 1, 0)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.AddCar(tt.in); got != tt.want {
				t.Errorf("AddCar(%v) == %v, want %v", tt.in, got, tt.want)
			}
		})
	}

}
