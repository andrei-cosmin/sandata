package data

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func getHead[K comparable](node *ChainNode[K]) *ChainNode[K] {
	previous := node

	for node != nil {
		previous = node
		node = node.PreviousChain
	}

	return previous
}

func getTail[K comparable](node *ChainNode[K]) *ChainNode[K] {
	var next *ChainNode[K]

	for node != nil {
		next = node
		node = node.NextChain
	}

	return next
}

// newChain - creates a new chain from the given keys
func newChain[K comparable](keys []K) *ChainNode[K] {
	// If there are no keys, return nil
	if len(keys) == 0 {
		return nil
	}

	// Create the root node
	root := &ChainNode[K]{
		Data: keys[0],
	}

	// Set previous to nil
	var previous *ChainNode[K]
	// Set the cursor to the root
	cursor := root

	// Iterate over the keys and create the chain
	for index := 1; index < len(keys); index++ {
		// Create the next node
		cursor.NextChain = &ChainNode[K]{
			Data: keys[index],
		}
		// Set the previous node
		cursor.PreviousChain = previous

		// Set the cursor to the next node
		previous = cursor
		cursor = cursor.NextChain
	}

	// Set the previous node of the last node
	cursor.PreviousChain = previous

	// Return the root of the chain
	return root
}

type trieEntry struct {
	keys  []string
	value int
}

var insertedEntries = []trieEntry{
	{[]string{"a"}, 4},
	{[]string{"a", "a"}, -5},
	{[]string{"a", "a", "a"}, 0},
	{[]string{"a", "b", "c"}, 1},
	{[]string{"a", "b", "d"}, 2},
	{[]string{"a", "b", "e"}, 3},
	{[]string{"a", "b", "f"}, 4},
	{[]string{"a", "b", "g"}, 5},
	{[]string{"a", "b", "h"}, 6},
	{[]string{"a", "b", "i"}, 7},
	{[]string{"a", "b", "j"}, 8},
	{[]string{"a", "b", "k"}, 9},
	{[]string{"a", "b", "l"}, 10},
	{[]string{"a", "b", "m"}, 11},
	{[]string{"a", "b", "n"}, 12},
	{[]string{"a", "b", "o"}, 13},
	{[]string{"a", "b", "p"}, 14},
	{[]string{"a", "b", "q"}, 15},
	{[]string{"a", "b", "r"}, 16},
	{[]string{"a", "b", "s"}, 17},
	{[]string{"a", "b", "c", "x"}, 200},
	{[]string{"a", "b", "c", "x", "d"}, 201},
	{[]string{"a", "b", "c", "x", "d", "f"}, 201},
}

var invalidEntries = []trieEntry{
	{keys: []string{"a", "b"}},
	{keys: []string{"a", "b", "cccc"}},
	{keys: []string{"a", "b", "c", "x", "t"}},
	{keys: []string{"a", "b", "c", "x", "d", "y"}},
	{keys: []string{"a", "b", "c", "d"}},
}

func getInvalidEntries() []trieEntry {
	generatedEntries := slices.Clone(invalidEntries)
	for _, entry := range insertedEntries {
		invalidKeys := slices.Insert(slices.Clone(entry.keys), 0, "invalid")
		generatedEntries = append(invalidEntries, trieEntry{keys: invalidKeys})
	}
	return generatedEntries
}

func TestSearchKeys(t *testing.T) {
	trie := NewTrie[string, int]()

	for _, entry := range insertedEntries {
		trie.Insert(entry.keys, entry.value)
	}

	for _, entry := range insertedEntries {
		result, found := trie.SearchKeys(entry.keys)
		assert.True(t, found)
		assert.Equal(t, entry.value, result)
	}
	for _, entry := range getInvalidEntries() {
		result, found := trie.SearchKeys(entry.keys)
		assert.False(t, found)
		assert.Equal(t, 0, result)
	}
}

func TestSearchChain(t *testing.T) {
	trie := NewTrie[string, int]()

	for _, entry := range insertedEntries {
		trie.Insert(entry.keys, entry.value)
	}

	for _, entry := range insertedEntries {
		result, found := trie.SearchChain(newChain[string](entry.keys))
		assert.True(t, found)
		assert.Equal(t, entry.value, result)
	}
	for _, entry := range getInvalidEntries() {
		result, found := trie.SearchChain(newChain[string](entry.keys))
		assert.False(t, found)
		assert.Equal(t, 0, result)
	}
}

func TestIteratorOverArray(t *testing.T) {
	trie := NewTrie[string, int]()

	for _, entry := range insertedEntries {
		trie.Insert(entry.keys, entry.value)
	}

	for _, entry := range insertedEntries {
		keys := entry.keys
		iterator := trie.Iterator()

		index := 0
		for index = range keys {
			if !iterator.Next(keys[index]) {
				break
			}
		}

		assert.True(t, iterator.HasValue())
		assert.Equal(t, entry.value, iterator.Value())
	}

	for _, entry := range insertedEntries {
		keys := entry.keys
		iterator := trie.Iterator()

		index := 0
		for index < len(keys) && iterator.Next(keys[index]) {
			index++
		}

		assert.True(t, iterator.HasValue())
		assert.Equal(t, entry.value, iterator.Value())
	}

	for _, entry := range getInvalidEntries() {
		keys := entry.keys
		iterator := trie.Iterator()

		for index := range keys {
			if !iterator.Next(keys[index]) {
				break
			}
		}

		assert.False(t, iterator.HasValue())
	}

	for _, entry := range getInvalidEntries() {
		keys := entry.keys
		iterator := trie.Iterator()

		index := 0
		for index < len(keys) && iterator.Next(keys[index]) {
			index++
		}

		assert.False(t, iterator.HasValue())
	}
}

func TestIteratorOverChain(t *testing.T) {
	trie := NewTrie[string, int]()

	for _, entry := range insertedEntries {
		trie.Insert(entry.keys, entry.value)
	}

	for _, entry := range insertedEntries {
		chain := newChain[string](entry.keys)
		iterator := trie.Iterator()
		for chain != nil && iterator.Next(chain.Data) {
			chain = chain.NextChain
		}

		assert.True(t, iterator.HasValue())
		assert.Equal(t, entry.value, iterator.Value())
	}

	for _, entry := range getInvalidEntries() {
		chain := newChain[string](entry.keys)
		iterator := trie.Iterator()
		for chain != nil && iterator.Next(chain.Data) {
			chain = chain.NextChain
		}

		assert.False(t, iterator.HasValue())
	}
}
