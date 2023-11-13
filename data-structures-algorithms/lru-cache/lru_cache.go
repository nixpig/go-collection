package lru_cache

type Key int
type Value int

type Cache struct {
	table       map[Key]CacheEntry[Key, Value]
	queue       QueueListNode[Value]
	maxSize     int
	currentSize int
}

type CacheEntry[K comparable, V any] struct {
	key   K
	value V
	node  QueueListNode[V]
}

type QueueList[V any] struct {
	front *QueueListNode[V]
	back  *QueueListNode[V]
}

type QueueListNode[V any] struct {
	value V
	next  *QueueListNode[V]
	prev  *QueueListNode[V]
}

func (q *QueueList[Value]) Enqueue(value Value) {
	node := QueueListNode[Value]{value: value}

	if q.back == nil {
		q.front = &node
		q.back = &node
	} else {
		q.back.next = &node
		q.back = &node
	}
}
