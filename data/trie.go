package data

// Iterator - interface for a generic trie iterator
type Iterator[K comparable, V any] interface {
	Next(K) bool
	HasValue() bool
	Value() V
}

// trieIterator - struct for a trie iterator
type trieIterator[K comparable, V any] struct {
	cursor *trieNode[K, V]
}

// Next - moves the iterator to the next node
func (n *trieIterator[K, V]) Next(key K) bool {
	n.cursor = n.cursor.paths[key]
	return n.cursor != nil
}

// HasValue - checks if the current node has a value
func (n *trieIterator[K, V]) HasValue() bool {
	return n.cursor != nil && n.cursor.flag
}

// Value - returns the value of the current node
func (n *trieIterator[K, V]) Value() V {
	return n.cursor.data
}

// trieNode - struct for a trie node
//   - paths map[K]*trieNode[K, V] - map of paths to other nodes
//   - data V - data stored in the node
//   - flag bool - flag to indicate if the node has a value
type trieNode[K comparable, V any] struct {
	paths map[K]*trieNode[K, V]
	data  V
	flag  bool
}

// Trie - struct for a trie
//   - root *trieNode[K, V] - root node of the trie
//   - empty V - empty value for the trie
type Trie[K comparable, V any] struct {
	root  *trieNode[K, V]
	empty V
}

// NewTrie method - creates a new trie
func NewTrie[K comparable, V any]() *Trie[K, V] {
	return &Trie[K, V]{
		root: &trieNode[K, V]{
			paths: make(map[K]*trieNode[K, V]),
		},
	}
}

// Iterator - returns a new iterator for the trie set to the root
func (t *Trie[K, V]) Iterator() Iterator[K, V] {
	return &trieIterator[K, V]{
		cursor: t.root,
	}
}

// Insert - inserts a value into the trie using the given keys
func (t *Trie[K, V]) Insert(keys []K, value V) {
	// Set the cursor to the root
	cursor := t.root

	// Get the last key
	lastKey := keys[len(keys)-1]

	// Iterate over the keys, except the last key
	for index := range len(keys) - 1 {
		// Get the key at the current index
		key := keys[index]

		// Check if the key is in the paths, and create it, if it doesn't exist
		if _, ok := cursor.paths[key]; !ok {
			cursor.paths[key] = &trieNode[K, V]{}
		}

		// Move the cursor to the next node
		cursor = cursor.paths[key]

		// Create the paths map, if it doesn't exist
		if cursor.paths == nil {
			cursor.paths = make(map[K]*trieNode[K, V])
		}
	}

	// Create the node for the last key, containing the value and the flag set to true
	cursor.paths[lastKey] = &trieNode[K, V]{
		data: value,
		flag: true,
	}
}

// SearchKeys - searches for a value in the trie using the given keys
func (t *Trie[K, V]) SearchKeys(keys []K) (V, bool) {
	// Set the cursor to the root
	cursor := t.root

	// Iterate over the keys
	for _, key := range keys {
		// Check if the key is in the paths, and move the cursor to the next node
		if _, ok := cursor.paths[key]; ok {
			cursor = cursor.paths[key]
		} else {
			// Return the empty value and false, if the key is not found
			return t.empty, false
		}
	}

	// Return the value and the flag of the cursor
	return cursor.data, cursor.flag
}

// SearchChain - searches for a value in the trie using the given chain of keys
func (t *Trie[K, V]) SearchChain(chain *ChainNode[K]) (V, bool) {
	// Set the cursor to the root
	cursor := t.root

	for {
		// Get the key from the chain
		key := chain.Data

		// Check if the key is in the paths, and move the cursor to the next node
		if _, ok := cursor.paths[key]; ok {
			cursor = cursor.paths[key]
		} else {
			// Return the empty value and false, if the key is not found
			return t.empty, false
		}

		// Check if the chain has a next chain, and move the chain cursor to the next chain
		if chain.HasNext() {
			chain = chain.NextChain
		} else {
			// Return the value and the flag of the cursor
			return cursor.data, cursor.flag
		}
	}
}
