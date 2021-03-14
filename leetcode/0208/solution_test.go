package _208

import "testing"

func TestTrie_Search(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	tests := []struct {
		name string
		key  string
		want bool
	}{
		{
			name: "testCase 1",
			key:  "apple",
			want: true,
		},
		{
			name: "testCase 2",
			key:  "app",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trie.Search(tt.key); got != tt.want {
				t.Errorf("Search(%v) == %v, want %v", tt.key, got, tt.want)
			}
		})
	}
	trie.Insert("app")
	if got := trie.Search("app"); got != true {
		t.Errorf("Search(%v) == %v, want %v", "app", got, true)
	}

}

func TestTrie_StartsWith(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	tests := []struct {
		name string
		key  string
		want bool
	}{
		{
			name: "testCase 1",
			key:  "app",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trie.StartsWith(tt.key); got != tt.want {
				t.Errorf("StartsWith(%v) == %v, want %v", tt.key, got, tt.want)
			}
		})
	}
}
