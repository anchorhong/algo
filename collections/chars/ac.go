package chars

import (
	"container/list"
)

// AcNodeTree is the implement of Aho-Corasick algorithm
type AcNodeTree struct {
	root *AcNode
}

type AcNode struct {
	// used to save word
	Word string
	// whether is word
	isWord bool
	// the fail link point
	Fail     *AcNode
	Children [alphaSize]*AcNode
}

// NewAcNodeTrie builds a Aho-Corasick tree
func NewAcNodeTrie(words []string) *AcNodeTree {
	root := &AcNode{}
	for _, word := range words {
		cur := root
		for _, c := range word {
			if cur.Children[c-'a'] == nil {
				cur.Children[c-'a'] = &AcNode{}
			}
			cur = cur.Children[c-'a']
		}
		cur.isWord = true
		cur.Word = word
	}
	tree := AcNodeTree{root}
	tree.buildFailPoint()
	return &tree
}

func (a *AcNodeTree) buildFailPoint() {
	root := a.root
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		p := queue.Remove(queue.Front()).(*AcNode)
		for i := 0; i < alphaSize; i++ {
			pc := p.Children[i]
			if pc == nil {
				continue
			}
			if p == root {
				pc.Fail = root
			} else {
				f := p.Fail
				for f != nil {
					fc := f.Children[i]
					if fc != nil {
						pc.Fail = fc
						break
					}
					f = f.Fail
				}
				if f == nil {
					pc.Fail = root
				}
			}
			queue.PushBack(pc)
		}
	}
}

// Search searches the key words which exist in the text string
func (a AcNodeTree) Search(text string) (res []string) {
	cur := a.root
	for i := 0; i < len(text); i++ {
		idx := text[i] - 'a'
		for cur.Children[idx] == nil && cur != a.root {
			cur = cur.Fail
		}
		cur = cur.Children[idx]
		if cur == nil {
			cur = a.root
		}
		tmp := cur
		for tmp != a.root {
			if tmp.isWord {
				res = append(res, tmp.Word)
			}
			tmp = tmp.Fail
		}
	}
	return
}
