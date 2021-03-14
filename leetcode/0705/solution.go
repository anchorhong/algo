package _705

import "container/list"

const bucketSize = 679

type MyHashSet struct {
	buckets []list.List
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		buckets: make([]list.List, bucketSize),
	}
}

func (this *MyHashSet) hash(key int) int {
	return key % bucketSize
}

func (this *MyHashSet) Add(key int) {
	bucketNum := this.hash(key)
	for elem := this.buckets[bucketNum].Front(); elem != nil; elem = elem.Next() {
		if elem.Value == key {
			return
		}
	}
	this.buckets[bucketNum].PushBack(key)
}

func (this *MyHashSet) Remove(key int) {
	bucketNum := this.hash(key)
	for elem := this.buckets[bucketNum].Front(); elem != nil; elem = elem.Next() {
		if elem.Value == key {
			this.buckets[bucketNum].Remove(elem)
			return
		}
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	bucketNum := this.hash(key)
	for elem := this.buckets[bucketNum].Front(); elem != nil; elem = elem.Next() {
		if elem.Value == key {
			return true
		}
	}
	return false
}
