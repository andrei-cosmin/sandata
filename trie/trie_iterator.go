package trie

// Iterator - interface for a generic trie iterator
type Iterator[K comparable, V any] interface {
	Next(K) bool
	HasValue() bool
	Value() V
}

// iterator - struct for a trie iterator
type iterator[K comparable, V any] struct {
	cursor *node[K, V]
}

// Next - moves the iterator to the next node
func (n *iterator[K, V]) Next(key K) bool {
	n.cursor = n.cursor.paths[key]
	return n.cursor != nil
}

// HasValue - checks if the current node has a value
func (n *iterator[K, V]) HasValue() bool {
	return n.cursor != nil && n.cursor.flag
}

// Value - returns the value of the current node
func (n *iterator[K, V]) Value() V {
	return n.cursor.data
}
