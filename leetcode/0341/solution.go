package _341

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedInteger interface {
	Add(elem NestedInteger)
	IsInteger() bool
	SetInteger(value int)
	GetInteger() int
	GetList() []*NestedInteger
}

type NestedIterator struct {
	elem []int
	size int
	pos  int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	nl := &NestedIterator{}
	var dfs func(NestedInteger)
	dfs = func(n NestedInteger) {
		if n.IsInteger() {
			nl.elem = append(nl.elem, n.GetInteger())
			nl.size++
		} else {
			for _, l := range n.GetList() {
				dfs(*l)
			}
		}
	}
	for _, nested := range nestedList {
		dfs(*nested)
	}
	return nl
}

func (n *NestedIterator) Next() int {
	old := n.pos
	n.pos++
	return n.elem[old]
}

func (n *NestedIterator) HasNext() bool {
	return n.pos < n.size
}
