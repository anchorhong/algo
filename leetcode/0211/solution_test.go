package _211

import "testing"

func TestWordDictionary_AddWord(t *testing.T) {

	wd := Constructor()
	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")

	tests := []struct {
		name   string
		search string
		want   bool
	}{
		{
			name:   "testCase1",
			search: "pad",
			want:   false,
		},
		{
			name:   "testCase2",
			search: "bad",
			want:   true,
		},
		{
			name:   "testCase3",
			search: ".ad",
			want:   true,
		},
		{
			name:   "testCase4",
			search: "b..",
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wd.Search(tt.search); got != tt.want {
				t.Errorf("Search(%v) == %v, want %v", tt.search, got, tt.want)
			}
		})
	}

}
