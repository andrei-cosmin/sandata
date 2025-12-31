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
