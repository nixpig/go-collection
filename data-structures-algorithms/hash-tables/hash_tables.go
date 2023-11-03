package hashtables

type HashTable[V comparable] struct {
	size int
	bins []*ListNode[V]
}

type ListNode[V comparable] struct {
	key   int
	value V
	next  *ListNode[V]
}

func hashFunction(key int, size int) int {
	constant := 23
	total := 0

	for i := range make([]int, key) {
		total = constant*total + int(i)
	}

	return total % size
}

func (ht *HashTable[V]) Insert(key int, value V) {
	hashValue := hashFunction(key, ht.size)

}
