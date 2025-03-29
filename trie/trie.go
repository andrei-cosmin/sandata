package trie

import "github.com/andrei-cosmin/sandata/chain"

// node - struct for a trie node
//   - paths map[K]*node[K, V] - map of paths to other nodes
//   - data V - data stored in the node
//   - flag bool - flag to indicate if the node has a value
type node[K comparable, V any] struct {
	paths map[K]*node[K, V]
	data  V
	flag  bool
}

// Trie - struct for a trie
//   - root *node[K, V] - root node of the trie
//   - empty V - empty value for the trie
type Trie[K comparable, V any] struct {
	root  *node[K, V]
	empty V
}

// New method - creates a new trie
func New[K comparable, V any]() *Trie[K, V] {
	return &Trie[K, V]{
		root: &node[K, V]{
			paths: make(map[K]*node[K, V]),
		},
	}
}

// Iterator - returns a new iterator for the trie set to the root
func (t *Trie[K, V]) Iterator() Iterator[K, V] {
	return &iterator[K, V]{
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
			cursor.paths[key] = &node[K, V]{}
		}

		// Move the cursor to the next node
		cursor = cursor.paths[key]

		// Create the paths map, if it doesn't exist
		if cursor.paths == nil {
			cursor.paths = make(map[K]*node[K, V])
		}
	}

	// Create the node for the last key, containing the value and the flag set to true
	cursor.paths[lastKey] = &node[K, V]{
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
func (t *Trie[K, V]) SearchChain(chain *chain.Node[K]) (V, bool) {
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
			chain = chain.Next
		} else {
			// Return the value and the flag of the cursor
			return cursor.data, cursor.flag
		}
	}
}
