package _721

import (
	"reflect"
	"testing"
)

func Test_accountsMerge1(t *testing.T) {
	type args struct {
		accounts [][]string
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]string
	}{
		{
			name: "testCase 1",
			args: args{
				accounts: [][]string{{"John", "johnsmith@mail.com", "john00@mail.com"}, {"John", "johnnybravo@mail.com"}, {"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"Mary", "mary@mail.com"}},
			},
			wantAns: [][]string{{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"}, {"John", "johnnybravo@mail.com"}, {"Mary", "mary@mail.com"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := accountsMerge1(tt.args.accounts); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("accountsMerge1() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
