package main

type CacheElement[K comparable, V any] struct {
	key   K
	value V
}

func NewCacheElement[K comparable, V any](key K, value V) *CacheElement[K, V] {
	return &CacheElement[K, V]{
		key:   key,
		value: value,
	}
}

func (ce *CacheElement[K, V]) GetKey() K {
	return ce.key
}

func (ce *CacheElement[K, V]) GetValue() V {
	return ce.value
}
