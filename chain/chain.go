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

package chain

// Node represents a node in a double linked list
//   - Prev *Node[T] - reference to the previous node in the list
//   - Next *Node[T] - reference to the next node in the list
//   - Data T - the trie which is stored in the node
type Node[T any] struct {
	Prev *Node[T]
	Next *Node[T]
	Data T
}

// RemoveNode removes the node from the list, and links the previous and next nodes together
//
// Example (remove `B`):
//
//	A <-> B -> C  =>  A <-> C, B
func (c *Node[T]) RemoveNode() {
	// Get the reference of the previous node
	prev := c.Prev

	// Get the reference of the next node
	next := c.Next

	// If the previous node is not nil, link it to the next node
	if prev != nil {
		prev.Next = next
	}

	// If the next node is not nil, link it to the previous node
	if next != nil {
		next.Prev = prev
	}

	// Isolate the current node
	c.Prev = nil
	c.Next = nil
}

// InsertAsHead inserts the given node as the head of the list, and clears out references
//
// NOTE: the current node will become the head of the [other] node,
// the [other] node will become the tail of the current node
//
// Example (insert `B` as head of `A`):
//
//	List with `B`: X <-> B <-> Y <-> Z <-> T
//	List with `A`: C <-> A <-> D
//	Result:	X <-> B <-> A <-> D, Y <-> Z <-> T, C
func (c *Node[T]) InsertAsHead(other *Node[T]) {
	// Get the reference of the previous head
	prev := other.Prev

	// Get reference of the next node
	next := c.Next

	// Disconnects the previous head from the [other] node
	if prev != nil {
		prev.Next = nil
	}

	// Disconnects the previous tail from the current node
	if next != nil {
		next.Prev = nil
	}

	// Link the current node as head to the [other] node
	c.Next = other

	// Link the [other] node as tail to the current node
	other.Prev = c
}

// Split splits the list at the current node
//
// NOTE: the previous head will become a tail, the current node will become the head of a new list
//
// Example (split at `B`):
//
//	X <-> Y <-> B <-> Z <-> T  =>  X <-> Y, B <-> Z <-> T
func (c *Node[T]) Split() {
	// Get the reference of the previous node
	prev := c.Prev

	// Set the previous node as the tail of the list
	if prev != nil {
		prev.Next = nil
	}

	// Set the current node as the head of the list
	c.Prev = nil
}

// InsertAsTail inserts the given node as the tail of the list, and clears out references
//
// NOTE: the current node will become the tail of the [other] node, the [other] node will become the head of the current node
//
// Example (insert `A` as tail of `B`):
//
//	List with `A`: C <-> A <-> D
//	List with `B`: X <-> B <-> Y <-> Z <-> T
//	Result:	X <-> B <-> A <-> D, Y <-> Z <-> T, C
func (c *Node[T]) InsertAsTail(other *Node[T]) {
	// Get the reference of the previous head
	prev := c.Prev

	// Get reference of the next node
	next := other.Next

	// Disconnects the previous head from the current node
	if prev != nil {
		prev.Next = nil
	}

	// Disconnects the previous tail from the [other] node
	if next != nil {
		next.Prev = nil
	}

	// Link the current node as tail to the [other] node
	c.Prev = other

	// Link the [other] node as head to the current node
	other.Next = c
}

// HasNext returns true if the current node has a next node
func (c *Node[T]) HasNext() bool {
	return c.Next != nil
}

// HasPrevious returns true if the current node has a previous node
func (c *Node[T]) HasPrevious() bool {
	return c.Prev != nil
}

// IsHead returns true if the current node is the head of the list
func (c *Node[T]) IsHead() bool {
	return c.Prev == nil
}

// IsTail returns true if the current node is the tail of the list
func (c *Node[T]) IsTail() bool {
	return c.Next == nil
}

// HeadOf returns the head of the chain
func HeadOf[K comparable](node *Node[K]) *Node[K] {
	previous := node

	for node != nil {
		previous = node
		node = node.Prev
	}

	return previous
}

// TailOf returns the tail of the chain
func TailOf[K comparable](node *Node[K]) *Node[K] {
	var next *Node[K]

	for node != nil {
		next = node
		node = node.Next
	}

	return next
}

// New creates a new chain from the given keys
func New[K comparable](keys []K) *Node[K] {
	// If there are no keys, return nil
	if len(keys) == 0 {
		return nil
	}

	// Create the root node
	root := &Node[K]{
		Data: keys[0],
	}

	// Set previous to nil
	var prev *Node[K]
	// Set the cursor to the root
	cursor := root

	// Iterate over the keys and create the chain
	for index := 1; index < len(keys); index++ {
		// Create the next node
		cursor.Next = &Node[K]{
			Data: keys[index],
		}
		// Set the previous node
		cursor.Prev = prev

		// Set the cursor to the next node
		prev = cursor
		cursor = cursor.Next
	}

	// Set the previous node of the last node
	cursor.Prev = prev

	// Return the root of the chain
	return root
}
