package _211

type WordDictionary struct {
	child [26]*WordDictionary
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}

func (w *WordDictionary) AddWord(word string) {
	cur := w
	for _, c := range word {
		if cur.child[c-'a'] == nil {
			cur.child[c-'a'] = &WordDictionary{}
		}
		cur = cur.child[c-'a']
	}
	cur.isEnd = true
}

func (w *WordDictionary) Search(word string) bool {
	if word == "" {
		return true
	}
	cur := w
	for i, c := range word {
		if c == '.' {
			for j, _ := range cur.child {
				sub := ""
				if i < len(word) - 1{
					sub = word[i+1:]
				}
				if cur.child[j] != nil && cur.child[j].Search(sub) {
					return true
				}
			}
			return false
		} else {
			if cur.child[c-'a'] == nil {
				return false
			}
			cur = cur.child[c-'a']
		}
	}
	return cur.isEnd
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
