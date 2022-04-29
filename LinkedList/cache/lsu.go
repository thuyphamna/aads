package cache

import (
	"errors"
	"fmt"
)

type LRUCache struct {
	capacity int
	size     int
	storage  map[int]*Node
	dblist   *DoublyLinkList
}

func CreateCache(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		storage:  make(map[int]*Node),
		dblist:   &DoublyLinkList{},
	}
}

func (cache *LRUCache) Get(key int) (int, error) {
	nodePointer, exist := cache.storage[key]
	if !exist {
		return 0, errors.New("Can't find the cache")
	} else {
		cache.dblist.MoveNodeToEnd(nodePointer)
		return nodePointer.value, nil
	}
}

func (cache *LRUCache) Put(key int, val int) {
	nodePointer, exist := cache.storage[key]
	if !exist {
		newNodePointer := &Node{
			key:   key,
			value: val,
		}
		cache.storage[key] = newNodePointer
		cache.dblist.AddToEnd(newNodePointer)
		cache.size++
		if cache.size > cache.capacity {
			lsuNode := cache.dblist.RemoveFromFront()
			delete(cache.storage, lsuNode.key)
			cache.size--
		}
	} else {
		nodePointer.value = val
		cache.dblist.MoveNodeToEnd(nodePointer)
	}
}

func (cache *LRUCache) PrintCache() {
	fmt.Printf("storage %v, size %v, capacity %v", cache.storage, cache.size, cache.capacity)
}
