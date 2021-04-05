package _344

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s    []byte
		want []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testCase 1",
			args: args{
				s:    []byte{'h', 'e', 'l', 'l', 'o'},
				want: []byte{'o', 'l', 'l', 'e', 'h'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			if !reflect.DeepEqual(tt.args.s, tt.args.want) {
				t.Errorf("want %v, but %v", tt.args.want, tt.args.s)
			}
		})
	}
}
