package _208

const alphaSize = 26

type Trie struct {
	children [alphaSize]*Trie
	isEnd    bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{isEnd: false}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this
	for _, s := range word {
		if cur.children[s-'a'] == nil {
			cur.children[s-'a'] = &Trie{}
		}
		cur = cur.children[s-'a']
	}
	cur.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this
	for _, s := range word {
		if cur.children[s-'a'] == nil {
			return false
		}
		cur = cur.children[s-'a']
	}
	return cur.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, s := range prefix {
		if cur.children[s-'a'] == nil {
			return false
		}
		cur = cur.children[s-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
