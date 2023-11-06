package hashtables

import ()

const mul = 10

type HashTable[V comparable] struct {
	bins map[int]*ListNode[V]
}

type ListNode[V comparable] struct {
	key   int
	value V
	next  *ListNode[V]
}

func NewHashTable[V comparable]() HashTable[V] {
	ht := HashTable[V]{
		bins: make(map[int]*ListNode[V]),
	}

	return ht
}

func hashFunction(key int, size int) int {
	return size % key
}

func (ht *HashTable[V]) Insert(key int, value V) {
	hashValue := hashFunction(key, mul)

	if ht.bins[hashValue] == nil {
		ht.bins[hashValue] = &ListNode[V]{
			key:   key,
			value: value,
		}

		return
	}

	current := ht.bins[hashValue]

	for current.key != key && current.next != nil {
		current = current.next
	}

	if current.key == key {
		current.value = value
	} else {
		current.next = &ListNode[V]{
			key:   key,
			value: value,
		}
	}
}

func (ht *HashTable[V]) Lookup(key int) *ListNode[V] {
	hashValue := hashFunction(key, mul)

	current := ht.bins[hashValue]

	if current == nil {
		return nil
	}

	for current.key != key && current.next != nil {
		current = current.next
	}

	if current.key == key {
		return current
	}

	return nil
}

func (ht *HashTable[V]) Remove(key int) bool {
	hashValue := hashFunction(key, mul)

	current := ht.bins[hashValue]
	var previous *ListNode[V]

	if current == nil {
		return false
	}

	for current.key != key && current.next != nil {
		previous = current
		current = current.next
	}

	if current.key == key {
		if previous != nil {
			previous.next = current.next
		} else {
			ht.bins[hashValue] = current.next
		}

		return true
	}

	return false
}
