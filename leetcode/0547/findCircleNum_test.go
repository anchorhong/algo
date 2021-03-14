package _547

import "testing"

func TestUnionFind_union(t *testing.T) {
	type fields struct {
		id  []int
		num int
	}
	type args struct {
		p int
		q int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//u := &UnionFind{
			//	id:  tt.fields.id,
			//	num: tt.fields.num,
			//}
		})
	}
}
