package ds

import "container/list"

type LRUCache struct {
	capacity int
	l        *list.List
	mapping  map[int]*node
}

type node struct {
	e     *list.Element
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		l:        list.New(),
		mapping:  make(map[int]*node, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if n, ok := this.mapping[key]; ok {
		this.l.MoveToFront(n.e)
		return n.value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if n, ok := this.mapping[key]; ok {
		this.l.MoveToFront(n.e)
		n.value = value
		return
	}

	if this.l.Len() >= this.capacity {
		lastElement := this.l.Back()
		this.l.Remove(lastElement)
		delete(this.mapping, lastElement.Value.(int))
	}

	// new key
	e := this.l.PushFront(key)
	n := &node{
		e: e,
		value: value,
	}
	this.mapping[key] = n
}
