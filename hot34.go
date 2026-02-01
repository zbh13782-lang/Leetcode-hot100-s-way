package hot

import "container/list"

type entry struct {
	key, value int
}

type LRUCache struct {
	cachelist *list.List
	capacity  int
	keyToNode map[int]*list.Element
}

func Constructor1(capacity int) LRUCache {
	return LRUCache{
		cachelist: list.New(),
		capacity:  capacity,
		keyToNode: map[int]*list.Element{},
	}
}

func (this *LRUCache) Get(key int) int {
	if this.keyToNode[key] == nil {
		return -1
	}
	this.cachelist.MoveToFront(this.keyToNode[key])
	return this.keyToNode[key].Value.(entry).value
}

func (this *LRUCache) Put(key int, value int) {
	node := this.keyToNode[key]
	if node != nil {
		node.Value = entry{key: key, value: value}
		this.cachelist.MoveToFront(node)
	}
	this.keyToNode[key] = this.cachelist.PushFront(entry{key, value})
	if this.cachelist.Len() > this.capacity {
		backnode := this.cachelist.Remove(this.cachelist.Back())
		delete(this.keyToNode, backnode.(entry).key)
	}
}
