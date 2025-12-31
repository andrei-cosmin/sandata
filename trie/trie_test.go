/*
 * MIT License
 *
 * Copyright (c) 2025 Andrei Casu-Pop
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
 * documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
 * rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to
 * permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
 * Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
 * WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
 * OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package trie

import (
	"slices"
	"testing"

	"github.com/andrei-cosmin/sandata/chain"
	"github.com/stretchr/testify/assert"
)

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
	{keys: []string{"b"}},
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
	trie := New[string, int]()

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
	trie := New[string, int]()

	for _, entry := range insertedEntries {
		trie.Insert(entry.keys, entry.value)
	}

	for _, entry := range insertedEntries {
		result, found := trie.SearchChain(chain.New[string](entry.keys))
		assert.True(t, found)
		assert.Equal(t, entry.value, result)
	}
	for _, entry := range getInvalidEntries() {
		result, found := trie.SearchChain(chain.New[string](entry.keys))
		assert.False(t, found)
		assert.Equal(t, 0, result)
	}
}

func TestIteratorOverArray(t *testing.T) {
	trie := New[string, int]()

	for _, entry := range insertedEntries {
		trie.Insert(entry.keys, entry.value)
	}

	for _, entry := range insertedEntries {
		keys := entry.keys
		iter := trie.Iterator()

		index := 0
		for index = range keys {
			if !iter.Next(keys[index]) {
				break
			}
		}

		assert.True(t, iter.HasValue())
		assert.Equal(t, entry.value, iter.Value())
	}

	for _, entry := range insertedEntries {
		keys := entry.keys
		iter := trie.Iterator()

		index := 0
		for index < len(keys) && iter.Next(keys[index]) {
			index++
		}

		assert.True(t, iter.HasValue())
		assert.Equal(t, entry.value, iter.Value())
	}

	for _, entry := range getInvalidEntries() {
		keys := entry.keys
		iter := trie.Iterator()

		for index := range keys {
			if !iter.Next(keys[index]) {
				break
			}
		}

		assert.False(t, iter.HasValue())
	}

	for _, entry := range getInvalidEntries() {
		keys := entry.keys
		iter := trie.Iterator()

		index := 0
		for index < len(keys) && iter.Next(keys[index]) {
			index++
		}

		assert.False(t, iter.HasValue())
	}
}

func TestIteratorOverChain(t *testing.T) {
	trie := New[string, int]()

	for _, entry := range insertedEntries {
		trie.Insert(entry.keys, entry.value)
	}

	for _, entry := range insertedEntries {
		cursor := chain.New[string](entry.keys)
		iter := trie.Iterator()
		for cursor != nil && iter.Next(cursor.Data) {
			cursor = cursor.Next
		}

		assert.True(t, iter.HasValue())
		assert.Equal(t, entry.value, iter.Value())
	}

	for _, entry := range getInvalidEntries() {
		cursor := chain.New[string](entry.keys)
		iter := trie.Iterator()
		for cursor != nil && iter.Next(cursor.Data) {
			cursor = cursor.Next
		}

		assert.False(t, iter.HasValue())
	}
}
