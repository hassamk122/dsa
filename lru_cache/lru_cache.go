package main

type LRUCache[K comparable, V any] struct {
	maxSize          int
	nodeMap          map[K]*Node[CacheElement[K, V]]
	doublyLinkedList *DoublyLinkedList[CacheElement[K, V]]
}

func NewLRUCache[K comparable, V any](maxSize int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		maxSize:          maxSize,
		nodeMap:          make(map[K]*Node[CacheElement[K, V]], maxSize),
		doublyLinkedList: NewDoublyLinkedList[CacheElement[K, V]](),
	}
}

func (lru *LRUCache[K, V]) Size() int {
	return lru.doublyLinkedList.size
}

func (lru *LRUCache[K, V]) Put(key K, value V) bool {
	newCacheElm := NewCacheElement(key, value)
	var llNode *Node[CacheElement[K, V]]

	if v, ok := lru.nodeMap[key]; ok && v != nil {
		v.Element = *newCacheElm
		lru.doublyLinkedList.MoveToFront(v)
		llNode = nil
	} else {
		if lru.Size() >= lru.maxSize {
			lru.evictElement()
		}

		llNode = lru.doublyLinkedList.Add(*newCacheElm)
		lru.nodeMap[key] = llNode
	}

	return llNode != nil
}

func (lru *LRUCache[K, V]) evictElement() {
	lruNode := lru.doublyLinkedList.RemoveLast()

	if !(lruNode == nil) {
		delete(lru.nodeMap, lruNode.Element.key)
	}
}

func (lru *LRUCache[K, V]) Get(key K) (V, bool) {
	var zero V
	node, ok := lru.nodeMap[key]
	if !ok || node == nil {
		return zero, false
	}

	lru.doublyLinkedList.MoveToFront(node)

	return node.Element.value, true
}
